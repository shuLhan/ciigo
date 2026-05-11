// SPDX-License-Identifier: GPL-3.0-only
// SPDX-FileCopyrightText: 2021 Shulhan <ms@kilabit.info>

package ciigo

import "git.sr.ht/~shulhan/pakakeh.go/lib/http"

const (
	defAddress = `:8080`
)

// ServeOptions define the options to use on Serve function.
// See [git.sr.ht/~shulhan/pakakeh.go/lib/http.ServerOptions] for more information.
type ServeOptions struct {
	http.ServerOptions

	// IsDevelopment if set to true, it will serve the ConvertOptions.Root
	// directory directly and watch all asciidoc files for changes and
	// convert it.
	// This is like running Watch, Convert and Serve at the same time.
	IsDevelopment bool
}

func (opts *ServeOptions) init() {
	if len(opts.Address) == 0 {
		opts.Address = defAddress
	}
}
