// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//go:generate go run generate_main.go

package ciigo

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/shuLhan/share/lib/memfs"
)

//
// Convert all markup files inside directory, recursively, into HTML files
// using "htmlTemplate" file as template.
//
func Convert(dir, htmlTemplate string) {
	if len(dir) == 0 {
		dir = "."
	}

	b, err := ioutil.ReadFile(htmlTemplate)
	if err != nil {
		log.Fatal("ciigo.Convert: " + err.Error())
	}

	htmlg := newHTMLGenerator(htmlTemplate, string(b))

	fileMarkups := listFileMarkups(dir)

	htmlg.convertFileMarkups(fileMarkups, true)
}

//
// Generate a static Go file to be used for building binary.
//
// It will convert all markup files inside root directory into HTML files,
// recursively; and read all the HTML files and files in "content/assets" and
// convert them into Go file in "out".
//
func Generate(root, out, htmlTemplate string) {
	b, err := ioutil.ReadFile(htmlTemplate)
	if err != nil {
		log.Fatal("ciigo.Generate: " + err.Error())
	}

	htmlg := newHTMLGenerator(htmlTemplate, string(b))
	fileMarkups := listFileMarkups(root)

	htmlg.convertFileMarkups(fileMarkups, false)

	mfs, err := memfs.New(nil, defExcludes, true)
	if err != nil {
		log.Fatal("ciigo.Generate: " + err.Error())
	}

	err = mfs.Mount(root)
	if err != nil {
		log.Fatalf("ciigo.Generate: Mount %s: %s", root, err.Error())
	}

	_, err = mfs.AddFile(htmlTemplate)
	if err != nil {
		log.Fatalf("ciigo.Generate: AddFile %s: %s", htmlTemplate, err.Error())
	}

	err = mfs.GoGenerate("", out, memfs.EncodingGzip)
	if err != nil {
		log.Fatal("ciigo.Generate: " + err.Error())
	}
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
