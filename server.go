// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strings"

	libhttp "git.sr.ht/~shulhan/pakakeh.go/lib/http"
	"git.sr.ht/~shulhan/pakakeh.go/lib/memfs"
)

// InitHTTPServer create an HTTP server to serve HTML files in directory
// defined in "[ConvertOptions].Root".
//
// The address parameter is optional, if not set its default to ":8080".
// The htmlTemplate parameter is optional, if not set its default to
// embedded HTML template.
func (ciigo *Ciigo) InitHTTPServer(opts ServeOptions) (err error) {
	var logp = `initServer`

	err = opts.init()
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	if opts.Mfs == nil {
		opts.IsDevelopment = true
		var mfsopts = &memfs.Options{
			Root:      opts.Root,
			Excludes:  defExcludes,
			TryDirect: true,
		}
		opts.Mfs, err = memfs.New(mfsopts)
		if err != nil {
			return fmt.Errorf(`%s: %w`, logp, err)
		}
	} else {
		opts.Mfs.Opts.TryDirect = opts.IsDevelopment
	}

	ciigo.serveOpts = opts

	var httpdOpts = libhttp.ServerOptions{
		Memfs:           opts.Mfs,
		Address:         opts.Address,
		EnableIndexHTML: opts.EnableIndexHTML,
	}

	ciigo.HTTPServer, err = libhttp.NewServer(httpdOpts)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	var epInSearch = libhttp.Endpoint{
		Method:       libhttp.RequestMethodGet,
		Path:         `/_internal/search`,
		RequestType:  libhttp.RequestTypeQuery,
		ResponseType: libhttp.ResponseTypeHTML,
		Call:         ciigo.onSearch,
	}

	err = ciigo.HTTPServer.RegisterEndpoint(epInSearch)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	var pathHTMLTemplate string

	if opts.IsDevelopment {
		pathHTMLTemplate = opts.HTMLTemplate
	}

	ciigo.converter, err = NewConverter(pathHTMLTemplate)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	if !opts.IsDevelopment {
		var tmplNode *memfs.Node

		tmplNode, _ = opts.Mfs.Get(internalTemplatePath)
		if tmplNode != nil {
			ciigo.converter.tmpl, err = ciigo.converter.tmpl.Parse(string(tmplNode.Content))
			if err != nil {
				return fmt.Errorf(`%s: %s`, logp, err)
			}
		}
	}

	if opts.IsDevelopment {
		ciigo.watcher, err = newWatcher(ciigo.converter, opts.ConvertOptions)
		if err != nil {
			return fmt.Errorf(`%s: %w`, logp, err)
		}

		ciigo.converter.convertFileMarkups(ciigo.watcher.fileMarkups, false)
	}

	return nil
}

// Serve start the HTTP web server.
func (ciigo *Ciigo) Serve() (err error) {
	var logp = `Serve`

	if ciigo.serveOpts.IsDevelopment {
		err = ciigo.watcher.start()
		if err != nil {
			return fmt.Errorf(`%s: %w`, logp, err)
		}
	}

	log.Printf(`ciigo: starting HTTP server at http://%s for %q`,
		ciigo.HTTPServer.Options.Address,
		ciigo.HTTPServer.Options.Memfs.Opts.Root)

	err = ciigo.HTTPServer.Start()
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	if ciigo.serveOpts.IsDevelopment {
		ciigo.watcher.stop()
	}

	return nil
}

func (ciigo *Ciigo) onSearch(epr *libhttp.EndpointRequest) (resBody []byte, err error) {
	var (
		logp = `onSearch`

		fhtml   *fileHTML
		buf     bytes.Buffer
		q       string
		results []memfs.SearchResult
	)

	q = epr.HTTPRequest.Form.Get(`q`)
	results = ciigo.HTTPServer.Options.Memfs.Search(strings.Fields(q), 0)

	err = ciigo.converter.tmplSearch.Execute(&buf, results)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	fhtml = &fileHTML{
		Body: template.HTML(buf.String()),
	}

	buf.Reset()

	err = ciigo.converter.tmpl.Execute(&buf, fhtml)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	resBody = buf.Bytes()

	return resBody, nil
}
