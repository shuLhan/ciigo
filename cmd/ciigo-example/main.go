// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//
// Program ciigo-example provide an example on how to build a binary that
// include the static, generated .go file.
//
package main

import (
	"log"

	"github.com/shuLhan/share/lib/memfs"

	"git.sr.ht/~shulhan/ciigo"
)

var ciigoFS *memfs.MemFS

func main() {
	opts := &ciigo.ServeOptions{
		ConvertOptions: ciigo.ConvertOptions{
			Root:         "_example",
			HtmlTemplate: "_example/html.tmpl",
		},
		Mfs:     ciigoFS,
		Address: ":8080",
	}

	err := ciigo.Serve(opts)
	if err != nil {
		log.Fatal(err)
	}
}
