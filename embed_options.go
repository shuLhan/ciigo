// Copyright 2021, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

import "github.com/shuLhan/share/lib/memfs"

//
// EmbedOptions define the options for calling GoEmbed function.
//
type EmbedOptions struct {
	ConvertOptions

	// PackageName the name of package in Go generated source code.
	// Default to memfs.DefaultEmbedPackageName if its empty.
	PackageName string

	// GenVarName the name of variable where all files in Root will be
	// stored.
	// Default to memfs.DefaultEmbedVarName if its empty.
	VarName string

	// GenGoFileName the file name of Go source code will be written.
	// Default to memfs.DefaultEmbedGoFileName if its empty.
	GoFileName string
}

func (opts *EmbedOptions) init() (err error) {
	err = opts.ConvertOptions.init()
	if err != nil {
		return err
	}
	if len(opts.PackageName) == 0 {
		opts.PackageName = memfs.DefaultEmbedPackageName
	}
	if len(opts.VarName) == 0 {
		opts.VarName = memfs.DefaultEmbedVarName
	}
	if len(opts.GoFileName) == 0 {
		opts.GoFileName = memfs.DefaultEmbedGoFileName
	}
	return nil
}
