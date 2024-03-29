// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

module git.sr.ht/~shulhan/ciigo

go 1.21

require (
	git.sr.ht/~shulhan/asciidoctor-go v0.5.2-0.20240305110034-dc67158aeeb6
	git.sr.ht/~shulhan/pakakeh.go v0.53.2-0.20240315075343-713d51e4792f
	github.com/yuin/goldmark v1.7.0
	github.com/yuin/goldmark-meta v1.1.0
)

require (
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

//replace git.sr.ht/~shulhan/asciidoctor-go => ../asciidoctor-go

//replace git.sr.ht/~shulhan/pakakeh.go => ../pakakeh.go
