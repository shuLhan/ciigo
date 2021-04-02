// Copyright 2021, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

import "github.com/shuLhan/share/lib/memfs"

const (
	defAddress = ":8080"
)

//
// ServeOptions contains the options to use on Serve function.
//
type ServeOptions struct {
	ConvertOptions

	// Mfs contains pointer to variable generated from Generate.
	// This option is used to use embedded files for serving on HTTP.
	Mfs *memfs.MemFS

	// Address to listen and serve for HTTP request.
	Address string
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
