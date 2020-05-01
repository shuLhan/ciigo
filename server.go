// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/shuLhan/share/lib/debug"
	libhttp "github.com/shuLhan/share/lib/http"
	libio "github.com/shuLhan/share/lib/io"
)

//
// server contains the HTTP server.
//
type server struct {
	http        *libhttp.Server
	opts        *libhttp.ServerOptions
	htmlg       *htmlGenerator
	fileMarkups []*fileMarkup
	dw          *libio.DirWatcher
}

//
// newServer create an HTTP server to serve HTML files in directory "root".
//
// The address parameter is optional, if not set its default to ":8080".
// The htmlTemplate parameter is optional, if not set its default to
// embedded HTML template.
//
func newServer(root, address, htmlTemplate string) (srv *server) {
	var err error

	srv = &server{
		opts: &libhttp.ServerOptions{
			Address:     address,
			Root:        root,
			Excludes:    defExcludes,
			Development: debug.Value > 0,
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

	srv.initHTMLGenerator(htmlTemplate)

	if srv.opts.Development {
		srv.fileMarkups = listFileMarkups(root)
		srv.htmlg.convertFileMarkups(srv.fileMarkups, false)
	}

	return srv
}

//
// start the web server.
//
func (srv *server) start() {
	if srv.opts.Development {
		srv.autoGenerate()
	}

	fmt.Printf("ciigo: starting HTTP server at %q for %q\n",
		srv.opts.Address, srv.opts.Root)

	err := srv.http.Start()
	if err != nil {
		log.Fatal("web: Start: " + err.Error())
	}
}

func (srv *server) autoGenerate() {
	srv.dw = &libio.DirWatcher{
		Path:  srv.opts.Root,
		Delay: time.Second,
		Includes: []string{
			`.*\.adoc$`,
			`.*\.md$`,
		},
		Excludes: []string{
			`assets/.*`,
			`.*\.html$`,
			`^\..*`,
		},
		Callback: srv.onChangeFileMarkup,
	}

	err := srv.dw.Start()
	if err != nil {
		log.Fatal("ciigo: autoGenerate: " + err.Error())
	}

	if len(srv.htmlg.path) > 0 {
		_, err = libio.NewWatcher(srv.htmlg.path, 0, srv.onChangeHTMLTemplate)
		if err != nil {
			log.Fatal("ciigo: autoGenerate: " + err.Error())
		}
	}
}

func (srv *server) initHTMLGenerator(htmlTemplate string) {
	if len(htmlTemplate) == 0 {
		srv.htmlg = newHTMLGenerator("", templateIndexHTML)
		return
	}

	var (
		bhtml       []byte
		err         error
		htmlContent string
	)

	htmlTemplate = filepath.Clean(htmlTemplate)

	if srv.opts.Development {
		bhtml, err = ioutil.ReadFile(htmlTemplate)
		if err != nil {
			log.Fatal("server.initHTMLGenerator: " + err.Error())
		}
	} else {
		tmplNode, err := srv.http.Memfs.Get(htmlTemplate)
		if err != nil {
			log.Fatalf("server.initHTMLGenerator: Memfs.Get %s: %s",
				htmlTemplate, err.Error())
		}
		bhtml, err = tmplNode.Decode()
		if err != nil {
			log.Fatal("server.initHTMLGenerator: " + err.Error())
		}
	}

	htmlContent = string(bhtml)
	srv.htmlg = newHTMLGenerator("", htmlContent)
}

//
// onChangeFileMarkup watch the markup files inside the "content" directory,
// and re-generate them into HTML file when changed.
//
func (srv *server) onChangeFileMarkup(ns *libio.NodeState) {
	if ns.State == libio.FileStateDeleted {
		fmt.Printf("ciigo: onChangeFileMarkup: %q deleted\n", ns.Node.SysPath)
		return
	}

	ext := strings.ToLower(path.Ext(ns.Node.SysPath))
	if !isExtensionMarkup(ext) {
		return
	}

	fmt.Println("ciigo: onChangeFileMarkup: " + ns.Node.SysPath)

	var (
		fmarkup *fileMarkup
		err     error
	)

	switch ns.State {
	case libio.FileStateCreated:
		fmarkup, err = newFileMarkup(ns.Node.SysPath, nil)
		if err != nil {
			log.Println(err)
			return
		}

		srv.fileMarkups = append(srv.fileMarkups, fmarkup)

	case libio.FileStateModified:
		for x := 0; x < len(srv.fileMarkups); x++ {
			if srv.fileMarkups[x].path == ns.Node.SysPath {
				fmarkup = srv.fileMarkups[x]
				break
			}
		}
		if fmarkup == nil {
			fmarkup, err = newFileMarkup(ns.Node.SysPath, nil)
			if err != nil {
				log.Println(err)
				return
			}

			srv.fileMarkups = append(srv.fileMarkups, fmarkup)
		}
	}

	fhtml := &fileHTML{
		path: fmarkup.basePath + ".html",
	}

	fhtml.rawBody.Reset()
	srv.htmlg.convert(fmarkup, fhtml, true)
}

func (srv *server) onChangeHTMLTemplate(ns *libio.NodeState) {
	if ns.State == libio.FileStateDeleted {
		fmt.Printf("watchHTMLTemplate: file %q deleted\n", ns.Node.SysPath)
		return
	}

	fmt.Println("web: recompiling HTML template  ...")

	err := srv.htmlg.reloadTemplate()
	if err != nil {
		log.Println("watchHTMLTemplate: loadTemplate: " + err.Error())
		return
	}

	fmt.Println("web: regenerate all markup files ... ")

	srv.htmlg.convertFileMarkups(srv.fileMarkups, true)
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
