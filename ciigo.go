// SPDX-FileCopyrightText: 2019 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

// Package ciigo is a program to write static web server with embedded files
// using the asciidoc markup languages.
//
// For more information see the README file at the page repository
// https://sr.ht/~shulhan/ciigo.
package ciigo

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/shuLhan/share/lib/memfs"
)

const (
	extAsciidoc          = `.adoc`
	extMarkdown          = `.md`
	internalTemplatePath = `_internal/.template`
)

// Version define the latest tagged release of this module.
var Version = `0.10.1`

// defExcludes define default files to be excludes on GoEmbed.
var defExcludes = []string{
	`.*\.adoc$`,
	`.*\.md$`,
	`^\..*`,
}

// Convert all markup files inside directory "dir" recursively into HTML
// files using ConvertOptions HtmlTemplate file as base template.
// If HtmlTemplate is empty it will default to use embedded HTML template.
// See template_index_html.go for template format.
func Convert(opts *ConvertOptions) (err error) {
	var (
		logp = `Convert`

		converter   *Converter
		fileMarkups map[string]*FileMarkup
	)

	if opts == nil {
		opts = &ConvertOptions{}
	}
	err = opts.init()
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	converter, err = NewConverter(opts.HtmlTemplate)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	fileMarkups, err = listFileMarkups(opts.Root, opts.excRE)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	converter.convertFileMarkups(fileMarkups, false)

	return nil
}

// GoEmbed generate a static Go file that embed all files inside Root except
// the one that being excluded explicitly by ConvertOptions Exclude.
//
// It convert all markup files inside directory "dir" into HTML files,
// recursively, and then embed them into Go file defined by
// EmbedOptions.GoFileName.
//
// If HtmlTemplate option is empty it default to use embedded HTML
// template.
// See template_index_html.go for template format.
func GoEmbed(opts *EmbedOptions) (err error) {
	var (
		logp = `GoEmbed`

		converter    *Converter
		fileMarkups  map[string]*FileMarkup
		mfs          *memfs.MemFS
		mfsOpts      *memfs.Options
		convertForce bool
	)

	if opts == nil {
		opts = &EmbedOptions{}
	}
	err = opts.init()
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	converter, err = NewConverter(opts.HtmlTemplate)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	fileMarkups, err = listFileMarkups(opts.Root, opts.excRE)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	if isHtmlTemplateNewer(opts) {
		convertForce = true
	}

	converter.convertFileMarkups(fileMarkups, convertForce)

	mfsOpts = &memfs.Options{
		Root:     opts.Root,
		Excludes: defExcludes,
		Embed:    opts.EmbedOptions,
	}

	mfs, err = memfs.New(mfsOpts)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	if len(opts.HtmlTemplate) > 0 {
		_, err = mfs.AddFile(internalTemplatePath, opts.HtmlTemplate)
		if err != nil {
			return fmt.Errorf(`%s: %w`, logp, err)
		}
	}

	err = mfs.GoEmbed()
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	return nil
}

// Serve the content at directory "dir" using HTTP server at specific
// "address".
func Serve(opts *ServeOptions) (err error) {
	var (
		logp = `Serve`
		srv  *server
	)

	if opts == nil {
		opts = &ServeOptions{}
	}
	err = opts.init()
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	srv, err = newServer(opts)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}
	err = srv.start()
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}
	return nil
}

// Watch any changes on markup files on directory Root recursively and
// changes on the HTML template file.
// If there is new or modified markup files it will convert them into HTML
// files using HTML template automatically.
//
// If the HTML template file modified, it will re-convert all markup files.
// If the HTML template file deleted, it will replace them with internal,
// default HTML template.
func Watch(opts *ConvertOptions) (err error) {
	var (
		logp = `Watch`

		converter *Converter
		w         *watcher
	)

	if opts == nil {
		opts = &ConvertOptions{}
	}
	err = opts.init()
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	converter, err = NewConverter(opts.HtmlTemplate)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	w, err = newWatcher(converter, opts)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	err = w.start()
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	return nil
}

// isHtmlTemplateNewer will return true if HtmlTemplate is not defined or
// newer than embedded GoFileName.
func isHtmlTemplateNewer(opts *EmbedOptions) bool {
	var (
		logp = `isHtmlTemplateNewer`

		fiHtmlTmpl fs.FileInfo
		fiGoEmbed  fs.FileInfo
		err        error
	)

	if len(opts.HtmlTemplate) == 0 {
		return true
	}

	fiHtmlTmpl, err = os.Stat(opts.HtmlTemplate)
	if err != nil {
		log.Fatalf(`%s: %s`, logp, err)
	}

	if len(opts.EmbedOptions.GoFileName) == 0 {
		// No output file for GoEmbed.
		return false
	}

	fiGoEmbed, err = os.Stat(opts.EmbedOptions.GoFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		log.Fatalf(`%s: %s`, logp, err)
	}

	return fiHtmlTmpl.ModTime().After(fiGoEmbed.ModTime())
}

// isExtensionMarkup return true if the file extension ext match with one of
// supported markup format.
func isExtensionMarkup(ext string) bool {
	if ext == extAsciidoc {
		return true
	}
	return ext == extMarkdown
}

// listFileMarkups find any markup files inside the content directory,
// recursively.
func listFileMarkups(dir string, excRE []*regexp.Regexp) (
	fileMarkups map[string]*FileMarkup, err error,
) {
	var (
		logp = `listFileMarkups`

		d        *os.File
		fi       os.FileInfo
		fmarkup  *FileMarkup
		name     string
		filePath string
		k        string
		ext      string
		fis      []os.FileInfo
		fmarkups map[string]*FileMarkup
	)

	d, err = os.Open(dir)
	if err != nil {
		if os.IsPermission(err) {
			return nil, nil
		}
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	fis, err = d.Readdir(0)
	if err != nil {
		return nil, fmt.Errorf(`%s: %w`, logp, err)
	}

	fileMarkups = make(map[string]*FileMarkup)

	for _, fi = range fis {
		name = fi.Name()
		filePath = filepath.Join(dir, name)

		if isExcluded(filePath, excRE) {
			continue
		}

		if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
			fi, err = os.Stat(filePath)
			if err != nil {
				return nil, fmt.Errorf(`%s: %w`, logp, err)
			}
		}

		if fi.IsDir() {
			if name[0] == '.' {
				// Skip any directory start with '.'.
				continue
			}
			fmarkups, err = listFileMarkups(filePath, excRE)
			if err != nil {
				return nil, fmt.Errorf(`%s: %s: %w`, logp, filePath, err)
			}
			for k, fmarkup = range fmarkups {
				fileMarkups[k] = fmarkup
			}
			continue
		}

		ext = strings.ToLower(filepath.Ext(name))
		if !isExtensionMarkup(ext) {
			continue
		}
		if fi.Size() == 0 {
			continue
		}
		fmarkup, err = NewFileMarkup(filePath, fi)
		if err != nil {
			return nil, fmt.Errorf(`%s: %s: %w`, logp, filePath, err)
		}
		fileMarkups[filePath] = fmarkup
	}

	return fileMarkups, nil
}

func isExcluded(path string, excs []*regexp.Regexp) bool {
	if len(excs) == 0 {
		return false
	}

	var re *regexp.Regexp
	for _, re = range excs {
		if re.MatchString(path) {
			return true
		}
	}
	return false
}
