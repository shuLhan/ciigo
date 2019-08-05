// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generate_main.go
// +build ignore

package main

import (
	"github.com/shuLhan/ciigo"
)

func main() {
	ciigo.Generate("./content", "cmd/ciigo-example/static.go")
}
