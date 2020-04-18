// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"github.com/bytesparadise/libasciidoc"
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	mdparser "github.com/yuin/goldmark/parser"
)

//
// htmlGenerator provide a template to write full HTML file.
//
type htmlGenerator struct {
	path       string
	mdg        goldmark.Markdown
	tmpl       *template.Template
	tmplSearch *template.Template
}

func newHTMLGenerator(file, content string) (htmlg *htmlGenerator) {
	var err error

	htmlg = &htmlGenerator{
		path: file,
		mdg: goldmark.New(
			goldmark.WithExtensions(
				meta.Meta,
			),
		),
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

	in, err := ioutil.ReadFile(fmarkup.path)
	if err != nil {
		log.Fatal("htmlGenerator.convert: " + err.Error())
	}

	switch fmarkup.kind {
	case markupKindAsciidoc:
		cfg := configuration.NewConfiguration()
		bufin := bytes.NewBuffer(in)

		doc, err := parser.ParseDocument(bufin, cfg)
		if err != nil {
			log.Fatal(err)
		}

		bufin = bytes.NewBuffer(in)
		md, err := libasciidoc.Convert(bufin, &fhtml.rawBody, cfg)
		if err != nil {
			log.Fatal(err)
		}

		fhtml.unpackAdocMetadata(doc, md)

	case markupKindMarkdown:
		ctx := mdparser.NewContext()
		err := htmlg.mdg.Convert(in, &fhtml.rawBody, mdparser.WithContext(ctx))
		if err != nil {
			log.Fatal(err)
		}

		fmarkup.metadata = meta.Get(ctx)
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
