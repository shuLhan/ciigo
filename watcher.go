// Copyright 2021, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"git.sr.ht/~shulhan/asciidoctor-go"
	"github.com/shuLhan/share/lib/clise"
	libio "github.com/shuLhan/share/lib/io"
	"github.com/shuLhan/share/lib/memfs"
)

//
// watcher watch for changes on all asciidoc file and convert them
// automatically.
//
type watcher struct {
	dir          string
	htmlTemplate string
	dw           *libio.DirWatcher
	fileMarkups  map[string]*fileMarkup
	tmpl         *template.Template
	changes      *clise.Clise
}

//
// newWatcher create a watcher that monitor every files changes in directory
// "dir" for new, modified, and deleted asciidoc files and HTML template file.
//
func newWatcher(dir, htmlTemplate string) (w *watcher, err error) {
	w = &watcher{
		dir:          dir,
		htmlTemplate: htmlTemplate,
		fileMarkups:  make(map[string]*fileMarkup),
		changes:      clise.New(1),
	}
	w.dw = &libio.DirWatcher{
		Options: memfs.Options{
			Root: dir,
			Includes: []string{
				`.*\.adoc$`,
			},
		},
		Delay:    time.Second,
		Callback: w.onChangeFileMarkup,
	}

	err = w.initTemplate()
	if err != nil {
		return nil, fmt.Errorf("newWatcher: %w", err)
	}

	return w, nil
}

//
// convert the markup into HTML.
//
func (w *watcher) convert(fmarkup *fileMarkup) (err error) {
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

	return w.write(fmarkup.fhtml)
}

//
// convertFileMarkups convert markup files into HTML.
//
func (w *watcher) convertFileMarkups(fileMarkups map[string]*fileMarkup) {
	for _, fmarkup := range fileMarkups {
		fmt.Printf("ciigo: converting %q to %q => ", fmarkup.path,
			fmarkup.fhtml.path)

		err := w.convert(fmarkup)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("OK")
		}
	}
}

func (w *watcher) initTemplate() (err error) {
	logp := "initTemplate"

	w.tmpl = template.New("")

	if len(w.htmlTemplate) == 0 {
		w.tmpl, err = w.tmpl.Parse(templateIndexHTML)
		if err != nil {
			return fmt.Errorf("%s: %w", logp, err)
		}
		return nil
	}

	w.htmlTemplate = filepath.Clean(w.htmlTemplate)

	bhtml, err := ioutil.ReadFile(w.htmlTemplate)
	if err != nil {
		return fmt.Errorf("%s: %w", logp, err)
	}

	w.tmpl, err = w.tmpl.Parse(string(bhtml))
	if err != nil {
		return fmt.Errorf("%s: %s", logp, err)
	}

	return nil
}

//
// onChangeFileMarkup watch the markup files inside the "content" directory,
// and re-generate them into HTML file when changed.
//
func (w *watcher) onChangeFileMarkup(ns *libio.NodeState) {
	var (
		logp = "onChangeFileMarkup"
		err  error
	)

	if ns.State == libio.FileStateDeleted {
		fmt.Printf("ciigo: %s: %q deleted\n", logp, ns.Node.SysPath)
		fmarkup, ok := w.fileMarkups[ns.Node.SysPath]
		if ok {
			delete(w.fileMarkups, ns.Node.SysPath)
			w.changes.Push(fmarkup)
		}
		return
	}

	ext := strings.ToLower(filepath.Ext(ns.Node.SysPath))
	if !isExtensionMarkup(ext) {
		return
	}

	fmarkup := w.fileMarkups[ns.Node.SysPath]
	if fmarkup == nil {
		fmt.Printf("ciigo: %s: %s created\n", logp, ns.Node.SysPath)
		fmarkup, err = newFileMarkup(ns.Node.SysPath, nil)
		if err != nil {
			log.Printf("%s: %s\n", logp, err)
			return
		}

		w.fileMarkups[ns.Node.SysPath] = fmarkup
	} else {
		fmt.Printf("ciigo: %s: %s updated\n", logp, ns.Node.SysPath)
	}

	err = w.convert(fmarkup)
	if err != nil {
		log.Printf("%s: %s\n", logp, err)
	}

	w.changes.Push(fmarkup)
}

//
// onChangeHTMLTemplate reload the HTML template and re-convert all markup
// files.
//
func (w *watcher) onChangeHTMLTemplate(ns *libio.NodeState) {
	var err error
	logp := "onChangeHTMLTemplate"

	if ns.State == libio.FileStateDeleted {
		log.Printf("ciigo: HTML template file %q has been deleted\n",
			ns.Node.SysPath)
		// Use the internal HTML template.
		w.tmpl, err = w.tmpl.Parse(templateIndexHTML)
		if err != nil {
			log.Printf("%s: %s", logp, err)
			return
		}
	} else {
		fmt.Printf("ciigo: recompiling HTML template %q ...\n", w.htmlTemplate)
		w.tmpl, err = template.ParseFiles(w.htmlTemplate)
		if err != nil {
			log.Printf("%s: %s\n", logp, err)
			return
		}
	}

	fmt.Printf("ciigo: regenerate all markup files ...\n")
	w.convertFileMarkups(w.fileMarkups)
}

//
// start watching for changes.
//
func (w *watcher) start() (err error) {
	err = w.dw.Start()
	if err != nil {
		return fmt.Errorf("start: %w", err)
	}
	if len(w.htmlTemplate) > 0 {
		_, err = libio.NewWatcher(w.htmlTemplate, 0, w.onChangeHTMLTemplate)
		if err != nil {
			return fmt.Errorf("start: %w", err)
		}
	}
	return nil
}

//
// write the HTML file.
//
func (w *watcher) write(fhtml *fileHTML) (err error) {
	f, err := os.Create(fhtml.path)
	if err != nil {
		return err
	}

	err = w.tmpl.Execute(f, fhtml)
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}
