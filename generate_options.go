// Copyright 2021, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

import "github.com/shuLhan/share/lib/memfs"

//
// GenerateOptions define the options for calling Generate function.
//
type GenerateOptions struct {
	ConvertOptions

	// GenPackageName the name of package in Go generated source code.
	// Default to memfs.DefaultGenPackageName if its empty.
	GenPackageName string

	// GenVarName the name of variable where all files in Root will be
	// stored.
	// Default to memfs.DefaultGenVarName if its empty.
	GenVarName string

	// GenGoFileName the file name of Go source code will be written.
	// Default to memfs.DefaultGenGoFileName if its empty.
	GenGoFileName string
}

func (opts *GenerateOptions) init() (err error) {
	err = opts.ConvertOptions.init()
	if err != nil {
		return err
	}
	if len(opts.GenPackageName) == 0 {
		opts.GenPackageName = memfs.DefaultGenPackageName
	}
	if len(opts.GenVarName) == 0 {
		opts.GenVarName = memfs.DefaultGenVarName
	}
	if len(opts.GenGoFileName) == 0 {
		opts.GenGoFileName = memfs.DefaultGenGoFileName
	}
	return nil
}
