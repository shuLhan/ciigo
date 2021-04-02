// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//
// Package ciigo is a program to write static web server with embedded files
// using asciidoc markup languages.
//
// For more information see the README file at the page repository
// https://sr.ht/~shulhan/ciigo.
//
package ciigo

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/shuLhan/share/lib/memfs"
)

const (
	defAddress           = ":8080"
	defDir               = "."
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
// files using "htmlTemplate" file as template.
// If htmlTemplate is empty it will default to use embedded HTML template.
// See template_index_html.go for template format.
//
func Convert(opts *ConvertOptions) (err error) {
	logp := "Convert"

	if opts == nil {
		opts = &ConvertOptions{}
	}
	opts.init()

	htmlg, err := newHTMLGenerator(nil, opts.HtmlTemplate, true)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	fileMarkups, err := listFileMarkups(opts.Root)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	htmlg.convertFileMarkups(fileMarkups)

	return nil
}

//
// Generate a static Go file to be used for building binary.
//
// It will convert all markup files inside directory "dir" into HTML files,
// recursively; and read all the HTML files and files in "content/assets" and
// convert them into Go file in "out".
//
// If htmlTemplate is empty it will default to use embedded HTML template.
// See template_index_html.go for template format.
//
func Generate(opts *GenerateOptions) (err error) {
	logp := "Generate"

	if opts == nil {
		opts = &GenerateOptions{}
	}
	opts.init()

	htmlg, err := newHTMLGenerator(nil, opts.HtmlTemplate, true)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	fileMarkups, err := listFileMarkups(opts.Root)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	htmlg.convertFileMarkups(fileMarkups)

	memfsOpts := &memfs.Options{
		Root:     opts.Root,
		Excludes: defExcludes,
	}
	mfs, err := memfs.New(memfsOpts)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	if len(opts.HtmlTemplate) > 0 {
		_, err = mfs.AddFile(internalTemplatePath, opts.HtmlTemplate)
		if err != nil {
			return fmt.Errorf("%s: %w", logp, err)
		}
	}

	err = mfs.GoGenerate(opts.GenPackageName, opts.GenVarName,
		opts.GenGoFileName, memfs.EncodingGzip)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	return nil
}

//
// Serve the content at directory "dir" using HTTP server at specific
// "address".
//
func Serve(mfs *memfs.MemFS, dir, address, htmlTemplate string) (err error) {
	if len(dir) == 0 {
		dir = defDir
	}
	if len(address) == 0 {
		address = defAddress
	}
	srv, err := newServer(mfs, dir, address, htmlTemplate)
	if err != nil {
		return fmt.Errorf("Serve: %w", err)
	}
	err = srv.start()
	if err != nil {
		return fmt.Errorf("Serve: %w", err)
	}
	return nil
}

//
// Watch any changes on asciidoc files on directory "dir" recursively and
// changes on the HTML template file.
// If there is new or modified asciidoc files it will convert them into HTML
// files using HTML template automatically.
//
// If the HTML template file modified, it will re-convert all asciidoc files.
// If the HTML template file deleted, it will replace them with internal,
// default HTML template.
//
func Watch(dir, htmlTemplate string) (err error) {
	htmlg, err := newHTMLGenerator(nil, htmlTemplate, true)
	if err != nil {
		return fmt.Errorf("Watch: %w", err)
	}

	w, err := newWatcher(htmlg, dir)
	if err != nil {
		return fmt.Errorf("Watch: %w", err)
	}

	err = w.start()
	if err != nil {
		return fmt.Errorf("Watch: %w", err)
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
func listFileMarkups(dir string) (fileMarkups map[string]*fileMarkup, err error) {
	logp := "listFileMarkups"

	d, err := os.Open(dir)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", logp, err)
	}

	fis, err := d.Readdir(0)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", logp, err)
	}

	fileMarkups = make(map[string]*fileMarkup)

	for _, fi := range fis {
		name := fi.Name()

		if fi.IsDir() && name[0] != '.' {
			newdir := filepath.Join(dir, fi.Name())
			fmarkups, err := listFileMarkups(newdir)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", logp, err)
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

		filePath := filepath.Join(dir, name)

		fmarkup := &fileMarkup{
			path:     filePath,
			info:     fi,
			basePath: strings.TrimSuffix(filePath, ext),
			fhtml:    &fileHTML{},
		}

		fmarkup.fhtml.path = fmarkup.basePath + ".html"

		fileMarkups[filePath] = fmarkup
	}

	return fileMarkups, nil
}
