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
	fhtml := &fileHTML{}

	for _, fmarkup := range fileMarkups {
		fhtml.reset()
		fhtml.path = fmarkup.basePath + ".html"

		fmt.Printf("ciigo: converting %q to %q ... ", fmarkup.path, fhtml.path)

		htmlg.convert(fmarkup, fhtml, force)

		fmt.Println("OK")
		fmt.Printf("  metadata: %+v\n", fmarkup.metadata)
	}
}

func (htmlg *htmlGenerator) convert(fmarkup *fileMarkup, fhtml *fileHTML, force bool) {
	if fmarkup.isHTMLLatest(fhtml.path) && !force {
		return
	}

	switch fmarkup.kind {
	case markupKindAsciidoc:
		doc, err := asciidoctor.Open(fmarkup.path)
		if err != nil {
			log.Fatal(err)
		}

		err = doc.ToEmbeddedHTML(&fhtml.rawBody)
		if err != nil {
			log.Fatal(err)
		}

		fhtml.unpackAdocMetadata(doc)
	}
	if fhtml.rawBody.Len() == 0 {
		fmt.Println("skip")
		return
	}

	fhtml.unpackMarkup(fmarkup)
	htmlg.write(fhtml)
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
