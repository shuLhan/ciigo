// Copyright 2021, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

const (
	// DefaultRoot define default Root value for GenerateOptions.
	DefaultRoot = "."
)

//
// ConvertOptions define the options to use on Convert function.
//
type ConvertOptions struct {
	// Root directory where its content will be embedded into Go source
	// code.
	// Default to DefaultRoot if its empty.
	Root string

	// HtmlTemplate the HTML template to be used when converting asciidoc
	// file into HTML.
	// If empty it will default to use embedded HTML template.
	// See template_index_html.go for template format.
	HtmlTemplate string
}

func (opts *ConvertOptions) init() {
	if len(opts.Root) == 0 {
		opts.Root = DefaultRoot
	}
}
