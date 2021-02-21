// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//
// Package ciigo is a program to write static web server with embedded files
// using asciidoc and markdown markup languages.
//
// For more information see the README file at the page repository
// https://sr.ht/~shulhan/ciigo.
//
package ciigo

import (
	"fmt"
	"io/ioutil"
	"log"
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
func Convert(dir, htmlTemplate string) {
	if len(dir) == 0 {
		dir = "."
	}

	contentHTML := templateIndexHTML

	if len(htmlTemplate) > 0 {
		b, err := ioutil.ReadFile(htmlTemplate)
		if err != nil {
			log.Fatal("ciigo.Convert: " + err.Error())
		}
		contentHTML = string(b)
	}

	htmlg := newHTMLGenerator(htmlTemplate, contentHTML)

	fileMarkups := listFileMarkups(dir)

	htmlg.convertFileMarkups(fileMarkups, true)
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
func Generate(opts *GenerateOptions) {
	if opts == nil {
		opts = &GenerateOptions{}
	}
	opts.init()

	contentHTML := templateIndexHTML

	if len(opts.HTMLTemplate) > 0 {
		b, err := ioutil.ReadFile(opts.HTMLTemplate)
		if err != nil {
			log.Fatal("ciigo.Generate: " + err.Error())
		}
		contentHTML = string(b)
	}

	htmlg := newHTMLGenerator(opts.HTMLTemplate, contentHTML)
	fileMarkups := listFileMarkups(opts.Root)

	htmlg.convertFileMarkups(fileMarkups, len(opts.HTMLTemplate) == 0)

	memfsOpts := &memfs.Options{
		Root:     opts.Root,
		Excludes: defExcludes,
	}
	mfs, err := memfs.New(memfsOpts)
	if err != nil {
		log.Fatal("ciigo.Generate: " + err.Error())
	}

	if len(opts.HTMLTemplate) > 0 {
		_, err = mfs.AddFile(internalTemplatePath, opts.HTMLTemplate)
		if err != nil {
			log.Fatalf("ciigo.Generate: AddFile %s: %s",
				opts.HTMLTemplate, err.Error())
		}
	}

	err = mfs.GoGenerate(opts.GenPackageName, opts.GenVarName,
		opts.GenGoFileName, memfs.EncodingGzip)
	if err != nil {
		log.Fatal("ciigo.Generate: " + err.Error())
	}
}

//
// Serve the content at directory "dir" using HTTP server at specific
// "address".
//
func Serve(mfs *memfs.MemFS, dir, address, htmlTemplate string) {
	if len(dir) == 0 {
		dir = defDir
	}
	if len(address) == 0 {
		address = defAddress
	}
	srv := newServer(mfs, dir, address, htmlTemplate)
	srv.start()
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
	w, err := newWatcher(dir, htmlTemplate)
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
func listFileMarkups(dir string) (fileMarkups []*fileMarkup) {
	d, err := os.Open(dir)
	if err != nil {
		log.Fatal("ciigo: listFileMarkups: os.Open: ", err)
	}

	fis, err := d.Readdir(0)
	if err != nil {
		log.Fatal("generate: " + err.Error())
	}

	for _, fi := range fis {
		name := fi.Name()

		if fi.IsDir() && name[0] != '.' {
			newdir := filepath.Join(dir, fi.Name())
			fileMarkups = append(fileMarkups, listFileMarkups(newdir)...)
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

		fileMarkups = append(fileMarkups, fmarkup)
	}

	return fileMarkups
}
