// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//
// Program ciigo-example provide an example on how to build a binary that
// include the static, generated .go file.
//
package main

import (
	"github.com/shuLhan/ciigo"
)

func main() {
	ciigo.Serve("_example", ":8080", "_example/html.tmpl")
}
