// SPDX-FileCopyrightText: 2021 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/shuLhan/share/lib/clise"
	"github.com/shuLhan/share/lib/memfs"
)

//
// watcher watch for changes on all markup files and convert them
// automatically to HTML.
//
type watcher struct {
	changes       *clise.Clise
	watchDir      *memfs.DirWatcher
	watchTemplate *memfs.Watcher
	htmlg         *htmlGenerator

	// fileMarkups contains all markup files found inside "dir".
	// Its used to convert all markup files when the template file
	// changes.
	fileMarkups map[string]*fileMarkup

	dir string
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
//	+-- watchFileMarkup --> UPDATE --> htmlGenerator.convert()
//	|
//	+-- watchHtmlTemplate +--> DELETE --> htmlGenerator.htmlTemplateUseInternal()
//	                      |
//	                      +--> UPDATE --> htmlGenerated.htmlTemplateReload()
//
func newWatcher(htmlg *htmlGenerator, convertOpts *ConvertOptions) (w *watcher, err error) {
	var (
		logp = "newWatcher"
	)

	w = &watcher{
		dir:     convertOpts.Root,
		htmlg:   htmlg,
		changes: clise.New(1),
	}
	w.watchDir = &memfs.DirWatcher{
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
		Delay: time.Second,
	}

	if len(convertOpts.Exclude) > 0 {
		w.watchDir.Options.Excludes = append(w.watchDir.Options.Excludes, convertOpts.Exclude)
	}

	w.fileMarkups, err = listFileMarkups(convertOpts.Root, convertOpts.excRE)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", logp, err)
	}

	return w, nil
}

//
// start watching for changes.
//
func (w *watcher) start() (err error) {
	err = w.watchDir.Start()
	if err != nil {
		return fmt.Errorf("start: %w", err)
	}

	go w.watchFileMarkup()

	if len(w.htmlg.htmlTemplate) > 0 {
		w.watchTemplate, err = memfs.NewWatcher(w.htmlg.htmlTemplate, 0)
		if err != nil {
			return fmt.Errorf("start: %w", err)
		}
		go w.watchHtmlTemplate()
	}
	return nil
}

//
// watchFileMarkup watch the markup files inside the "content" directory,
// and re-generate them into HTML file when changed.
//
func (w *watcher) watchFileMarkup() {
	var (
		logp = "watchFileMarkup"

		ns      memfs.NodeState
		fmarkup *fileMarkup
		err     error
	)

	for ns = range w.watchDir.C {
		ext := strings.ToLower(filepath.Ext(ns.Node.SysPath))
		if !isExtensionMarkup(ext) {
			continue
		}

		switch ns.State {
		case memfs.FileStateDeleted:
			fmt.Printf("%s: %q deleted\n", logp, ns.Node.SysPath)
			fmarkup, ok := w.fileMarkups[ns.Node.SysPath]
			if ok {
				delete(w.fileMarkups, ns.Node.SysPath)
				w.changes.Push(fmarkup)
			}
			continue

		case memfs.FileStateCreated:
			fmt.Printf("%s: %s created\n", logp, ns.Node.SysPath)
			fmarkup, err = newFileMarkup(ns.Node.SysPath, nil)
			if err != nil {
				log.Printf("%s: %s\n", logp, err)
				continue
			}

			w.fileMarkups[ns.Node.SysPath] = fmarkup

		case memfs.FileStateUpdateMode:
			fmt.Printf("%s: %s mode updated\n", logp, ns.Node.SysPath)
			continue

		case memfs.FileStateUpdateContent:
			fmt.Printf("%s: %s content updated\n", logp, ns.Node.SysPath)
			fmarkup = w.fileMarkups[ns.Node.SysPath]
			if fmarkup == nil {
				log.Printf("%s: %s not found\n", logp, ns.Node.SysPath)

				fmarkup, err = newFileMarkup(ns.Node.SysPath, nil)
				if err != nil {
					log.Printf("%s: %s\n", logp, err)
					continue
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
}

//
// watchHtmlTemplate reload the HTML template and re-convert all markup
// files.
//
func (w *watcher) watchHtmlTemplate() {
	var (
		logp = "watchHtmlTemplate"

		ns  memfs.NodeState
		err error
	)

	for ns = range w.watchTemplate.C {
		if ns.State == memfs.FileStateDeleted {
			log.Printf("%s: HTML template file %q has been deleted\n",
				logp, ns.Node.SysPath)
			err = w.htmlg.htmlTemplateUseInternal()
		} else {
			fmt.Printf("%s: recompiling HTML template %q ...\n", logp,
				ns.Node.SysPath)
			err = w.htmlg.htmlTemplateReload()
		}
		if err != nil {
			log.Printf("%s: %s", logp, err)
			continue
		}

		fmt.Printf("%s: regenerate all markup files ...\n", logp)
		w.htmlg.convertFileMarkups(w.fileMarkups, true)
	}
}
