// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"html/template"
	"strings"

	"git.sr.ht/~shulhan/asciidoctor-go"
)

const (
	metadataStylesheet = `stylesheet`
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
