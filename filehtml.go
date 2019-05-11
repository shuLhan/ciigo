// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

import (
	"html/template"
	"strings"
)

//
// fileHTML represent an HTML metadata for header and its body.
//
type fileHTML struct {
	Title    string
	Styles   []string
	Body     template.HTML
	Metadata map[string]string

	path    string
	rawBody strings.Builder
}

//
// reset all fields to its empty values.
//
func (fhtml *fileHTML) reset() {
	fhtml.Title = ""
	fhtml.Styles = fhtml.Styles[:0]
	fhtml.Body = template.HTML("")

	fhtml.path = ""
	fhtml.rawBody.Reset()
}

//
// unpackAdoc convert the asciidoc metadata to its HTML representation and
// rawBody to template.HTML.
//
func (fhtml *fileHTML) unpackAdoc(fa *fileAdoc) {
	fhtml.Metadata = make(map[string]string)

	for k, v := range fa.metadata {
		switch k {
		case "title":
			fhtml.Title = v.(string)
		case "stylesheet":
			fhtml.Styles = append(fhtml.Styles, v.(string))
		default:
			fhtml.Metadata[k] = v.(string)
		}
	}

	fhtml.Body = template.HTML(fhtml.rawBody.String()) // nolint:gosec
}
