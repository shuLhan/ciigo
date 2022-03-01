// SPDX-FileCopyrightText: 2021 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"fmt"
	"regexp"
)

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

	// Exclude define regular expresion to exclude certain paths from
	// being scanned.
	Exclude string

	// HtmlTemplate the HTML template to be used when converting markup
	// file into HTML.
	// If empty it will default to use embedded HTML template.
	// See template_index_html.go for template format.
	HtmlTemplate string

	excRE []*regexp.Regexp
}

func (opts *ConvertOptions) init() (err error) {
	var (
		logp = "ConvertOptions.init"
	)

	if len(opts.Root) == 0 {
		opts.Root = DefaultRoot
	}
	if len(opts.Exclude) > 0 {
		var re *regexp.Regexp

		re, err = regexp.Compile(opts.Exclude)
		if err != nil {
			return fmt.Errorf("%s: %w", logp, err)
		}

		opts.excRE = append(opts.excRE, re)
		defExcludes = append(defExcludes, opts.Exclude)
	}
	return nil

}
