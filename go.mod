// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

module git.sr.ht/~shulhan/ciigo

go 1.19

require (
	git.sr.ht/~shulhan/asciidoctor-go v0.5.0
	github.com/shuLhan/share v0.50.0
	github.com/yuin/goldmark v1.5.6
	github.com/yuin/goldmark-meta v1.1.0
)

require (
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

//replace git.sr.ht/~shulhan/asciidoctor-go => ../asciidoctor-go

//replace github.com/shuLhan/share => ../share
