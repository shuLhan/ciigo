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
	opts := ciigo.GenerateOptions{
		Root:           "_example",
		HTMLTemplate:   "_example/html.tmpl",
		GenPackageName: "main",
		GenVarName:     "ciigoFS",
		GenGoFileName:  "cmd/ciigo-example/static.go",
	}
	err := ciigo.Generate(&opts)
	if err != nil {
		log.Fatal(err)
	}
}
