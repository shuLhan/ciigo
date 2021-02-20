// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"git.sr.ht/~shulhan/asciidoctor-go"
)

//
// htmlGenerator provide a template to write full HTML file.
//
type htmlGenerator struct {
	path       string
	tmpl       *template.Template
	tmplSearch *template.Template
}

func newHTMLGenerator(file, content string) (htmlg *htmlGenerator) {
	var err error

	htmlg = &htmlGenerator{
		path: file,
	}

	htmlg.tmpl = template.New("")
	htmlg.tmpl, err = htmlg.tmpl.Parse(content)
	if err != nil {
		log.Fatal("newHTMLGenerator: ", err.Error())
	}

	htmlg.tmplSearch = template.New("search")
	htmlg.tmplSearch, err = htmlg.tmplSearch.Parse(templateSearch)
	if err != nil {
		log.Fatal("newHTMLGenerator: " + err.Error())
	}

	return
}

func (htmlg *htmlGenerator) reloadTemplate() (err error) {
	htmlg.tmpl, err = template.ParseFiles(htmlg.path)

	return
}

func (htmlg *htmlGenerator) convertFileMarkups(fileMarkups []*fileMarkup, force bool) {
	for _, fmarkup := range fileMarkups {
		fmt.Printf("ciigo: converting %q to %q ... ", fmarkup.path,
			fmarkup.fhtml.path)

		htmlg.convert(fmarkup, force)

		fmt.Println("OK")
		fmt.Printf("  metadata: %+v\n", fmarkup.metadata)
	}
}

func (htmlg *htmlGenerator) convert(fmarkup *fileMarkup, force bool) {
	if fmarkup.isHTMLLatest() && !force {
		return
	}

	doc, err := asciidoctor.Open(fmarkup.path)
	if err != nil {
		log.Fatal(err)
	}

	fmarkup.fhtml.rawBody.Reset()
	err = doc.ToHTMLBody(&fmarkup.fhtml.rawBody)
	if err != nil {
		log.Fatal(err)
	}

	fmarkup.fhtml.unpackAdocMetadata(doc)

	htmlg.write(fmarkup.fhtml)
}

//
// write the HTML file.
//
func (htmlg *htmlGenerator) write(fhtml *fileHTML) {
	f, err := os.Create(fhtml.path)
	if err != nil {
		log.Fatal("htmlGenerator: write: os.Create: " + err.Error())
	}

	err = htmlg.tmpl.Execute(f, fhtml)
	if err != nil {
		log.Fatal("htmlGenerator: write: Execute: " + err.Error())
	}

	err = f.Close()
	if err != nil {
		log.Fatal("htmlGenerator: write: Close: " + err.Error())
	}
}
