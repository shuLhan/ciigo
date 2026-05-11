// SPDX-License-Identifier: GPL-3.0-or-later
// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>

package ciigo

import (
	"log"
	"os"
	"regexp"
	"sort"
	"testing"

	"git.sr.ht/~shulhan/pakakeh.go/lib/memfs"
	"git.sr.ht/~shulhan/pakakeh.go/lib/test"
)

func TestMain(m *testing.M) {
	if staticfs != nil {
		css, err := staticfs.Get(`/ciigo.css`)
		if err != nil {
			log.Fatal(err)
		}
		css.Content = []byte(`body{}`)
	}
	os.Exit(m.Run())
}

func TestListFileMarkups(t *testing.T) {
	type testCase struct {
		excRegex string
		exp      []string
	}

	var cases = []testCase{{
		excRegex: `(ex)/.*`,
		exp: []string{
			`testdata/in/clu/de/file.adoc`,
			`testdata/in/clu/de/markdown.md`,
		},
	}, {
		excRegex: `(in|ex)/.*`,
	}, {
		excRegex: `de`,
	}, {
		excRegex: `file$`,
		exp: []string{
			`testdata/ex/clu/de/file.adoc`,
			`testdata/ex/clu/de/markdown.md`,
			`testdata/in/clu/de/file.adoc`,
			`testdata/in/clu/de/markdown.md`,
		},
	}}

	var (
		dir = `testdata`

		c     testCase
		excre *regexp.Regexp
		list  map[string]*FileMarkup
		got   []string
		k     string
		err   error
	)

	for _, c = range cases {
		excre = regexp.MustCompile(c.excRegex)

		list, err = listFileMarkups(dir, []*regexp.Regexp{excre})
		if err != nil {
			t.Fatal(err)
		}

		got = make([]string, 0, len(list))
		for k = range list {
			got = append(got, k)
		}

		sort.Strings(got)

		test.Assert(t, `list`, c.exp, got)
	}
}

func TestGoEmbed(t *testing.T) {
	t.Skip(`TODO: compare the list and content of files, instead of Go code`)
	type testCase struct {
		tag  string
		opts EmbedOptions
	}

	var listCase = []testCase{{
		opts: EmbedOptions{
			ConvertOptions: ConvertOptions{
				Root: `testdata/in/`,
			},
			EmbedOptions: memfs.EmbedOptions{
				PackageName:    `mypackage`,
				VarName:        `memfsIn`,
				WithoutModTime: true,
				GoFileName:     `testdata/goembed/default.go`,
				CommentHeader: `// SPDX-License-Identifier: GPL-3.0-or-later
// SPDX-FileCopyrightText: 2026 Shulhan <ms@kilabit.info>
`,
			},
		},
	}}
	for _, tcase := range listCase {
		err := GoEmbed(tcase.opts)
		if err != nil {
			t.Fatal(err)
		}
		got, err := os.ReadFile(tcase.opts.EmbedOptions.GoFileName)
		if err != nil {
			t.Fatal(err)
		}
		exp, err := os.ReadFile(tcase.opts.EmbedOptions.GoFileName + `.exp`)
		if err != nil {
			t.Fatal(err)
		}
		test.Assert(t, tcase.tag, string(exp), string(got))
	}
}
