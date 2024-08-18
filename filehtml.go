// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"fmt"
	"html/template"
	"strings"

	"git.sr.ht/~shulhan/asciidoctor-go"
)

const (
	metadataStylesheet = `stylesheet`
	metadataTitle      = `title`
)

// fileHTML represent an HTML metadata for header and its body.
type fileHTML struct {
	Title       string
	EmbeddedCSS *template.CSS
	Styles      []string
	Body        template.HTML
	Metadata    map[string]string

	rawBody strings.Builder
}

func newFileHTML() (fhtml *fileHTML) {
	fhtml = &fileHTML{
		Metadata: map[string]string{},
	}
	return fhtml
}

func (fhtml *fileHTML) unpackAdocMetadata(doc *asciidoctor.Document) {
	var (
		k string
		v string
	)

	fhtml.Title = doc.Title.String()
	fhtml.Styles = fhtml.Styles[:0]

	for k, v = range doc.Attributes.Entry {
		switch k {
		case metadataStylesheet:
			fhtml.Styles = append(fhtml.Styles, v)
		case asciidoctor.DocAttrAuthorNames:
			fhtml.Metadata[asciidoctor.DocAttrAuthor] = v
		case asciidoctor.DocAttrDescription,
			asciidoctor.DocAttrGenerator,
			asciidoctor.DocAttrKeywords:
			fhtml.Metadata[k] = v
		}
	}

	if len(fhtml.Styles) == 0 {
		fhtml.EmbeddedCSS = embeddedCSS()
	}
}

func (fhtml *fileHTML) unpackMarkdownMetadata(metadata map[string]any) {
	var (
		key  string
		val  any
		vstr string
		ok   bool
	)

	fhtml.Styles = fhtml.Styles[:0]

	for key, val = range metadata {
		vstr, ok = val.(string)
		if !ok {
			vstr = fmt.Sprintf(`%s`, val)
		}

		key = strings.ToLower(key)
		switch key {
		case metadataStylesheet:
			fhtml.Styles = append(fhtml.Styles, vstr)
		case metadataTitle:
			fhtml.Title = vstr
		default:
			// Metadata `author_names`, `description`,
			// `generator`, and `keywords` goes here.
			fhtml.Metadata[key] = vstr
		}
	}

	if len(fhtml.Styles) == 0 {
		fhtml.EmbeddedCSS = embeddedCSS()
	}
}
