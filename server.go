// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	libhttp "github.com/shuLhan/share/lib/http"
	"github.com/shuLhan/share/lib/memfs"
)

// server contains the HTTP server that serve the generated HTML files.
type server struct {
	http    *libhttp.Server
	htmlg   *htmlGenerator
	watcher *watcher
	opts    ServeOptions
}

// newServer create an HTTP server to serve HTML files in directory "root".
//
// The address parameter is optional, if not set its default to ":8080".
// The htmlTemplate parameter is optional, if not set its default to
// embedded HTML template.
func newServer(opts *ServeOptions) (srv *server, err error) {
	var (
		logp = "newServer"
	)

	if opts.Mfs == nil {
		opts.Mfs = &memfs.MemFS{
			Opts: &memfs.Options{
				Root:     opts.Root,
				Excludes: defExcludes,
			},
		}
		opts.IsDevelopment = true
	}

	opts.Mfs.Opts.TryDirect = opts.IsDevelopment

	srv = &server{
		opts: *opts,
	}

	httpdOpts := &libhttp.ServerOptions{
		Memfs:   opts.Mfs,
		Address: opts.Address,
	}

	srv.http, err = libhttp.NewServer(httpdOpts)
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

	srv.htmlg, err = newHTMLGenerator(opts.Mfs, opts.HtmlTemplate, opts.IsDevelopment)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", logp, err)
	}

	if opts.IsDevelopment {
		srv.watcher, err = newWatcher(srv.htmlg, &opts.ConvertOptions)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", logp, err)
		}

		srv.htmlg.convertFileMarkups(srv.watcher.fileMarkups, false)
	}

	return srv, nil
}

// start the web server.
func (srv *server) start() (err error) {
	logp := "start"

	if srv.opts.IsDevelopment {
		err := srv.watcher.start()
		if err != nil {
			return fmt.Errorf("%s: %w", logp, err)
		}
	}

	fmt.Printf("ciigo: starting HTTP server at %q for %q\n",
		srv.http.Options.Address, srv.http.Options.Memfs.Opts.Root)

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
	results := srv.http.Options.Memfs.Search(strings.Fields(q), 0)

	err = srv.htmlg.tmplSearch.Execute(&bufSearch, results)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", logp, err)
	}

	fhtml := &fileHtml{
		Body: template.HTML(bufSearch.String()), //nolint: gosec
	}

	err = srv.htmlg.tmpl.Execute(&buf, fhtml)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", logp, err)
	}

	resBody = buf.Bytes()

	return resBody, nil
}
