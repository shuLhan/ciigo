// SPDX-FileCopyrightText: 2021 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import "github.com/shuLhan/share/lib/memfs"

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
