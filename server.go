// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/shuLhan/share/lib/debug"
	libhttp "github.com/shuLhan/share/lib/http"
	"github.com/shuLhan/share/lib/memfs"
)

//
// server contains the HTTP server that serve the generated HTML files.
//
type server struct {
	http    *libhttp.Server
	opts    *libhttp.ServerOptions
	htmlg   *htmlGenerator
	watcher *watcher
}

//
// newServer create an HTTP server to serve HTML files in directory "root".
//
// The address parameter is optional, if not set its default to ":8080".
// The htmlTemplate parameter is optional, if not set its default to
// embedded HTML template.
//
func newServer(mfs *memfs.MemFS, root, address, htmlTemplate string) (srv *server) {
	var err error

	srv = &server{
		opts: &libhttp.ServerOptions{
			Options: memfs.Options{
				Root:        root,
				Excludes:    defExcludes,
				Development: debug.Value > 0,
			},
			Memfs:   mfs,
			Address: address,
		},
	}

	srv.http, err = libhttp.NewServer(srv.opts)
	if err != nil {
		log.Fatal("ciigo: libhttp.NewServer: " + err.Error())
	}

	epInSearch := &libhttp.Endpoint{
		Method:       libhttp.RequestMethodGet,
		Path:         "/_internal/search",
		RequestType:  libhttp.RequestTypeQuery,
		ResponseType: libhttp.ResponseTypeHTML,
		Call:         srv.onSearch,
	}

	err = srv.http.RegisterEndpoint(epInSearch)
	if err != nil {
		log.Fatal("ciigo: " + err.Error())
	}

	srv.htmlg, err = newHTMLGenerator(mfs, htmlTemplate, srv.opts.Development)
	if err != nil {
		log.Fatal("ciigo: " + err.Error())
	}

	if srv.opts.Development {
		srv.watcher, err = newWatcher(srv.htmlg, root)
		if err != nil {
			log.Fatal("ciigo: " + err.Error())
		}

		srv.watcher.fileMarkups = listFileMarkups(root)
		srv.htmlg.convertFileMarkups(srv.watcher.fileMarkups)
	}

	return srv
}

//
// start the web server.
//
func (srv *server) start() {
	if srv.opts.Development {
		err := srv.watcher.start()
		if err != nil {
			log.Fatal("ciigo: " + err.Error())
		}
	}

	fmt.Printf("ciigo: starting HTTP server at %q for %q\n",
		srv.opts.Address, srv.opts.Root)

	err := srv.http.Start()
	if err != nil {
		log.Fatal("ciigo: " + err.Error())
	}
}

func (srv *server) onSearch(res http.ResponseWriter, req *http.Request, reqBody []byte) (
	resBody []byte, err error,
) {
	var bufSearch, buf bytes.Buffer

	q := req.Form.Get("q")
	results := srv.http.Memfs.Search(strings.Fields(q), 0)

	err = srv.htmlg.tmplSearch.Execute(&bufSearch, results)
	if err != nil {
		return nil, fmt.Errorf("ciigo.onSearch: " + err.Error())
	}

	fhtml := &fileHTML{
		Body: template.HTML(bufSearch.String()), //nolint: gosec
	}

	err = srv.htmlg.tmpl.Execute(&buf, fhtml)
	if err != nil {
		return nil, fmt.Errorf("ciigo.onSearch: " + err.Error())
	}

	resBody = buf.Bytes()

	return resBody, nil
}
