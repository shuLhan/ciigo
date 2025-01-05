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

	libhttp "git.sr.ht/~shulhan/pakakeh.go/lib/http"
	"git.sr.ht/~shulhan/pakakeh.go/lib/memfs"
)

const (
	extAsciidoc          = `.adoc`
	extMarkdown          = `.md`
	internalTemplatePath = `_internal/.template`
)

// Version define the latest tagged release of this module.
var Version = `0.14.0`

// defExcludes define default files to be excludes on GoEmbed.
var defExcludes = []string{
	`.*\.adoc$`,
	`.*\.md$`,
	`^\..*`,
}

// Ciigo provides customizable and reusable instance of ciigo for embedding,
// converting, and/or serving HTTP server.
// This type is introduced so one can add HTTP handler or endpoint along
// with serving the files.
type Ciigo struct {
	HTTPServer *libhttp.Server
	converter  *Converter
	watcher    *watcher
	serveOpts  ServeOptions
}

// Convert all markup files inside directory "dir" recursively into HTML
// files using ConvertOptions HTMLTemplate file as base template.
// If HTMLTemplate is empty it will default to use embedded HTML template.
// See template_index_html.go for template format.
func Convert(opts ConvertOptions) (err error) {
	var ciigo = &Ciigo{}
	return ciigo.Convert(opts)
}

// GoEmbed generate a static Go file that embed all files inside Root except
// the one that being excluded explicitly by ConvertOptions Exclude.
//
// It convert all markup files inside directory "dir" into HTML files,
// recursively, and then embed them into Go file defined by
// EmbedOptions.GoFileName.
//
// If HTMLTemplate option is empty it default to use embedded HTML
// template.
// See template_index_html.go for template format.
func GoEmbed(opts EmbedOptions) (err error) {
	var ciigo = &Ciigo{}
	return ciigo.GoEmbed(opts)
}

// Serve the content under directory "[ServeOptions].ConvertOptions.Root"
// using HTTP server at specific "[ServeOptions].Address".
func Serve(opts ServeOptions) (err error) {
	var ciigo = &Ciigo{}
	err = ciigo.InitHTTPServer(opts)
	if err != nil {
		return err
	}
	return ciigo.Serve()
}

// Watch any changes on markup files on directory Root recursively and
// changes on the HTML template file.
// If there is new or modified markup files it will convert them into HTML
// files using HTML template automatically.
//
// If the HTML template file modified, it will re-convert all markup files.
// If the HTML template file deleted, it will replace them with internal,
// default HTML template.
func Watch(opts ConvertOptions) (err error) {
	var ciigo = &Ciigo{}
	return ciigo.Watch(opts)
}

// isHTMLTemplateNewer will return true if HTMLTemplate is not defined or
// newer than embedded GoFileName.
func isHTMLTemplateNewer(opts EmbedOptions) bool {
	var (
		logp = `isHTMLTemplateNewer`

		fiHTMLTmpl fs.FileInfo
		fiGoEmbed  fs.FileInfo
		err        error
	)

	if len(opts.HTMLTemplate) == 0 {
		return false
	}

	fiHTMLTmpl, err = os.Stat(opts.HTMLTemplate)
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

	return fiHTMLTmpl.ModTime().After(fiGoEmbed.ModTime())
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
		log.Printf(`%s: %s`, logp, err)
		return nil, nil
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
				// Broken symlink.
				log.Printf(`%s: %s`, logp, err)
				continue
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

// Convert all markup files inside directory [ConvertOptions.Root]
// recursively into HTML files using [ConvertOptions.HTMLTemplate] file as
// base template.
// If HTMLTemplate is empty it use the default embedded HTML template.
// See template_index_html.go for template format.
func (ciigo *Ciigo) Convert(opts ConvertOptions) (err error) {
	var logp = `Convert`

	err = opts.init()
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	ciigo.serveOpts.ConvertOptions = opts

	ciigo.converter, err = NewConverter(opts.HTMLTemplate)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	var fileMarkups map[string]*FileMarkup

	fileMarkups, err = listFileMarkups(opts.Root, opts.excRE)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	ciigo.converter.convertFileMarkups(fileMarkups, false)

	return nil
}

// GoEmbed embed the file system (directories and files) inside the
// [ConvertOptions.Root] into a Go code.
// One can exclude files by writing regular expression in
// [ConvertOptions.Exclude].
func (ciigo *Ciigo) GoEmbed(embedOpts EmbedOptions) (err error) {
	var logp = `GoEmbed`

	err = embedOpts.init()
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	ciigo.converter, err = NewConverter(embedOpts.HTMLTemplate)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	var fileMarkups map[string]*FileMarkup

	fileMarkups, err = listFileMarkups(embedOpts.Root, embedOpts.excRE)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	var convertForce = isHTMLTemplateNewer(embedOpts)

	ciigo.converter.convertFileMarkups(fileMarkups, convertForce)

	var mfsOpts = &memfs.Options{
		Root:     embedOpts.Root,
		Excludes: defExcludes,
		Embed:    embedOpts.EmbedOptions,
	}

	var mfs *memfs.MemFS

	mfs, err = memfs.New(mfsOpts)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	if len(embedOpts.HTMLTemplate) > 0 {
		_, err = mfs.AddFile(internalTemplatePath, embedOpts.HTMLTemplate)
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

// Watch start a watcher on [ConvertOptions.Root] directory that monitor any
// changes to markup files and convert them to HTML files.
func (ciigo *Ciigo) Watch(convertOpts ConvertOptions) (err error) {
	var logp = `Watch`

	err = convertOpts.init()
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	ciigo.converter, err = NewConverter(convertOpts.HTMLTemplate)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	ciigo.watcher, err = newWatcher(ciigo.converter, convertOpts)
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}

	err = ciigo.watcher.start()
	if err != nil {
		return fmt.Errorf(`%s: %w`, logp, err)
	}
	return nil
}
