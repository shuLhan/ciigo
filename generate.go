// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

//go:generate go run generate_main.go

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/shuLhan/share/lib/memfs"
)

//
// Generate a static Go file to be used for building binary.
//
// It will convert all asciidoc files inside root directory into HTML files,
// recursively; and read all the HTML files and files in "content/assets" and
// convert them into Go file in "out".
//
func Generate(root, out string) {
	htmlg := newHTMLGenerator()
	adocs := listAdocFiles(root)
	htmlg.convertAdocs(adocs, false)

	excs := []string{
		`.*\.adoc$`,
	}

	mfs, err := memfs.New(nil, excs, true)
	if err != nil {
		log.Fatal("ciigo: Generate: " + err.Error())
	}

	mfs.Mount(root)

	mfs.GoGenerate("", out)
}

//
// listAdocFiles find any ".adoc" file inside the content directory,
// recursively.
//
func listAdocFiles(dir string) (adocs []*fileAdoc) {
	d, err := os.Open(dir)
	if err != nil {
		log.Fatal("ciigo: listAdocFiles: os.Open: ", err)
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
		if fi.IsDir() {
			newdir := filepath.Join(dir, fi.Name())
			adocs = append(adocs, listAdocFiles(newdir)...)
			continue
		}

		ext := strings.ToLower(filepath.Ext(name))
		if ext != ".adoc" {
			continue
		}
		if fi.Size() == 0 {
			continue
		}

		adoc := &fileAdoc{
			path:     filepath.Join(dir, name),
			info:     fi,
			basePath: filepath.Join(dir, strings.TrimSuffix(name, ext)),
		}
		adocs = append(adocs, adoc)
	}

	return adocs
}
