// SPDX-License-Identifier: GPL-3.0-only
// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>

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
// This method initialize the [Ciigo.HTTPServer] so one can register
// additional endpoints before running the HTTP server.
func (ciigo *Ciigo) InitHTTPServer(serveOpts ServeOptions, convertOpts ConvertOptions) (err error) {
	var logp = `InitHTTPServer`

	err = convertOpts.init()
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	serveOpts.init()

	if serveOpts.Memfs == nil {
		serveOpts.IsDevelopment = true
		var mfsopts = &memfs.Options{
			Root:      convertOpts.Root,
			Excludes:  defExcludes,
			TryDirect: true,
		}
		serveOpts.Memfs, err = memfs.New(mfsopts)
		if err != nil {
			return fmt.Errorf(`%s: %w`, logp, err)
		}
	} else {
		serveOpts.Memfs.Opts.TryDirect = serveOpts.IsDevelopment
	}
	if staticfs != nil {
		staticfs.Opts.TryDirect = serveOpts.IsDevelopment
	}

	if serveOpts.HandleFS == nil {
		serveOpts.HandleFS = ciigo.onGet
	} else {
		userHandle := serveOpts.HandleFS
		serveOpts.HandleFS = ciigo.userHandleFS(userHandle)
	}

	ciigo.HTTPServer, err = libhttp.NewServer(serveOpts.ServerOptions)
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

	if serveOpts.IsDevelopment {
		pathHTMLTemplate = convertOpts.HTMLTemplate
	}

	ciigo.converter, err = NewConverter(pathHTMLTemplate)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	if !serveOpts.IsDevelopment {
		var tmplNode *memfs.Node

		tmplNode, _ = serveOpts.Memfs.Get(internalTemplatePath)
		if tmplNode != nil {
			ciigo.converter.tmpl, err = ciigo.converter.tmpl.Parse(string(tmplNode.Content))
			if err != nil {
				return fmt.Errorf(`%s: %s`, logp, err)
			}
		}
	}

	if serveOpts.IsDevelopment {
		ciigo.watcher, err = newWatcher(ciigo.converter, convertOpts)
		if err != nil {
			return fmt.Errorf(`%s: %w`, logp, err)
		}

		ciigo.converter.convertFileMarkups(ciigo.watcher.fileMarkups, false)
	}

	ciigo.serveOpts = serveOpts
	ciigo.convertOpts = convertOpts

	return nil
}

// userHandleFS wrap the user HandleFS with internal onGet.
func (ciigo *Ciigo) userHandleFS(userHandle libhttp.FSHandler) libhttp.FSHandler {
	return func(node *memfs.Node, res http.ResponseWriter, req *http.Request) (
		out *memfs.Node, statusCode int,
	) {
		out, statusCode = userHandle(node, res, req)
		if statusCode != 0 {
			return out, statusCode
		}
		return ciigo.onGet(node, res, req)
	}
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
	node *memfs.Node, w http.ResponseWriter, req *http.Request,
) (out *memfs.Node, statusCode int) {
	var (
		logp = `onGet`
		file = req.URL.Path
	)

	if node == nil {
		// File does not exist.
		ext := path.Ext(req.URL.Path)
		if ext == `.adoc` || ext == `.md` {
			// Redirect to .html.
			req.URL.Path = strings.TrimSuffix(req.URL.Path, ext) + `.html`
			req.URL.Path = path.Join(ciigo.serveOpts.BasePath, req.URL.Path)
			http.Redirect(w, req, req.URL.String(), http.StatusSeeOther)
			return nil, http.StatusSeeOther
		}
	} else if ciigo.serveOpts.IsDevelopment {
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
	if !ciigo.serveOpts.IsDevelopment {
		return node, 0
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
		return node, 0
	}
	var err error
	if isNew || ciigo.converter.shouldConvert(fmarkup, ciigo.serveOpts.IsDevelopment) {
		err = ciigo.converter.ToHTMLFile(fmarkup)
		if err != nil {
			log.Printf(`%s: failed to convert markup file %q: %s`,
				logp, fmarkup.path, err)
			return node, 0
		}
	}
	out, err = ciigo.serveOpts.Memfs.Get(file)
	if err != nil {
		log.Printf(`%s: failed to get %q: %s`, logp, file, err)
		return node, 0
	}
	return out, 0
}
