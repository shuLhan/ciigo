// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

module git.sr.ht/~shulhan/ciigo

go 1.18

require (
	git.sr.ht/~shulhan/asciidoctor-go v0.4.1
	github.com/shuLhan/share v0.44.0
	github.com/yuin/goldmark v1.5.4
	github.com/yuin/goldmark-meta v1.1.0
)

require (
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

//replace git.sr.ht/~shulhan/asciidoctor-go => ../asciidoctor-go

//replace github.com/shuLhan/share => ../share
