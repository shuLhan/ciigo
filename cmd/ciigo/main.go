// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//
// ciigo is a CLI to convert, generate, and/or serve a directory that contains
// markup files, as HTML files.
//
// Usage
//
// The following section describe how to use ciigo CLI.
//
//	ciigo [-template <file>] convert <dir>
//
// Scan the "dir" recursively to find markup files (.adoc or .md) and convert
// them into HTML files.
// The template "file" is optional, default to "templates/html.tmpl" in the
// current directory.
//
//	ciigo [-template <file>] [-out <file>] generate <dir>
//
// Convert all the markup files inside directory "dir" recursively and then
// embed them into ".go" source file.
// The output file is optional, default to "ciigo_static.go" in current
// directory.
//
//	ciigo [-template <file>] [-address <ip:port>] serve <dir>
//
// Serve all files inside directory "dir" using HTTP server, watch changes on
// markup files and convert them to HTML files.
// If the address is not set, its default to ":8080".
//
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/shuLhan/ciigo"
	"github.com/shuLhan/share/lib/debug"
)

func main() {
	isHelp := flag.Bool("help", false, "print help")

	htmlTemplate := flag.String("template", "templates/html.tmpl",
		"path to HTML template")
	outputFile := flag.String("out", "ciigo_static.go",
		"path to output of .go generated file")
	address := flag.String("address", ":8080",
		"the binding address for HTTP server")

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
		dir = "."
	}

	command = strings.ToLower(command)
	switch command {
	case "convert":
		ciigo.Convert(dir, *htmlTemplate)
	case "generate":
		ciigo.Generate(dir, *outputFile, *htmlTemplate)
	case "serve":
		debug.Value = 2
		srv := ciigo.NewServer(dir, *address, *htmlTemplate)
		srv.Start()
	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Println(`
=  ciigo

A CLI to convert, generate, and/or serve a directory that contains markup
files, as HTML files.

==  Usage

ciigo [-template <file>] convert <dir>

	Scan the "dir" recursively to find markup files (.adoc or .md)
	and convert them into HTML files.
	The template "file" is optional, default to "templates/html.tmpl" in
	the current directory.

ciigo [-template <file>] [-out <file>] generate <dir>

	Convert all markup files inside directory "dir" recursively and then
	embed them into ".go" source file.
	The output file is optional, default to "ciigo_static.go" in current
	directory.

ciigo [-template <file>] [-address <ip:port>] serve <dir>

	Serve all files inside directory "dir" using HTTP server, watch
	changes on markup files and convert them to HTML files automatically.
	If the address is not set, its default to ":8080".`)
}
