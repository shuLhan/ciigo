// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//
// ciigo is a CLI to convert, embed, and/or serve a directory that contains
// markup files, as HTML files.
//
// Usage
//
// The following section describe how to use ciigo CLI.
//
//	ciigo [-template <file>] [-exclude <regex>] convert <dir>
//
// Scan the "dir" recursively to find markup files and convert them into HTML
// files.
// The template "file" is optional, default to embedded HTML template.
//
//	ciigo [-template <file>] [-exclude <regex>] [-out <file>] embed <dir>
//
// Convert all the markup files inside directory "dir" recursively and then
// embed them into ".go" source file.
// The output file is optional, default to "ciigo_static.go" in current
// directory.
//
//	ciigo [-template <file>] [-exclude <regex>] [-address <ip:port>] serve <dir>
//
// Serve all files inside directory "dir" using HTTP server, watch changes on
// markup files and convert them to HTML files.
// If the address is not set, its default to ":8080".
//
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"git.sr.ht/~shulhan/ciigo"
	"github.com/shuLhan/share/lib/debug"
	"github.com/shuLhan/share/lib/memfs"
)

func main() {
	flag.Usage = usage
	isHelp := flag.Bool("help", false, "print help")

	htmlTemplate := flag.String("template", "", "path to HTML template")
	outputFile := flag.String("out", "ciigo_static.go",
		"path to output of .go embed file")
	address := flag.String("address", ":8080",
		"the binding address for HTTP server")
	exclude := flag.String("exclude", "",
		"a regex to exclude certain paths from being scanned during covert, embeded, watch, or serve")

	flag.Parse()

	if *isHelp {
		usage()
		os.Exit(0)
	}

	command := flag.Arg(0)
	if len(command) == 0 {
		usage()
		os.Exit(1)
	}

	dir := flag.Arg(1)
	if len(dir) == 0 {
		dir = ciigo.DefaultRoot
	}

	var err error
	command = strings.ToLower(command)

	switch command {
	case "convert":
		opts := ciigo.ConvertOptions{
			Root:         dir,
			HtmlTemplate: *htmlTemplate,
			Exclude:      *exclude,
		}
		err = ciigo.Convert(&opts)

	case "embed":
		genOpts := ciigo.EmbedOptions{
			ConvertOptions: ciigo.ConvertOptions{
				Root:         dir,
				HtmlTemplate: *htmlTemplate,
				Exclude:      *exclude,
			},
			EmbedOptions: memfs.EmbedOptions{
				GoFileName: *outputFile,
			},
		}
		err = ciigo.GoEmbed(&genOpts)

	case "serve":
		debug.Value = 1
		opts := ciigo.ServeOptions{
			ConvertOptions: ciigo.ConvertOptions{
				Root:         dir,
				HtmlTemplate: *htmlTemplate,
				Exclude:      *exclude,
			},
			Address: *address,
		}
		err = ciigo.Serve(&opts)

	default:
		usage()
		os.Exit(1)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func usage() {
	fmt.Println(`
=  ciigo

A CLI to convert, embed, and/or serve a directory that contains asciidoc markup
files, as HTML files.

==  Usage

ciigo [-template <file>] [-exclude <regex>] convert <dir>

	Scan the "dir" recursively to find markup files.
	and convert them into HTML files.
	The template "file" is optional, default to embedded HTML template.

ciigo [-template <file>] [-exclude <regex>] [-out <file>] embed <dir>

	Convert all markup files inside directory "dir" recursively and then
	embed them into ".go" source file.
	The output file is optional, default to "ciigo_static.go" in current
	directory.

ciigo [-template <file>]  [-exclude <regex>] [-address <ip:port>] serve <dir>

	Serve all files inside directory "dir" using HTTP server, watch
	changes on markup files and convert them to HTML files automatically.
	If the address is not set, its default to ":8080".`)
}
