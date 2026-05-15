// SPDX-License-Identifier: GPL-3.0-only
// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>

// ciigo is a CLI to convert, embed, and/or serve a directory that contains
// markup files, as HTML files.
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"kilabit.info/ciigo"
	libhttp "kilabit.info/pakakeh.go/lib/http"
	"kilabit.info/pakakeh.go/lib/memfs"
	"kilabit.info/pakakeh.go/lib/systemd"
)

const (
	cmdConvert = "convert"
	cmdEmbed   = "embed"
	cmdHelp    = "help"
	cmdServe   = "serve"
	cmdVersion = "version"
)

func main() {
	var serveOpts = ciigo.ServeOptions{
		ServerOptions: libhttp.ServerOptions{
			EnableIndexHTML: true,
		},
		IsDevelopment: true,
	}
	var (
		htmlTemplate *string
		outputFile   *string
		exclude      *string
	)
	var shutdownIdleDuration string

	flag.Usage = usage

	htmlTemplate = flag.String("template", "", "path to HTML template")
	outputFile = flag.String("out", "ciigo_static.go",
		"path to output of .go embed file")

	flag.StringVar(&serveOpts.Address, `address`, `127.0.0.1:6320`,
		`Set the listening address for HTTP server.`)
	flag.StringVar(&serveOpts.BasePath, `base-path`, ``,
		`Set the base path (or prefix) for handling HTTP request.`)
	flag.StringVar(&shutdownIdleDuration, `shutdown-idle`, ``,
		`Set the duration when server will shutting down after idle.`)

	exclude = flag.String("exclude", "",
		"a regex to exclude certain paths from being scanned during covert, embeded, watch, or serve")

	var flagPackageName = flag.String(`package-name`, `main`, `package name for embed`)
	var flagVarName = flag.String(`var-name`, `memFS`, `variable name for embed`)

	flag.Parse()

	var err error
	if shutdownIdleDuration != `` {
		serveOpts.ShutdownIdleDuration, err = time.ParseDuration(shutdownIdleDuration)
		if err != nil {
			log.Fatalf(`invalid shutdown-idle %s: %s`, shutdownIdleDuration, err)
		}
	}

	var (
		command     = strings.ToLower(flag.Arg(0))
		convertOpts = ciigo.ConvertOptions{
			Root:         flag.Arg(1),
			HTMLTemplate: *htmlTemplate,
			Exclude:      []string{*exclude},
		}
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
		var listeners []net.Listener
		listeners, err = systemd.Listeners(true)
		if err != nil {
			log.Fatal(err)
		}
		if len(listeners) > 1 {
			log.Fatal(`too many listeners received for activation`)
		}
		if len(listeners) == 1 {
			serveOpts.Listener = listeners[0]
			gotAddr := serveOpts.Listener.Addr().String()
			if gotAddr != serveOpts.Address {
				log.Fatalf(`invalid Listener address, got %s, want %s`,
					gotAddr, serveOpts.Address)
			}
		}
		err = ciigo.Serve(serveOpts, convertOpts)

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

ciigo [-template <file>]  [-exclude <regex>] [-address <ip:port>]
	[-base-path <string>]
	[-shutdown-idle <duration>]
	serve <dir>

	Serve all files inside directory "dir" using HTTP server, watch
	changes on markup files and convert them to HTML files automatically.
	If the address is not set, its default to ":8080".

	The -base-path set the base path or prefix for serving HTTP request.

	The -shutdown-idle set the duration when the serve will automatically
	shutting down when idle.

ciigo version

	Print the current ciigo version.`)
}
