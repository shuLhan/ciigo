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

// watcher watch for changes on all markup files and convert them
// automatically to HTML.
type watcher struct {
	changes       *clise.Clise
	watchDir      *memfs.DirWatcher
	watchTemplate *memfs.Watcher
	converter     *Converter

	// fileMarkups contains all markup files found inside "dir".
	// Its used to convert all markup files when the template file
	// changes.
	fileMarkups map[string]*FileMarkup

	dir string
}

// newWatcher create a watcher that monitor every files changes in directory
// "dir" for new, modified, and deleted markup files and HTML template file.
//
// The watcher depends on Converter to convert the markup to HTML using
// the HTML template in Converter.
//
//	watcher
//	|
//	+-- watchFileMarkup --> UPDATE --> Converter.convertFileMarkups()
//	|
//	+-- watchHTMLTemplate +--> DELETE --> Converter.htmlTemplateUseInternal()
//	                      |
//	                      +--> UPDATE --> Converter.convertFileMarkups()
func newWatcher(converter *Converter, convertOpts *ConvertOptions) (w *watcher, err error) {
	var (
		logp = `newWatcher`
	)

	w = &watcher{
		dir:       convertOpts.Root,
		converter: converter,
		changes:   clise.New(1),
	}
	w.watchDir = &memfs.DirWatcher{
		Options: memfs.Options{
			Root: convertOpts.Root,
			Includes: []string{
				`.*\.adoc$`,
				`.*\.md$`,
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
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	return w, nil
}

// start watching for changes.
func (w *watcher) start() (err error) {
	err = w.watchDir.Start()
	if err != nil {
		return fmt.Errorf(`start: %w`, err)
	}

	go w.watchFileMarkup()

	if len(w.converter.htmlTemplate) > 0 {
		w.watchTemplate, err = memfs.NewWatcher(w.converter.htmlTemplate, 0)
		if err != nil {
			return fmt.Errorf(`start: %w`, err)
		}
		go w.watchHTMLTemplate()
	}
	return nil
}

func (w *watcher) stop() {
	w.watchDir.Stop()
	if w.watchTemplate != nil {
		w.watchTemplate.Stop()
	}
}

// watchFileMarkup watch the markup files inside the "content" directory,
// and re-generate them into HTML file when changed.
func (w *watcher) watchFileMarkup() {
	var (
		logp = `watchFileMarkup`

		ns      memfs.NodeState
		fmarkup *FileMarkup
		ext     string
		err     error
		ok      bool
	)

	for ns = range w.watchDir.C {
		ext = strings.ToLower(filepath.Ext(ns.Node.SysPath))
		if !isExtensionMarkup(ext) {
			continue
		}

		switch ns.State {
		case memfs.FileStateDeleted:
			log.Printf(`%s: %q deleted`, logp, ns.Node.SysPath)
			fmarkup, ok = w.fileMarkups[ns.Node.SysPath]
			if ok {
				delete(w.fileMarkups, ns.Node.SysPath)
				w.changes.Push(fmarkup)
			}
			continue

		case memfs.FileStateCreated:
			log.Printf(`%s: %s created`, logp, ns.Node.SysPath)
			fmarkup, err = NewFileMarkup(ns.Node.SysPath, nil)
			if err != nil {
				log.Printf("%s: %s\n", logp, err)
				continue
			}

			w.fileMarkups[ns.Node.SysPath] = fmarkup

		case memfs.FileStateUpdateMode:
			log.Printf(`%s: %s mode updated`, logp, ns.Node.SysPath)
			continue

		case memfs.FileStateUpdateContent:
			log.Printf(`%s: %s content updated`, logp, ns.Node.SysPath)
			fmarkup = w.fileMarkups[ns.Node.SysPath]
			if fmarkup == nil {
				log.Printf("%s: %s not found\n", logp, ns.Node.SysPath)

				fmarkup, err = NewFileMarkup(ns.Node.SysPath, nil)
				if err != nil {
					log.Printf("%s: %s\n", logp, err)
					continue
				}

				w.fileMarkups[ns.Node.SysPath] = fmarkup
			}
		}

		err = w.converter.ToHTMLFile(fmarkup)
		if err != nil {
			log.Printf(`%s: %s`, logp, err)
		}

		w.changes.Push(fmarkup)
	}
}

// watchHTMLTemplate reload the HTML template and re-convert all markup
// files.
func (w *watcher) watchHTMLTemplate() {
	var (
		logp = `watchHTMLTemplate`

		ns  memfs.NodeState
		err error
	)

	for ns = range w.watchTemplate.C {
		if ns.State == memfs.FileStateDeleted {
			log.Printf("%s: HTML template file %q has been deleted\n",
				logp, ns.Node.SysPath)
			err = w.converter.htmlTemplateUseInternal()
		} else {
			log.Printf(`%s: recompiling HTML template %q ...`, logp, ns.Node.SysPath)
			err = w.converter.SetHTMLTemplateFile(w.converter.htmlTemplate)
		}
		if err != nil {
			log.Printf(`%s: %s`, logp, err)
			continue
		}

		log.Printf(`%s: regenerate all markup files ...`, logp)
		w.converter.convertFileMarkups(w.fileMarkups, true)
	}
}

// waitChanges wait for changes on file markup and return it.
func (w *watcher) waitChanges() (fmarkup *FileMarkup) {
	var ok bool

	for {
		fmarkup, ok = w.changes.Pop().(*FileMarkup)
		if ok {
			break
		}
	}
	return fmarkup
}
