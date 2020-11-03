// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//
// Package ciigo is a program to write static web server with embedded files
// using asciidoc and markdown markup languages.
//
// For more information see the README file at the page repository
// https://github.com/shuLhan/ciigo.
//
package ciigo

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/shuLhan/share/lib/memfs"
)

const (
	defAddress  = ":8080"
	defDir      = "."
	dirAssets   = "assets"
	extAsciidoc = ".adoc"
)

const (
	metadataAuthor     = "author"
	metadataDate       = "date"
	metadataStylesheet = "stylesheet"
	metadataTitle      = "title"
)

const (
	markupKindUnknown byte = iota
	markupKindAsciidoc
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
func Generate(dir, out, htmlTemplate string) {
	contentHTML := templateIndexHTML

	if len(htmlTemplate) > 0 {
		b, err := ioutil.ReadFile(htmlTemplate)
		if err != nil {
			log.Fatal("ciigo.Generate: " + err.Error())
		}
		contentHTML = string(b)
	}

	htmlg := newHTMLGenerator(htmlTemplate, contentHTML)
	fileMarkups := listFileMarkups(dir)

	htmlg.convertFileMarkups(fileMarkups, len(htmlTemplate) == 0)

	mfs, err := memfs.New(dir, nil, defExcludes, true)
	if err != nil {
		log.Fatal("ciigo.Generate: " + err.Error())
	}

	if len(htmlTemplate) > 0 {
		_, err = mfs.AddFile(htmlTemplate)
		if err != nil {
			log.Fatalf("ciigo.Generate: AddFile %s: %s", htmlTemplate, err.Error())
		}
	}

	err = mfs.GoGenerate("", out, memfs.EncodingGzip)
	if err != nil {
		log.Fatal("ciigo.Generate: " + err.Error())
	}
}

//
// Serve the content at directory "dir" using HTTP server at specific
// "address".
//
func Serve(dir, address, htmlTemplate string) {
	if len(dir) == 0 {
		dir = defDir
	}
	if len(address) == 0 {
		address = defAddress
	}
	srv := newServer(dir, address, htmlTemplate)
	srv.start()
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

		if name == dirAssets {
			continue
		}
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

		markupf := &fileMarkup{
			kind:     markupKind(ext),
			path:     filepath.Join(dir, name),
			info:     fi,
			basePath: filepath.Join(dir, strings.TrimSuffix(name, ext)),
		}
		fileMarkups = append(fileMarkups, markupf)
	}

	return fileMarkups
}

func markupKind(ext string) byte {
	switch ext {
	case extAsciidoc:
		return markupKindAsciidoc
	}
	return markupKindUnknown
}
