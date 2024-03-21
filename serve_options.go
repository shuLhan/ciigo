// SPDX-FileCopyrightText: 2021 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import "git.sr.ht/~shulhan/pakakeh.go/lib/memfs"

const (
	defAddress = `:8080`
)

// ServeOptions contains the options to use on Serve function.
type ServeOptions struct {
	// Mfs contains pointer to variable generated from Generate.
	// This option is used to use embedded files for serving on HTTP.
	Mfs *memfs.MemFS

	// Address to listen and serve for HTTP request.
	Address string

	ConvertOptions

	// If true, the serve command generate index.html automatically if its
	// not exist in the directory.
	// The index.html contains the list of files inside the requested
	// path.
	EnableIndexHTML bool

	// IsDevelopment if set to true, it will serve the ConvertOptions.Root
	// directory directly and watch all asciidoc files for changes and
	// convert it.
	// This is like running Watch, Convert and Serve at the same time.
	IsDevelopment bool
}

func (opts *ServeOptions) init() (err error) {
	err = opts.ConvertOptions.init()
	if err != nil {
		return err
	}
	if len(opts.Address) == 0 {
		opts.Address = defAddress
	}
	return nil
}
