// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
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

func (fhtml *fileHTML) unpackAdocMetadata(doc types.Document, md types.Metadata) {
	fhtml.Metadata = make(map[string]string)
	fhtml.Date = md.LastUpdated
	fhtml.Title = md.Title

	for k, v := range doc.Attributes {
		switch k {
		case metadataAuthor:
			fhtml.Author, _ = v.(string)
		case metadataDate:
			fhtml.Date, _ = v.(string)
		case metadataTitle:
			fhtml.Title, _ = v.(string)
		case metadataStylesheet:
			fhtml.Styles = append(fhtml.Styles, v.(string))
		default:
			fhtml.Metadata[k] = fmt.Sprintf("%v", v)
		}
	}
}
