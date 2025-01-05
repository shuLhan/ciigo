// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"time"

	"git.sr.ht/~shulhan/asciidoctor-go"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

// Converter a single, reusable AsciiDoc converter.
type Converter struct {
	tmpl       *template.Template
	tmplSearch *template.Template

	// htmlTemplateModtime modification time for HTMLTemplate.
	htmlTemplateModtime time.Time

	htmlTemplate string // Path to HTML template in storage.
}

// NewConverter create and initialize Converter with HTML template.
// If htmlTemplate is empty, it will use the internal, predefined template.
func NewConverter(htmlTemplate string) (converter *Converter, err error) {
	var (
		logp = `NewConverter`

		tmplContent string
	)

	converter = &Converter{}

	converter.tmpl = template.New(``)

	if len(htmlTemplate) == 0 {
		tmplContent = templateIndexHTML
	} else {
		converter.htmlTemplate = filepath.Clean(htmlTemplate)

		var fi os.FileInfo

		fi, err = os.Stat(converter.htmlTemplate)
		if err != nil {
			return nil, fmt.Errorf(`%s: %w`, logp, err)
		}

		converter.htmlTemplateModtime = fi.ModTime()

		var bhtml []byte

		bhtml, err = os.ReadFile(converter.htmlTemplate)
		if err != nil {
			return nil, fmt.Errorf(`%s: %w`, logp, err)
		}

		tmplContent = string(bhtml)
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
		if !isForce {
			if !converter.shouldConvert(fmarkup) {
				continue
			}
		}

		err = converter.ToHTMLFile(fmarkup)
		if err != nil {
			log.Printf(`%s: %s`, logp, err)
		} else {
			log.Printf(`%s: converting %s`, logp, fmarkup.path)
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

// SetHTMLTemplateFile set the HTML template from file.
func (converter *Converter) SetHTMLTemplateFile(pathHTMLTemplate string) (err error) {
	converter.tmpl, err = template.ParseFiles(pathHTMLTemplate)
	if err != nil {
		return err
	}
	converter.htmlTemplate = pathHTMLTemplate
	return nil
}

// ToHTMLFile convert the AsciiDoc file to HTML.
func (converter *Converter) ToHTMLFile(fmarkup *FileMarkup) (err error) {
	var (
		logp = `ToHTMLFile`

		fhtml *fileHTML
		f     *os.File
	)

	switch fmarkup.kind {
	case markupKindAdoc:
		fhtml, err = converter.adocToHTML(fmarkup)
	case markupKindMarkdown:
		fhtml, err = converter.markdownToHTML(fmarkup)
	}
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	fhtml.Body = template.HTML(fhtml.rawBody.String())

	f, err = os.Create(fmarkup.pathHTML)
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

func (converter *Converter) adocToHTML(fmarkup *FileMarkup) (fhtml *fileHTML, err error) {
	var (
		logp = `adocToHTML`
		doc  *asciidoctor.Document
	)

	doc, err = asciidoctor.Open(fmarkup.path)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	fhtml = newFileHTML()

	err = doc.ToHTMLBody(&fhtml.rawBody)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	fhtml.unpackAdocMetadata(doc)

	return fhtml, nil
}

func (converter *Converter) markdownToHTML(fmarkup *FileMarkup) (fhtml *fileHTML, err error) {
	var (
		logp = `markdownToHTML`
		mdg  = goldmark.New(
			goldmark.WithExtensions(
				meta.Meta,
			),
		)

		in        []byte
		parserCtx parser.Context
	)

	in, err = os.ReadFile(fmarkup.path)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	fhtml = newFileHTML()
	parserCtx = parser.NewContext()

	err = mdg.Convert(in, &fhtml.rawBody, parser.WithContext(parserCtx))
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	fhtml.unpackMarkdownMetadata(meta.Get(parserCtx))

	return fhtml, nil
}

// shouldConvert will return true if the file markup fmarkup needs to be
// converted to HTML.
// It return true if the HTML file not exist or the template or markup file is
// newer than the HTML file.
func (converter *Converter) shouldConvert(fmarkup *FileMarkup) bool {
	var fi os.FileInfo
	fi, _ = os.Stat(fmarkup.pathHTML)
	if fi == nil {
		// HTML file may not exist.
		return true
	}

	var htmlModtime = fi.ModTime()
	var err error

	if len(converter.htmlTemplate) != 0 {
		fi, err = os.Stat(converter.htmlTemplate)
		if err != nil {
			// The template file may has been deleted.
			return true
		}

		if fi.ModTime().After(htmlModtime) {
			converter.htmlTemplateModtime = fi.ModTime()
			return true
		}
	}

	fi, err = os.Stat(fmarkup.path)
	if err != nil {
		// The markup file may has been deleted.
		return false
	}
	if fi.ModTime().After(htmlModtime) || fmarkup.info.Size() != fi.Size() {
		fmarkup.info = fi
		return true
	}
	return false
}
