// SPDX-License-Identifier: GPL-3.0-or-later
// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>

package ciigo

import (
	"fmt"
	"html/template"
	"log"
	"strings"

	"kilabit.info/asciidoctor-go"
)

const (
	metadataTitle = `title`
)

// fileHTML represent an HTML metadata for header and its body.
type fileHTML struct {
	Title       string
	EmbeddedCSS template.CSS
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

	// By default `:stylesheet:` is set in Attributes.
	// If user remove it, by using `:stylesheet!:`, the key would not
	// exist in the map.
	var withStylesheet bool

	fhtml.Title = doc.Title.String()
	fhtml.Styles = fhtml.Styles[:0]

	for k, v = range doc.Attributes.Entry {
		switch k {
		case asciidoctor.DocAttrStylesheet:
			vals := strings.Split(v, ",")
			for _, v = range vals {
				v = strings.TrimSpace(v)
				if v == `` || v == `default` {
					withStylesheet = true
				} else {
					fhtml.Styles = append(fhtml.Styles, v)
				}
			}
		case asciidoctor.DocAttrAuthorNames:
			fhtml.Metadata[asciidoctor.DocAttrAuthor] = v
		default:
			fhtml.Metadata[k] = v
		}
	}
	fhtml.initCSS(withStylesheet)
}

func (fhtml *fileHTML) unpackMarkdownMetadata(metadata map[string]any) {
	var (
		key  string
		val  any
		vstr string
		ok   bool
	)
	// On markdown, the default stylesheet is true.
	// User need to explicitly set `stylesheet: false`
	// on metadata to turn it off.
	var withStylesheet = true

	fhtml.Styles = fhtml.Styles[:0]

	for key, val = range metadata {
		vstr, ok = val.(string)
		if !ok {
			vstr = fmt.Sprintf(`%s`, val)
		}

		key = strings.ToLower(key)
		switch key {
		case asciidoctor.DocAttrStylesheet:
			vals := strings.Split(vstr, ",")
			for _, v := range vals {
				v = strings.TrimSpace(v)

				switch v {
				case ``:
					// Skip it.
				case `false`:
					withStylesheet = false
				case `default`:
					// Use embedded stylesheet.
				default:
					fhtml.Styles = append(fhtml.Styles, vstr)
				}
			}
		case metadataTitle:
			fhtml.Title = vstr
		default:
			// Metadata `author_names`, `description`,
			// `generator`, and `keywords` goes here.
			fhtml.Metadata[key] = vstr
		}
	}
	fhtml.initCSS(withStylesheet)
}

func (fhtml *fileHTML) initCSS(withStylesheet bool) {
	var logp = `initCSS`

	if !withStylesheet {
		// User explicitly unset the default stylesheet.
		return
	}
	css, err := staticfs.Get(`/ciigo.css`)
	if err != nil {
		log.Printf(`%s: %s`, logp, err)
		return
	}
	fhtml.EmbeddedCSS = template.CSS(string(css.Content))
}
