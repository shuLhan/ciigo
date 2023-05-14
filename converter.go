// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"git.sr.ht/~shulhan/asciidoctor-go"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

// Converter a single, reusable AsciiDoc converter.
type Converter struct {
	tmpl         *template.Template
	tmplSearch   *template.Template
	htmlTemplate string // Path to HTML template in storage.
}

// NewConverter create and initialize Converter with HTML template.
// If htmlTemplate is empty, it will use the internal, predefined template.
func NewConverter(htmlTemplate string) (converter *Converter, err error) {
	var (
		logp = `NewConverter`

		tmplContent string
		bhtml       []byte
	)

	converter = &Converter{}

	converter.tmpl = template.New(``)

	if len(htmlTemplate) == 0 {
		tmplContent = templateIndexHTML
	} else {
		converter.htmlTemplate = filepath.Clean(htmlTemplate)

		bhtml, err = os.ReadFile(converter.htmlTemplate)
		if err != nil {
			tmplContent = templateIndexHTML
		} else {
			tmplContent = string(bhtml)
		}
	}

	converter.tmpl, err = converter.tmpl.Parse(tmplContent)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	converter.tmplSearch = template.New(`search`)
	converter.tmplSearch, err = converter.tmplSearch.Parse(templateSearch)
	if err != nil {
		return nil, fmt.Errorf(`%s: %s: %w`, logp, templateSearch, err)
	}

	return converter, nil
}

// convertFileMarkups convert markup files into HTML.
func (converter *Converter) convertFileMarkups(fileMarkups map[string]*FileMarkup, isForce bool) {
	var (
		logp = `convertFileMarkups`

		fmarkup *FileMarkup
		err     error
	)

	for _, fmarkup = range fileMarkups {
		if !fmarkup.isNewerThanHtml() {
			if !isForce {
				continue
			}
		}

		err = converter.ToHtmlFile(fmarkup)
		if err != nil {
			log.Printf(`%s: %s`, logp, err)
		} else {
			fmt.Printf("%s: converting %s\n", logp, fmarkup.path)
		}
	}
}

func (converter *Converter) htmlTemplateUseInternal() (err error) {
	converter.tmpl, err = converter.tmpl.Parse(templateIndexHTML)
	if err != nil {
		return err
	}
	return nil
}

// SetHtmlTemplateFile set the HTML template from file.
func (converter *Converter) SetHtmlTemplateFile(pathHtmlTemplate string) (err error) {
	converter.tmpl, err = template.ParseFiles(pathHtmlTemplate)
	if err != nil {
		return err
	}
	converter.htmlTemplate = pathHtmlTemplate
	return nil
}

// ToHtmlFile convert the AsciiDoc file to HTML.
func (converter *Converter) ToHtmlFile(fmarkup *FileMarkup) (err error) {
	var (
		logp = `ToHtmlFile`

		fhtml *fileHtml
		f     *os.File
	)

	switch fmarkup.kind {
	case markupKindAdoc:
		fhtml, err = converter.adocToHtml(fmarkup)
	case markupKindMarkdown:
		fhtml, err = converter.markdownToHtml(fmarkup)
	}
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	fhtml.Body = template.HTML(fhtml.rawBody.String())

	f, err = os.Create(fmarkup.pathHtml)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	err = converter.tmpl.Execute(f, fhtml)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	return nil
}

func (converter *Converter) adocToHtml(fmarkup *FileMarkup) (fhtml *fileHtml, err error) {
	var (
		logp = `adocToHtml`
		doc  *asciidoctor.Document
	)

	doc, err = asciidoctor.Open(fmarkup.path)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	fhtml = newFileHtml()

	err = doc.ToHTMLBody(&fhtml.rawBody)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	fhtml.unpackAdocMetadata(doc)

	return fhtml, nil
}

func (converter *Converter) markdownToHtml(fmarkup *FileMarkup) (fhtml *fileHtml, err error) {
	var (
		logp = `markdownToHtml`
		mdg  = goldmark.New(
			goldmark.WithExtensions(
				meta.Meta,
			),
		)

		in        []byte
		parserCtx parser.Context
	)

	in, err = ioutil.ReadFile(fmarkup.path)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	fhtml = newFileHtml()
	parserCtx = parser.NewContext()

	err = mdg.Convert(in, &fhtml.rawBody, parser.WithContext(parserCtx))
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	fhtml.unpackMarkdownMetadata(meta.Get(parserCtx))

	return fhtml, nil
}
