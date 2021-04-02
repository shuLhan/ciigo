// Copyright 2021, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/shuLhan/share/lib/clise"
	libio "github.com/shuLhan/share/lib/io"
	"github.com/shuLhan/share/lib/memfs"
)

//
// watcher watch for changes on all asciidoc file and convert them
// automatically.
//
type watcher struct {
	dir         string
	htmlg       *htmlGenerator
	dw          *libio.DirWatcher
	fileMarkups map[string]*fileMarkup
	changes     *clise.Clise
}

//
// newWatcher create a watcher that monitor every files changes in directory
// "dir" for new, modified, and deleted asciidoc files and HTML template file.
//
// The watcher depends on htmlGenerator to convert the markup to HTML using
// the HTML template in htmlGenerator.
//
//	watcher
//	|
//	+-- onChangeFileMarkup --> UPDATE --> htmlGenerator.convert()
//	|
//	+-- onChangeHtmlTemplate +--> DELETE --> htmlGenerator.htmlTemplateUseInternal()
//	                         |
//	                         +--> UPDATE --> htmlGenerated.htmlTemplateReload()
//
func newWatcher(htmlg *htmlGenerator, dir, exclude string) (w *watcher, err error) {
	w = &watcher{
		dir:         dir,
		htmlg:       htmlg,
		fileMarkups: make(map[string]*fileMarkup),
		changes:     clise.New(1),
	}
	w.dw = &libio.DirWatcher{
		Options: memfs.Options{
			Root: dir,
			Includes: []string{
				`.*\.adoc$`,
			},
			Excludes: []string{
				`^\..*`,
				`node_modules/.*`,
				`vendor/.*`,
			},
		},
		Delay:    time.Second,
		Callback: w.onChangeFileMarkup,
	}

	if len(exclude) > 0 {
		w.dw.Options.Excludes = append(w.dw.Options.Excludes, exclude)
	}

	return w, nil
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

	ext := strings.ToLower(filepath.Ext(ns.Node.SysPath))
	if !isExtensionMarkup(ext) {
		return
	}

	if ns.State == libio.FileStateDeleted {
		fmt.Printf("%s: %q deleted\n", logp, ns.Node.SysPath)
		fmarkup, ok := w.fileMarkups[ns.Node.SysPath]
		if ok {
			delete(w.fileMarkups, ns.Node.SysPath)
			w.changes.Push(fmarkup)
		}
		return
	}

	fmarkup := w.fileMarkups[ns.Node.SysPath]
	if fmarkup == nil {
		fmt.Printf("%s: %s created\n", logp, ns.Node.SysPath)
		fmarkup, err = newFileMarkup(ns.Node.SysPath, nil)
		if err != nil {
			log.Printf("%s: %s\n", logp, err)
			return
		}

		w.fileMarkups[ns.Node.SysPath] = fmarkup
	} else {
		fmt.Printf("%s: %s updated\n", logp, ns.Node.SysPath)
	}

	err = w.htmlg.convert(fmarkup)
	if err != nil {
		log.Printf("%s: %s\n", logp, err)
	}

	w.changes.Push(fmarkup)
}

//
// onChangeHtmlTemplate reload the HTML template and re-convert all markup
// files.
//
func (w *watcher) onChangeHtmlTemplate(ns *libio.NodeState) {
	var err error
	logp := "onChangeHtmlTemplate"

	if ns.State == libio.FileStateDeleted {
		log.Printf("%s: HTML template file %q has been deleted\n",
			logp, ns.Node.SysPath)
		err = w.htmlg.htmlTemplateUseInternal()
	} else {
		fmt.Printf("%s: recompiling HTML template %q ...\n", logp,
			ns.Node.SysPath)
		err = w.htmlg.htmlTemplateReload()
	}
	if err != nil {
		log.Printf("%s: %s\n", logp, err)
		return
	}

	fmt.Printf("%s: regenerate all markup files ...\n", logp)
	w.htmlg.convertFileMarkups(w.fileMarkups)
}

//
// start watching for changes.
//
func (w *watcher) start() (err error) {
	err = w.dw.Start()
	if err != nil {
		return fmt.Errorf("start: %w", err)
	}
	if len(w.htmlg.htmlTemplate) > 0 {
		_, err = libio.NewWatcher(w.htmlg.htmlTemplate, 0, w.onChangeHtmlTemplate)
		if err != nil {
			return fmt.Errorf("start: %w", err)
		}
	}
	return nil
}
