// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"github.com/shuLhan/ciigo"
)

const (
	dirRoot = "./content"
)

func main() {
	ciigo.Generate(dirRoot, "cmd/ciigo/static.go")
}
