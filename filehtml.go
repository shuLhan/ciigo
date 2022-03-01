// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"html/template"
	"io/fs"
	"os"
	"strings"

	"git.sr.ht/~shulhan/asciidoctor-go"
)

//
// fileHtml represent an HTML metadata for header and its body.
//
type fileHtml struct {
	Title       string
	EmbeddedCSS *template.CSS
	Styles      []string
	Body        template.HTML
	Metadata    map[string]string

	path    string
	finfo   fs.FileInfo
	rawBody strings.Builder
}

func newFileHtml(path string) (fhtml *fileHtml) {
	fhtml = &fileHtml{
		path: path,
	}
	fhtml.finfo, _ = os.Stat(path)
	return fhtml
}

func (fhtml *fileHtml) unpackAdocMetadata(doc *asciidoctor.Document) {
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
