// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

import (
	"html/template"
	"strings"

	"git.sr.ht/~shulhan/asciidoctor-go"
)

//
// fileHTML represent an HTML metadata for header and its body.
//
type fileHTML struct {
	Title       string
	EmbeddedCSS *template.CSS
	Styles      []string
	Body        template.HTML
	Metadata    map[string]string

	path    string
	rawBody strings.Builder
}

func (fhtml *fileHTML) unpackAdocMetadata(doc *asciidoctor.Document) {
	fhtml.Title = doc.Title.String()
	fhtml.Styles = fhtml.Styles[:0]
	fhtml.Metadata = make(map[string]string, len(doc.Attributes))

	for k, v := range doc.Attributes {
		switch k {
		case metadataStylesheet:
			fhtml.Styles = append(fhtml.Styles, v)
		default:
			fhtml.Metadata[k] = v
		}
	}

	if len(fhtml.Styles) == 0 {
		fhtml.EmbeddedCSS = embeddedCSS()
	}

	fhtml.Body = template.HTML(fhtml.rawBody.String())
}
