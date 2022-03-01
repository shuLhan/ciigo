// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

	"git.sr.ht/~shulhan/asciidoctor-go"
	"github.com/shuLhan/share/lib/memfs"
)

//
// htmlGenerator provide a template to write full HTML file.
//
type htmlGenerator struct {
	tmpl         *template.Template
	tmplSearch   *template.Template
	htmlTemplate string
}

func newHTMLGenerator(mfs *memfs.MemFS, htmlTemplate string, devel bool) (
	htmlg *htmlGenerator, err error,
) {
	var logp = "newHTMLGenerator"

	htmlg = &htmlGenerator{}

	htmlg.tmpl = template.New("")

	if len(htmlTemplate) == 0 {
		htmlg.tmpl, err = htmlg.tmpl.Parse(templateIndexHTML)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", logp, err)
		}
	} else if mfs == nil || devel {
		htmlg.htmlTemplate = filepath.Clean(htmlTemplate)

		bhtml, err := ioutil.ReadFile(htmlg.htmlTemplate)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", logp, err)
		}

		htmlg.tmpl, err = htmlg.tmpl.Parse(string(bhtml))
		if err != nil {
			return nil, fmt.Errorf("%s: %s", logp, err)
		}
	} else {
		// Load HTML template from memory file system.
		tmplNode, err := mfs.Get(internalTemplatePath)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", logp, err)
		}

		htmlg.tmpl, err = htmlg.tmpl.Parse(string(tmplNode.Content))
		if err != nil {
			return nil, fmt.Errorf("%s: %s", logp, err)
		}
	}

	htmlg.tmplSearch = template.New("search")
	htmlg.tmplSearch, err = htmlg.tmplSearch.Parse(templateSearch)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", logp, err)
	}

	return htmlg, nil
}

//
// convert the markup into HTML.
//
func (htmlg *htmlGenerator) convert(fmarkup *fileMarkup) (err error) {
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

	return htmlg.write(fmarkup.fhtml)
}

//
// convertFileMarkups convert markup files into HTML.
//
func (htmlg *htmlGenerator) convertFileMarkups(fileMarkups map[string]*fileMarkup, isForce bool) {
	logp := "convertFileMarkups"
	for _, fmarkup := range fileMarkups {
		if !fmarkup.isNewerThanHtml() {
			if !isForce {
				continue
			}
		}

		err := htmlg.convert(fmarkup)
		if err != nil {
			fmt.Printf("%s: %s\n", logp, err)
		} else {
			fmt.Printf("%s: converting %s\n", logp, fmarkup.path)
		}
	}
}

func (htmlg *htmlGenerator) htmlTemplateReload() (err error) {
	htmlg.tmpl, err = template.ParseFiles(htmlg.htmlTemplate)
	if err != nil {
		return err
	}
	return nil
}

func (htmlg *htmlGenerator) htmlTemplateUseInternal() (err error) {
	htmlg.tmpl, err = htmlg.tmpl.Parse(templateIndexHTML)
	if err != nil {
		return err
	}
	return nil
}

//
// write the HTML file.
//
func (htmlg *htmlGenerator) write(fhtml *fileHtml) (err error) {
	f, err := os.Create(fhtml.path)
	if err != nil {
		return err
	}

	err = htmlg.tmpl.Execute(f, fhtml)
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}
