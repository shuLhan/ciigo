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
// watcher watch for changes on all markup files and convert them
// automatically to HTML.
//
type watcher struct {
	dir   string
	htmlg *htmlGenerator
	dw    *libio.DirWatcher

	// fileMarkups contains all markup files found inside "dir".
	// Its used to convert all markup files when the template file
	// changes.
	fileMarkups map[string]*fileMarkup

	changes *clise.Clise
}

//
// newWatcher create a watcher that monitor every files changes in directory
// "dir" for new, modified, and deleted markup files and HTML template file.
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
func newWatcher(htmlg *htmlGenerator, convertOpts *ConvertOptions) (w *watcher, err error) {
	logp := "newWatcher"

	w = &watcher{
		dir:     convertOpts.Root,
		htmlg:   htmlg,
		changes: clise.New(1),
	}
	w.dw = &libio.DirWatcher{
		Options: memfs.Options{
			Root: convertOpts.Root,
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

	if len(convertOpts.Exclude) > 0 {
		w.dw.Options.Excludes = append(w.dw.Options.Excludes, convertOpts.Exclude)
	}

	w.fileMarkups, err = listFileMarkups(convertOpts.Root, convertOpts.excRE)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", logp, err)
	}

	return w, nil
}

//
// onChangeFileMarkup watch the markup files inside the "content" directory,
// and re-generate them into HTML file when changed.
//
func (w *watcher) onChangeFileMarkup(ns *libio.NodeState) {
	var (
		logp    = "onChangeFileMarkup"
		fmarkup *fileMarkup
		err     error
	)

	ext := strings.ToLower(filepath.Ext(ns.Node.SysPath))
	if !isExtensionMarkup(ext) {
		return
	}

	switch ns.State {
	case libio.FileStateDeleted:
		fmt.Printf("%s: %q deleted\n", logp, ns.Node.SysPath)
		fmarkup, ok := w.fileMarkups[ns.Node.SysPath]
		if ok {
			delete(w.fileMarkups, ns.Node.SysPath)
			w.changes.Push(fmarkup)
		}
		return

	case libio.FileStateCreated:
		fmt.Printf("%s: %s created\n", logp, ns.Node.SysPath)
		fmarkup, err = newFileMarkup(ns.Node.SysPath, nil)
		if err != nil {
			log.Printf("%s: %s\n", logp, err)
			return
		}

		w.fileMarkups[ns.Node.SysPath] = fmarkup

	case libio.FileStateUpdateMode:
		fmt.Printf("%s: %s mode updated\n", logp, ns.Node.SysPath)
		return

	case libio.FileStateUpdateContent:
		fmt.Printf("%s: %s content updated\n", logp, ns.Node.SysPath)
		fmarkup = w.fileMarkups[ns.Node.SysPath]
		if fmarkup == nil {
			log.Printf("%s: %s not found\n", logp, ns.Node.SysPath)

			fmarkup, err = newFileMarkup(ns.Node.SysPath, nil)
			if err != nil {
				log.Printf("%s: %s\n", logp, err)
				return
			}

			w.fileMarkups[ns.Node.SysPath] = fmarkup
		}
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
