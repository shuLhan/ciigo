// SPDX-License-Identifier: GPL-3.0-only
// SPDX-FileCopyrightText: 2026 Shulhan <ms@kilabit.info>

package main

import (
	"log"

	"git.sr.ht/~shulhan/pakakeh.go/lib/memfs"
)

func main() {
	var opts = &memfs.Options{
		Embed: memfs.EmbedOptions{
			PackageName: `ciigo`,
			VarName:     `staticfs`,
			GoFileName:  `staticfs.go`,
			CommentHeader: `// SPDX-License-Identifier: GPL-3.0-only
// SPDX-FileCopyrightText: 2026 Shulhan <ms@kilabit.info>

`,
		},
		Root: `./_static`,
	}

	staticfs, err := memfs.New(opts)
	if err != nil {
		log.Fatal(err)
	}

	err = staticfs.GoEmbed()
	if err != nil {
		log.Fatal(err)
	}
}
