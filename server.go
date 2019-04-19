// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

import (
	"fmt"
	"log"
	"path"
	"time"

	"github.com/shuLhan/share/lib/debug"
	libhttp "github.com/shuLhan/share/lib/http"
	libio "github.com/shuLhan/share/lib/io"
)

//
// Server contains the HTTP server.
//
type Server struct {
	http  *libhttp.Server
	opts  *libhttp.ServerOptions
	htmlg *htmlGenerator
	adocs []*fileAdoc
	dw    *libio.DirWatcher
}

//
// NewServer create an HTTP server to serve HTML files in directory "content".
// The address parameter is optional, if not set its default to ":8080".
//
func NewServer(address string) (srv *Server) {
	var err error

	if len(address) == 0 {
		address = ":8080"
	}

	srv = new(Server)

	srv.opts = &libhttp.ServerOptions{
		Address: address,
		Root:    dirRoot,
		Excludes: []string{
			`.*\.adoc$`,
		},
		Development: debug.Value > 0,
	}

	srv.http, err = libhttp.NewServer(srv.opts)
	if err != nil {
		log.Fatal("web: libhttp.NewServer: " + err.Error())
	}

	if srv.opts.Development {
		srv.htmlg = newHTMLGenerator()
		srv.adocs = listAdocFiles(dirRoot)
	}

	return srv
}

//
// Start the web server.
//
func (srv *Server) Start() {
	if srv.opts.Development {
		srv.autoGenerate()
	}

	err := srv.http.Start()
	if err != nil {
		log.Fatal("web: Start: " + err.Error())
	}
}

func (srv *Server) autoGenerate() {
	srv.dw = &libio.DirWatcher{
		Path:  dirRoot,
		Delay: time.Second,
		Includes: []string{
			`.*\.adoc$`,
		},
		Excludes: []string{
			`assets/.*`,
			`.*\.html$`,
		},
		Callback: srv.onChangeAdoc,
	}

	err := srv.dw.Start()
	if err != nil {
		log.Fatal("ciigo: autoGenerate: " + err.Error())
	}

	_, err = libio.NewWatcher(srv.htmlg.path, 0, srv.onChangeHTMLTemplate)
	if err != nil {
		log.Fatal("ciigo: autoGenerate: " + err.Error())
	}
}

//
// onChangeAdoc watch the ".adoc" file in the "content" directory, and
// re-generate the HTML file if changed.
//
func (srv *Server) onChangeAdoc(ns *libio.NodeState) {
	if ns.State == libio.FileStateDeleted {
		fmt.Printf("ciigo: onChangeAdoc: %q deleted\n", ns.Node.SysPath)
		return
	}

	ext := path.Ext(ns.Node.SysPath)
	if ext != ".adoc" {
		return
	}

	fmt.Println("ciigo: onChangeAdoc: " + ns.Node.SysPath)

	var (
		adoc *fileAdoc
		err  error
	)

	switch ns.State {
	case libio.FileStateCreated:
		adoc, err = newFileAdoc(ns.Node.SysPath, nil)
		if err != nil {
			log.Println(err)
			return
		}

		srv.adocs = append(srv.adocs, adoc)

	case libio.FileStateModified:
		for x := 0; x < len(srv.adocs); x++ {
			if srv.adocs[x].path == ns.Node.SysPath {
				adoc = srv.adocs[x]
				break
			}
		}
		if adoc == nil {
			adoc, err = newFileAdoc(ns.Node.SysPath, nil)
			if err != nil {
				log.Println(err)
				return
			}

			srv.adocs = append(srv.adocs, adoc)
		}
	}

	fhtml := &fileHTML{
		path: adoc.basePath + ".html",
	}

	fhtml.rawBody.Reset()
	adoc.toHTML(fhtml.path, &fhtml.rawBody, true)
	fhtml.unpackAdoc(adoc)

	srv.htmlg.write(fhtml)
}

func (srv *Server) onChangeHTMLTemplate(ns *libio.NodeState) {
	if ns.State == libio.FileStateDeleted {
		fmt.Printf("watchHTMLTemplate: file %q deleted\n", ns.Node.SysPath)
		return
	}

	fmt.Println("web: recompiling HTML template  ...")

	err := srv.htmlg.loadTemplate()
	if err != nil {
		log.Println("watchHTMLTemplate: loadTemplate: " + err.Error())
		return
	}

	fmt.Println("web: regenerate all .adoc files ... ")

	srv.htmlg.convertAdocs(srv.adocs, true)
}
