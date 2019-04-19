// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

//
// htmlGenerator provide a template to write full HTML file.
//
type htmlGenerator struct {
	path string
	tmpl *template.Template
}

func newHTMLGenerator() (htmlg *htmlGenerator) {
	htmlg = &htmlGenerator{
		path: "./templates/html.tmpl",
	}

	err := htmlg.loadTemplate()
	if err != nil {
		log.Fatal("htmlGenerator: loadTemplate: ", err.Error())
	}

	return
}

func (htmlg *htmlGenerator) loadTemplate() (err error) {
	htmlg.tmpl, err = template.ParseFiles(htmlg.path)

	return
}

func (htmlg *htmlGenerator) convertAdocs(adocs []*fileAdoc, force bool) {
	fhtml := &fileHTML{}

	for _, adoc := range adocs {
		fhtml.reset()
		fhtml.path = adoc.basePath + ".html"

		fmt.Printf("ciigo: converting %q to %q ... ", adoc.path, fhtml.path)

		adoc.toHTML(fhtml.path, &fhtml.rawBody, force)
		if fhtml.rawBody.Len() == 0 {
			fmt.Println("skip")
			continue
		}

		fhtml.unpackAdoc(adoc)
		fhtml.Body = template.HTML(fhtml.rawBody.String()) //nolint: gosec

		htmlg.write(fhtml)

		fmt.Println("OK")
		fmt.Printf("  metadata: %+v\n", adoc.metadata)
	}
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
