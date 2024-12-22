// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

module git.sr.ht/~shulhan/ciigo

go 1.23.4

require (
	git.sr.ht/~shulhan/asciidoctor-go v0.6.2-0.20250106025231-ed20fc1c23e8
	git.sr.ht/~shulhan/pakakeh.go v0.59.0
	github.com/yuin/goldmark v1.7.8
	github.com/yuin/goldmark-meta v1.1.0
)

require (
	golang.org/x/exp v0.0.0-20250103183323-7d7fa50e5329 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

//replace git.sr.ht/~shulhan/asciidoctor-go => ../asciidoctor-go

//replace git.sr.ht/~shulhan/pakakeh.go => ../pakakeh.go
