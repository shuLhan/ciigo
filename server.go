// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
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
	if opts.IsDevelopment {
		httpdOpts.HandleFS = ciigo.onGet
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

// onGet when user reload the page from browser, inspect the HTML file by
// checking if its older that the adoc.
// If yes, it will auto convert the adoc and return the new content of HTML
// files.
func (ciigo *Ciigo) onGet(
	node *memfs.Node, _ http.ResponseWriter, req *http.Request,
) (out *memfs.Node) {
	var (
		logp = `onGet`
		file string
	)

	if node == nil {
		file = req.URL.Path
	} else {
		if node.IsDir() {
			file = path.Join(node.Path, `index.html`)
		} else {
			if len(req.URL.Path) > len(node.Path) {
				file = req.URL.Path
			} else {
				file = node.Path
			}
		}
	}
	if file[len(file)-1] == '/' {
		file = path.Join(file, `index.html`)
	}

	var (
		fmarkup *FileMarkup
		isNew   bool
	)
	fmarkup, isNew = ciigo.watcher.getFileMarkupByHTML(file)
	if fmarkup == nil {
		// File is not HTML or no markup files created from it.
		return node
	}
	var err error
	if isNew || ciigo.converter.shouldConvert(fmarkup) {
		err = ciigo.converter.ToHTMLFile(fmarkup)
		if err != nil {
			log.Printf(`%s: failed to convert markup file %q: %s`,
				logp, fmarkup.path, err)
			return node
		}
	}
	out, err = ciigo.serveOpts.Mfs.Get(file)
	if err != nil {
		log.Printf(`%s: failed to get %q: %s`, logp, file, err)
		return node
	}
	return out
}
