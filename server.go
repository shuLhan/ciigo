// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

import (
	"fmt"
	"log"
	"path"
	"strings"
	"time"

	"github.com/shuLhan/share/lib/debug"
	libhttp "github.com/shuLhan/share/lib/http"
	libio "github.com/shuLhan/share/lib/io"
)

//
// Server contains the HTTP server.
//
type Server struct {
	http        *libhttp.Server
	opts        *libhttp.ServerOptions
	htmlg       *htmlGenerator
	markupFiles []*markupFile
	dw          *libio.DirWatcher
}

//
// NewServer create an HTTP server to serve HTML files in directory "content".
// The address parameter is optional, if not set its default to ":8080".
//
func NewServer(root, address string) (srv *Server) {
	var err error

	if len(root) == 0 {
		root = dirRoot
	}
	if len(address) == 0 {
		address = ":8080"
	}

	srv = &Server{
		opts: &libhttp.ServerOptions{
			Address:     address,
			Root:        root,
			Excludes:    defExcludes,
			Development: debug.Value > 0,
		},
	}

	srv.http, err = libhttp.NewServer(srv.opts)
	if err != nil {
		log.Fatal("web: libhttp.NewServer: " + err.Error())
	}

	if srv.opts.Development {
		srv.htmlg = newHTMLGenerator()
		srv.markupFiles = listMarkupFiles(root)
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

	fmt.Printf("ciigo: starting HTTP server at %s for %s\n",
		srv.opts.Address, srv.opts.Root)

	err := srv.http.Start()
	if err != nil {
		log.Fatal("web: Start: " + err.Error())
	}
}

func (srv *Server) autoGenerate() {
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
		Callback: srv.onChangeMarkupFile,
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
// onChangeMarkupFile watch the markup files inside the "content" directory,
// and re-generate them into HTML file when changed.
//
func (srv *Server) onChangeMarkupFile(ns *libio.NodeState) {
	if ns.State == libio.FileStateDeleted {
		fmt.Printf("ciigo: onChangeMarkupFile: %q deleted\n", ns.Node.SysPath)
		return
	}

	ext := strings.ToLower(path.Ext(ns.Node.SysPath))
	if !isExtensionMarkup(ext) {
		return
	}

	fmt.Println("ciigo: onChangeMarkupFile: " + ns.Node.SysPath)

	var (
		fmarkup *markupFile
		err     error
	)

	switch ns.State {
	case libio.FileStateCreated:
		fmarkup, err = newMarkupFile(ns.Node.SysPath, nil)
		if err != nil {
			log.Println(err)
			return
		}

		srv.markupFiles = append(srv.markupFiles, fmarkup)

	case libio.FileStateModified:
		for x := 0; x < len(srv.markupFiles); x++ {
			if srv.markupFiles[x].path == ns.Node.SysPath {
				fmarkup = srv.markupFiles[x]
				break
			}
		}
		if fmarkup == nil {
			fmarkup, err = newMarkupFile(ns.Node.SysPath, nil)
			if err != nil {
				log.Println(err)
				return
			}

			srv.markupFiles = append(srv.markupFiles, fmarkup)
		}
	}

	fhtml := &fileHTML{
		path: fmarkup.basePath + ".html",
	}

	fhtml.rawBody.Reset()
	srv.htmlg.convert(fmarkup, fhtml, true)
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

	fmt.Println("web: regenerate all markup files ... ")

	srv.htmlg.convertMarkupFiles(srv.markupFiles, true)
}
