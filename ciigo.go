// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//
// Package ciigo is a program to write static web server with embedded files
// using asciidoc and markdown markup languages.
//
// For more information see the README file at the page repository
// https://github.com/shuLhan/ciigo.
//
package ciigo

const (
	dirAssets   = "assets"
	dirRoot     = "./content"
	extAsciidoc = ".adoc"
	extMarkdown = ".md"
)

const (
	markupKindUnknown byte = iota
	markupKindAsciidoc
	markupKindMarkdown
)

var (
	defExcludes = []string{
		`.*\.adoc$`,
		`.*\.md$`,
		`^\..*`,
	}
)

func isExtensionMarkup(ext string) bool {
	return ext == extAsciidoc || ext == extMarkdown
}

func markupKind(ext string) byte {
	switch ext {
	case extAsciidoc:
		return markupKindAsciidoc
	case extMarkdown:
		return markupKindMarkdown
	}
	return markupKindUnknown
}
