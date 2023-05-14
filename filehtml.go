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

// fileHtml represent an HTML metadata for header and its body.
type fileHtml struct {
	Title       string
	EmbeddedCSS *template.CSS
	Styles      []string
	Body        template.HTML
	Metadata    map[string]string

	rawBody strings.Builder
}

func newFileHtml() (fhtml *fileHtml) {
	fhtml = &fileHtml{
		Metadata: map[string]string{},
	}
	return fhtml
}

func (fhtml *fileHtml) unpackAdocMetadata(doc *asciidoctor.Document) {
	var (
		k string
		v string
	)

	fhtml.Title = doc.Title.String()
	fhtml.Styles = fhtml.Styles[:0]

	for k, v = range doc.Attributes {
		switch k {
		case metadataStylesheet:
			fhtml.Styles = append(fhtml.Styles, v)
		case asciidoctor.MetaNameAuthorNames:
			fhtml.Metadata[asciidoctor.MetaNameAuthor] = v
		case asciidoctor.MetaNameDescription,
			asciidoctor.MetaNameGenerator,
			asciidoctor.MetaNameKeywords:
			fhtml.Metadata[k] = v
		}
	}

	if len(fhtml.Styles) == 0 {
		fhtml.EmbeddedCSS = embeddedCSS()
	}
}

func (fhtml *fileHtml) unpackMarkdownMetadata(metadata map[string]any) {
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
