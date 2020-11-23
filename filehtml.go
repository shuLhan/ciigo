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
	Author      string
	Date        string
	EmbeddedCSS *template.CSS
	Styles      []string
	Body        template.HTML
	Metadata    map[string]string

	path    string
	rawBody strings.Builder
}

//
// reset all fields to its empty values.
//
func (fhtml *fileHTML) reset() {
	fhtml.Title = ""
	fhtml.Author = ""
	fhtml.Date = ""
	fhtml.EmbeddedCSS = nil
	fhtml.Styles = fhtml.Styles[:0]
	fhtml.Body = template.HTML("")

	fhtml.path = ""
	fhtml.rawBody.Reset()
}

//
// unpackMarkup convert the markup metadata to its HTML representation and
// rawBody to template.HTML.
//
func (fhtml *fileHTML) unpackMarkup(fa *fileMarkup) {
	fhtml.Metadata = make(map[string]string)

	for k, v := range fa.metadata {
		switch k {
		case metadataAuthor:
			fhtml.Author = v.(string)
		case metadataDate:
			fhtml.Date = v.(string)
		case metadataTitle:
			fhtml.Title = v.(string)
		case metadataStylesheet:
			fhtml.Styles = append(fhtml.Styles, v.(string))
		default:
			fhtml.Metadata[k] = v.(string)
		}
	}
	if len(fhtml.Styles) == 0 {
		fhtml.EmbeddedCSS = embeddedCSS()
	}

	fhtml.Body = template.HTML(fhtml.rawBody.String()) // nolint:gosec
}

func (fhtml *fileHTML) unpackAdocMetadata(doc *asciidoctor.Document) {
	fhtml.Metadata = make(map[string]string)
	if len(doc.Revision.Date) > 0 {
		fhtml.Date = doc.Revision.Date
	} else {
		fhtml.Date = doc.LastUpdated
	}
	fhtml.Title = doc.Title.String()
	if len(doc.Authors) > 0 {
		fhtml.Author = doc.Authors[0].FullName()
	}

	for k, v := range doc.Attributes {
		switch k {
		case metadataStylesheet:
			fhtml.Styles = append(fhtml.Styles, v)
		default:
			fhtml.Metadata[k] = v
		}
	}
}
