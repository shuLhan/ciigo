// SPDX-FileCopyrightText: 2021 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import "github.com/shuLhan/share/lib/memfs"

//
// EmbedOptions define the options for calling GoEmbed function.
//
type EmbedOptions struct {
	ConvertOptions
	memfs.EmbedOptions
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
