// SPDX-FileCopyrightText: 2021 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"git.sr.ht/~shulhan/pakakeh.go/lib/clise"
	"git.sr.ht/~shulhan/pakakeh.go/lib/watchfs/v2"
)

// watcher watch for changes on all markup files and convert them
// automatically to HTML.
type watcher struct {
	changes       *clise.Clise
	watchDir      *watchfs.DirWatcher
	watchTemplate *watchfs.FileWatcher
	converter     *Converter

	// fileMarkups contains all markup files found inside the
	// [options.Root] directory recursively.
	// Its used to convert all markup files when the template file
	// changes.
	fileMarkups map[string]*FileMarkup

	opts watchfs.DirWatcherOptions
}

// newWatcher create a watcher that monitor every files changes in directory
// [options.Root] for new, modified, and deleted markup files and HTML
// template file.
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
func newWatcher(
	converter *Converter, convertOpts ConvertOptions,
) (w *watcher, err error) {
	w = &watcher{
		converter: converter,
		changes:   clise.New(1),
	}
	w.opts = watchfs.DirWatcherOptions{
		FileWatcherOptions: watchfs.FileWatcherOptions{
			File:     filepath.Join(convertOpts.Root, `.ciigo_rescan`),
			Interval: time.Second,
		},
		Root: convertOpts.Root,
		Includes: []string{
			`.*\.(adoc|md)$`,
		},
		Excludes: []string{
			`^\..*`,
			`node_modules/.*`,
			`vendor/.*`,
		},
	}

	w.opts.Excludes = append(w.opts.Excludes, convertOpts.Exclude...)

	w.watchDir, err = watchfs.WatchDir(w.opts)
	if err != nil {
		return nil, err
	}

	w.scanFileMarkup()

	return w, nil
}

// getFileMarkupByHTML get the file markup based on the HTML file name.
func (w *watcher) getFileMarkupByHTML(fileHTML string) (
	fmarkup *FileMarkup, isNew bool,
) {
	// Use file extension to handle insensitive cases of '.html' suffix.
	var ext = filepath.Ext(fileHTML)
	if strings.ToLower(ext) != `.html` {
		return nil, false
	}

	var (
		pathMarkup string
		ok         bool
	)
	pathMarkup, ok = strings.CutSuffix(fileHTML, ext)
	if !ok {
		return nil, false
	}
	pathMarkup = filepath.Join(w.opts.Root, pathMarkup)

	var pathMarkupAdoc = pathMarkup + `.adoc`
	fmarkup = w.fileMarkups[pathMarkupAdoc]
	if fmarkup != nil {
		return fmarkup, false
	}

	var pathMarkupMd = pathMarkup + `.md`
	fmarkup = w.fileMarkups[pathMarkupMd]
	if fmarkup != nil {
		return fmarkup, false
	}

	// Directly check on the file system.

	var fi os.FileInfo
	var err error
	fi, err = os.Stat(pathMarkupAdoc)
	if err == nil {
		fmarkup, err = NewFileMarkup(pathMarkupAdoc, fi)
		if err != nil {
			return nil, false
		}
		w.fileMarkups[pathMarkupAdoc] = fmarkup
		return fmarkup, true
	}

	fi, err = os.Stat(pathMarkupMd)
	if err == nil {
		fmarkup, err = NewFileMarkup(pathMarkupMd, fi)
		if err != nil {
			return nil, false
		}
		w.fileMarkups[pathMarkupMd] = fmarkup
		return fmarkup, true
	}
	return nil, false
}

func (w *watcher) scanFileMarkup() {
	w.fileMarkups = make(map[string]*FileMarkup)
	var files = w.watchDir.Files()
	for path, fi := range files {
		fmarkup, err := NewFileMarkup(path, fi)
		if err != nil {
			continue
		}
		w.fileMarkups[path] = fmarkup
	}
}

// start watching for changes.
func (w *watcher) start() (err error) {
	go w.watchFileMarkup()

	if len(w.converter.htmlTemplate) > 0 {
		var opts = watchfs.FileWatcherOptions{
			File:     w.converter.htmlTemplate,
			Interval: 5 * time.Second,
		}
		w.watchTemplate = watchfs.WatchFile(opts)
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

// watchFileMarkup watch the file ".ciigo_rescan" inside the "content"
// directory and reconvert all the markup into HTML files when its changes.
func (w *watcher) watchFileMarkup() {
	var (
		logp = `watchFileMarkup`

		listfi  []os.FileInfo
		fmarkup *FileMarkup
		err     error
		ok      bool
	)

	for listfi = range w.watchDir.C {
		if len(listfi) == 0 {
			continue
		}

		for _, fi := range listfi {
			var name = fi.Name()

			if fi.Size() == watchfs.FileFlagDeleted {
				log.Printf(`%s: %q deleted`, logp, name)
				fmarkup, ok = w.fileMarkups[name]
				if ok {
					delete(w.fileMarkups, name)
					w.changes.Push(fmarkup)
				}
				continue
			}

			fmarkup = w.fileMarkups[name]
			if fmarkup == nil {
				log.Printf(`%s: %s created`, logp, name)
				fmarkup, err = NewFileMarkup(name, nil)
				if err != nil {
					log.Printf(`%s: %s`, logp, err)
					continue
				}

				w.fileMarkups[name] = fmarkup
			}

			err = w.converter.ToHTMLFile(fmarkup)
			if err != nil {
				log.Printf(`%s: %s`, logp, err)
			}

			log.Printf(`%s: %q converted`, logp, fmarkup.path)
			w.changes.Push(fmarkup)
		}
	}
}

// watchHTMLTemplate reload the HTML template and re-convert all markup
// files.
func (w *watcher) watchHTMLTemplate() {
	var (
		logp = `watchHTMLTemplate`

		err error
	)

	for fi := range w.watchTemplate.C {
		if fi == nil {
			log.Printf(`%s: HTML template file has been deleted`, logp)
			err = w.converter.htmlTemplateUseInternal()
			continue
		}

		log.Printf(`%s: recompiling HTML template %q ...`, logp, fi.Name())
		err = w.converter.SetHTMLTemplateFile(w.converter.htmlTemplate)
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
