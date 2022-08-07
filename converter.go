// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"git.sr.ht/~shulhan/asciidoctor-go"
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
func (converter *Converter) convertFileMarkups(fileMarkups map[string]*fileMarkup, isForce bool) {
	var (
		logp = `convertFileMarkups`

		fmarkup *fileMarkup
		err     error
	)

	for _, fmarkup = range fileMarkups {
		if !fmarkup.isNewerThanHtml() {
			if !isForce {
				continue
			}
		}

		err = converter.ToHtmlFile(fmarkup.path, fmarkup.pathHtml)
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
func (converter *Converter) ToHtmlFile(pathAdoc, pathHtml string) (err error) {
	var (
		logp  = `ToHtmlFile`
		fhtml = newFileHtml()

		htmlBody string
		doc      *asciidoctor.Document
		f        *os.File
	)

	doc, err = asciidoctor.Open(pathAdoc)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	err = doc.ToHTMLBody(&fhtml.rawBody)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	fhtml.unpackAdocMetadata(doc)

	htmlBody = fhtml.rawBody.String()
	fhtml.Body = template.HTML(htmlBody)

	f, err = os.Create(pathHtml)
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
