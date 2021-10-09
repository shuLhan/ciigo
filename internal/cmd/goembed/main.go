// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generate_main.go

package main

import (
	"log"

	"git.sr.ht/~shulhan/ciigo"
)

func main() {
	opts := ciigo.EmbedOptions{
		ConvertOptions: ciigo.ConvertOptions{
			Root:         "_example",
			HtmlTemplate: "_example/html.tmpl",
		},
		PackageName: "main",
		VarName:     "ciigoFS",
		GoFileName:  "cmd/ciigo-example/static.go",
	}
	err := ciigo.GoEmbed(&opts)
	if err != nil {
		log.Fatal(err)
	}
}
