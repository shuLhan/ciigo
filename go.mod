// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

module git.sr.ht/~shulhan/ciigo

go 1.20

require (
	git.sr.ht/~shulhan/asciidoctor-go v0.5.1
	git.sr.ht/~shulhan/pakakeh.go v0.53.2-0.20240305092154-76510776395b
	github.com/yuin/goldmark v1.6.0
	github.com/yuin/goldmark-meta v1.1.0
)

require (
	github.com/shuLhan/share v0.51.0 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

//replace git.sr.ht/~shulhan/asciidoctor-go => ../asciidoctor-go

//replace git.sr.ht/~shulhan/pakakeh.go => ../pakakeh.go
