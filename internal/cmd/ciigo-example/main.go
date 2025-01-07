// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

// Program ciigo-example provide an example on how to build a binary that
// include the static, generated .go file.
package main

import (
	"flag"
	"log"
	"strings"

	"git.sr.ht/~shulhan/pakakeh.go/lib/memfs"

	"git.sr.ht/~shulhan/ciigo"
)

const (
	cmdEmbed = `embed`
)

var ciigoFS *memfs.MemFS

func main() {
	var (
		opts = ciigo.ServeOptions{
			ConvertOptions: ciigo.ConvertOptions{
				Root:         `_example`,
				HTMLTemplate: `_example/html.tmpl`,
			},
			Mfs:           ciigoFS,
			Address:       `127.0.0.1:8080`,
			IsDevelopment: true,
		}

		cmd string
		err error
	)

	flag.Parse()

	cmd = strings.ToLower(flag.Arg(0))
	if cmd == cmdEmbed {
		doEmbed()
		return
	}

	err = ciigo.Serve(opts)
	if err != nil {
		log.Fatal(err)
	}
}

func doEmbed() {
	var opts = ciigo.EmbedOptions{
		ConvertOptions: ciigo.ConvertOptions{
			Root:         `_example`,
			HTMLTemplate: `_example/html.tmpl`,
			Exclude: []string{
				`.*/\..*$`,
			},
		},
		EmbedOptions: memfs.EmbedOptions{
			CommentHeader: `// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later
`,
			PackageName:    `main`,
			VarName:        `ciigoFS`,
			GoFileName:     `internal/cmd/ciigo-example/static.go`,
			WithoutModTime: true,
		},
	}

	var err = ciigo.GoEmbed(opts)
	if err != nil {
		log.Fatal(err)
	}
}
