// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

import (
	"bytes"
	"fmt"
	"html/template"
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
func newServer(opts *ServeOptions) (srv *server, err error) {
	logp := "newServer"

	srv = &server{
		opts: &libhttp.ServerOptions{
			Options: memfs.Options{
				Root:        opts.Root,
				Excludes:    defExcludes,
				Development: debug.Value > 0,
			},
			Memfs:   opts.Mfs,
			Address: opts.Address,
		},
	}

	srv.http, err = libhttp.NewServer(srv.opts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", logp, err)
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
		return nil, fmt.Errorf("%s: %w", logp, err)
	}

	srv.htmlg, err = newHTMLGenerator(opts.Mfs, opts.HtmlTemplate, srv.opts.Development)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", logp, err)
	}

	if srv.opts.Development {
		srv.watcher, err = newWatcher(srv.htmlg, opts.Root)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", logp, err)
		}

		srv.watcher.fileMarkups, err = listFileMarkups(opts.Root)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", logp, err)
		}

		srv.htmlg.convertFileMarkups(srv.watcher.fileMarkups)
	}

	return srv, nil
}

//
// start the web server.
//
func (srv *server) start() (err error) {
	logp := "start"

	if srv.opts.Development {
		err := srv.watcher.start()
		if err != nil {
			return fmt.Errorf("%s: %w", logp, err)
		}
	}

	fmt.Printf("ciigo: starting HTTP server at %q for %q\n",
		srv.opts.Address, srv.opts.Root)

	err = srv.http.Start()
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	return nil
}

func (srv *server) onSearch(epr *libhttp.EndpointRequest) (resBody []byte, err error) {
	var bufSearch, buf bytes.Buffer
	logp := "onSearch"

	q := epr.HttpRequest.Form.Get("q")
	results := srv.http.Memfs.Search(strings.Fields(q), 0)

	err = srv.htmlg.tmplSearch.Execute(&bufSearch, results)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", logp, err)
	}

	fhtml := &fileHTML{
		Body: template.HTML(bufSearch.String()), //nolint: gosec
	}

	err = srv.htmlg.tmpl.Execute(&buf, fhtml)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", logp, err)
	}

	resBody = buf.Bytes()

	return resBody, nil
}
