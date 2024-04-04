// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

module git.sr.ht/~shulhan/ciigo

go 1.21

require (
	git.sr.ht/~shulhan/asciidoctor-go v0.5.2
	git.sr.ht/~shulhan/pakakeh.go v0.54.0
	github.com/yuin/goldmark v1.7.1
	github.com/yuin/goldmark-meta v1.1.0
)

require (
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

//replace git.sr.ht/~shulhan/asciidoctor-go => ../asciidoctor-go

//replace git.sr.ht/~shulhan/pakakeh.go => ../pakakeh.go
