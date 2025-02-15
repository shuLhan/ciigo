// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

// ciigo is a CLI to convert, embed, and/or serve a directory that contains
// markup files, as HTML files.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"git.sr.ht/~shulhan/ciigo"
	"git.sr.ht/~shulhan/pakakeh.go/lib/memfs"
)

const (
	cmdConvert = "convert"
	cmdEmbed   = "embed"
	cmdHelp    = "help"
	cmdServe   = "serve"
	cmdVersion = "version"
)

func main() {
	var (
		htmlTemplate *string
		outputFile   *string
		address      *string
		exclude      *string
	)

	flag.Usage = usage

	htmlTemplate = flag.String("template", "", "path to HTML template")
	outputFile = flag.String("out", "ciigo_static.go",
		"path to output of .go embed file")
	address = flag.String(`address`, `127.0.0.1:6320`,
		"the binding address for HTTP server")
	exclude = flag.String("exclude", "",
		"a regex to exclude certain paths from being scanned during covert, embeded, watch, or serve")

	var flagPackageName = flag.String(`package-name`, `main`, `package name for embed`)
	var flagVarName = flag.String(`var-name`, `memFS`, `variable name for embed`)

	flag.Parse()

	var (
		command     = strings.ToLower(flag.Arg(0))
		convertOpts = ciigo.ConvertOptions{
			Root:         flag.Arg(1),
			HTMLTemplate: *htmlTemplate,
			Exclude:      []string{*exclude},
		}
		err error
	)

	if len(command) == 0 {
		usage()
		os.Exit(1)
	}
	if len(convertOpts.Root) == 0 {
		convertOpts.Root = ciigo.DefaultRoot
	}

	switch command {
	case cmdConvert:
		err = ciigo.Convert(convertOpts)

	case cmdEmbed:
		var embedOpts = ciigo.EmbedOptions{
			ConvertOptions: convertOpts,
			EmbedOptions: memfs.EmbedOptions{
				GoFileName:  *outputFile,
				PackageName: *flagPackageName,
				VarName:     *flagVarName,
			},
		}

		err = ciigo.GoEmbed(embedOpts)

	case cmdHelp:
		usage()

	case cmdServe:
		var serveOpts = ciigo.ServeOptions{
			ConvertOptions:  convertOpts,
			Address:         *address,
			EnableIndexHTML: true,
			IsDevelopment:   true,
		}
		err = ciigo.Serve(serveOpts)

	case cmdVersion:
		fmt.Println(ciigo.Version)

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

ciigo help

	Print the usage (this output).

ciigo [-template <file>] [-exclude <regex>] [-out <file>]
	[-package-name <string>] [-var-name <string>]
	embed <dir>

	Convert all markup files inside directory "dir" recursively and then
	embed them into ".go" source file.
	The output file is optional, default to "ciigo_static.go" in current
	directory.
	The package name default to main.
	The variable name default to memFS.

ciigo [-template <file>]  [-exclude <regex>] [-address <ip:port>] serve <dir>

	Serve all files inside directory "dir" using HTTP server, watch
	changes on markup files and convert them to HTML files automatically.
	If the address is not set, its default to ":8080".

ciigo version

	Print the current ciigo version.`)
}
