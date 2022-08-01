// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generate_main.go

package main

import (
	"log"

	"github.com/shuLhan/share/lib/memfs"

	"git.sr.ht/~shulhan/ciigo"
)

func main() {
	var (
		opts = ciigo.EmbedOptions{
			ConvertOptions: ciigo.ConvertOptions{
				Root:         "_example",
				HtmlTemplate: "_example/html.tmpl",
			},
			EmbedOptions: memfs.EmbedOptions{
				CommentHeader: `// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later
`,
				PackageName: "main",
				VarName:     "ciigoFS",
				GoFileName:  "cmd/ciigo-example/static.go",
			},
		}

		err error
	)

	err = ciigo.GoEmbed(&opts)
	if err != nil {
		log.Fatal(err)
	}
}
