// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generate_main.go

package main

import "git.sr.ht/~shulhan/ciigo"

func main() {
	ciigo.Generate("./_example", "cmd/ciigo-example/static.go",
		"_example/html.tmpl")
}
