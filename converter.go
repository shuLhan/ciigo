// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"git.sr.ht/~shulhan/asciidoctor-go"
)

// Converter a single, reusable AsciiDoc converter.
type Converter struct {
	tmpl         *template.Template
	tmplSearch   *template.Template
	htmlTemplate string
}

// NewConverter create and initialize Converter with HTML template.
// If htmlTemplate is empty, it will use the internal, predefined template.
func NewConverter(htmlTemplate string) (converter *Converter, err error) {
	var (
		logp = "NewConverter"

		tmplContent string
		bhtml       []byte
	)

	converter = &Converter{}

	converter.tmpl = template.New("")

	if len(htmlTemplate) == 0 {
		tmplContent = templateIndexHTML
	} else {
		converter.htmlTemplate = filepath.Clean(htmlTemplate)

		bhtml, err = os.ReadFile(converter.htmlTemplate)
		if err != nil {
			return nil, fmt.Errorf("%s: %s: %w", logp, converter.htmlTemplate, err)
		}

		tmplContent = string(bhtml)
	}

	converter.tmpl, err = converter.tmpl.Parse(tmplContent)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", logp, err)
	}

	converter.tmplSearch = template.New("search")
	converter.tmplSearch, err = converter.tmplSearch.Parse(templateSearch)
	if err != nil {
		return nil, fmt.Errorf("%s: %s: %w", logp, templateSearch, err)
	}

	return converter, nil
}

// convert the markup into HTML.
func (converter *Converter) convert(fmarkup *fileMarkup) (err error) {
	doc, err := asciidoctor.Open(fmarkup.path)
	if err != nil {
		return err
	}

	fmarkup.fhtml.rawBody.Reset()
	err = doc.ToHTMLBody(&fmarkup.fhtml.rawBody)
	if err != nil {
		return err
	}

	fmarkup.fhtml.unpackAdocMetadata(doc)

	return converter.write(fmarkup.fhtml)
}

// convertFileMarkups convert markup files into HTML.
func (converter *Converter) convertFileMarkups(fileMarkups map[string]*fileMarkup, isForce bool) {
	logp := "convertFileMarkups"
	for _, fmarkup := range fileMarkups {
		if !fmarkup.isNewerThanHtml() {
			if !isForce {
				continue
			}
		}

		err := converter.convert(fmarkup)
		if err != nil {
			fmt.Printf("%s: %s\n", logp, err)
		} else {
			fmt.Printf("%s: converting %s\n", logp, fmarkup.path)
		}
	}
}

func (converter *Converter) htmlTemplateReload() (err error) {
	converter.tmpl, err = template.ParseFiles(converter.htmlTemplate)
	if err != nil {
		return err
	}
	return nil
}

func (converter *Converter) htmlTemplateUseInternal() (err error) {
	converter.tmpl, err = converter.tmpl.Parse(templateIndexHTML)
	if err != nil {
		return err
	}
	return nil
}

// write the HTML file.
func (converter *Converter) write(fhtml *fileHtml) (err error) {
	f, err := os.Create(fhtml.path)
	if err != nil {
		return err
	}

	err = converter.tmpl.Execute(f, fhtml)
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}
