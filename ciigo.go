// SPDX-FileCopyrightText: 2019 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

//
// Package ciigo is a program to write static web server with embedded files
// using the asciidoc markup languages.
//
// For more information see the README file at the page repository
// https://sr.ht/~shulhan/ciigo.
//
package ciigo

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/shuLhan/share/lib/memfs"
)

const (
	extAsciidoc          = ".adoc"
	internalTemplatePath = "_internal/.template"
)

const (
	metadataStylesheet = "stylesheet"
)

//nolint: gochecknoglobals
var (
	defExcludes = []string{
		`.*\.adoc$`,
		`^\..*`,
	}
)

//
// Convert all markup files inside directory "dir" recursively into HTML
// files using ConvertOptions HtmlTemplate file as base template.
// If HtmlTemplate is empty it will default to use embedded HTML template.
// See template_index_html.go for template format.
//
func Convert(opts *ConvertOptions) (err error) {
	var (
		logp        = "Convert"
		htmlg       *htmlGenerator
		fileMarkups map[string]*fileMarkup
	)

	if opts == nil {
		opts = &ConvertOptions{}
	}
	err = opts.init()
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	htmlg, err = newHTMLGenerator(nil, opts.HtmlTemplate, true)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	fileMarkups, err = listFileMarkups(opts.Root, opts.excRE)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	htmlg.convertFileMarkups(fileMarkups, false)

	return nil
}

//
// GoEmbed generate a static Go file that embed all files inside Root except
// the one that being excluded explicitly by ConvertOptions Exclude.
//
// It convert all markup files inside directory "dir" into HTML files,
// recursively, and then embed them into Go file defined by
// EmbedOptions.GoFileName.
//
// If HtmlTemplate option is empty it default to use embedded HTML
// template.
// See template_index_html.go for template format.
//
func GoEmbed(opts *EmbedOptions) (err error) {
	var (
		logp        = "GoEmbed"
		htmlg       *htmlGenerator
		fileMarkups map[string]*fileMarkup
		mfs         *memfs.MemFS
	)

	if opts == nil {
		opts = &EmbedOptions{}
	}
	err = opts.init()
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	htmlg, err = newHTMLGenerator(nil, opts.HtmlTemplate, true)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	fileMarkups, err = listFileMarkups(opts.Root, opts.excRE)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	htmlg.convertFileMarkups(fileMarkups, false)

	memfsOpts := &memfs.Options{
		Root:     opts.Root,
		Excludes: defExcludes,
		Embed:    opts.EmbedOptions,
	}

	mfs, err = memfs.New(memfsOpts)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	if len(opts.HtmlTemplate) > 0 {
		_, err = mfs.AddFile(internalTemplatePath, opts.HtmlTemplate)
		if err != nil {
			return fmt.Errorf("%s: %w", logp, err)
		}
	}

	err = mfs.GoEmbed()
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	return nil
}

//
// Serve the content at directory "dir" using HTTP server at specific
// "address".
//
func Serve(opts *ServeOptions) (err error) {
	var (
		logp = "Serve"
		srv  *server
	)

	if opts == nil {
		opts = &ServeOptions{}
	}
	err = opts.init()
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	srv, err = newServer(opts)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}
	err = srv.start()
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}
	return nil
}

//
// Watch any changes on markup files on directory Root recursively and
// changes on the HTML template file.
// If there is new or modified markup files it will convert them into HTML
// files using HTML template automatically.
//
// If the HTML template file modified, it will re-convert all markup files.
// If the HTML template file deleted, it will replace them with internal,
// default HTML template.
//
func Watch(opts *ConvertOptions) (err error) {
	var (
		logp  = "Watch"
		htmlg *htmlGenerator
		w     *watcher
	)

	if opts == nil {
		opts = &ConvertOptions{}
	}
	err = opts.init()
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	htmlg, err = newHTMLGenerator(nil, opts.HtmlTemplate, true)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	w, err = newWatcher(htmlg, opts)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	err = w.start()
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	return nil
}

func isExtensionMarkup(ext string) bool {
	return ext == extAsciidoc
}

//
// listFileMarkups find any markup files inside the content directory,
// recursively.
//
func listFileMarkups(dir string, excRE []*regexp.Regexp) (
	fileMarkups map[string]*fileMarkup, err error,
) {
	var (
		logp = "listFileMarkups"
		d    *os.File
		fis  []os.FileInfo
	)

	d, err = os.Open(dir)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", logp, err)
	}

	fis, err = d.Readdir(0)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", logp, err)
	}

	fileMarkups = make(map[string]*fileMarkup)

	for _, fi := range fis {
		name := fi.Name()
		filePath := filepath.Join(dir, name)

		if fi.IsDir() {
			if name[0] == '.' {
				// Skip any directory start with '.'.
				continue
			}
			fmarkups, err := listFileMarkups(filePath, excRE)
			if err != nil {
				return nil, fmt.Errorf("%s: %s: %w", logp, filePath, err)
			}
			for k, v := range fmarkups {
				fileMarkups[k] = v
			}
			continue
		}

		ext := strings.ToLower(filepath.Ext(name))
		if !isExtensionMarkup(ext) {
			continue
		}
		if fi.Size() == 0 {
			continue
		}
		if isExcluded(filePath, excRE) {
			continue
		}
		fmarkup, err := newFileMarkup(filePath, fi)
		if err != nil {
			return nil, fmt.Errorf("%s: %s: %w", logp, filePath, err)
		}
		fileMarkups[filePath] = fmarkup
	}

	return fileMarkups, nil
}

func isExcluded(path string, excs []*regexp.Regexp) bool {
	if len(excs) == 0 {
		return false
	}
	for _, re := range excs {
		if re.MatchString(path) {
			return true
		}
	}
	return false
}
