// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"testing"

	"git.sr.ht/~shulhan/pakakeh.go/lib/memfs"
	"git.sr.ht/~shulhan/pakakeh.go/lib/test"
)

func TestMain(m *testing.M) {
	var templateCSS = template.CSS(`body{}`)

	_embeddedCSS = &templateCSS

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
	type testCase struct {
		tag  string
		opts EmbedOptions
	}

	var (
		outDir = `testdata/goembed/out`

		tdata *test.Data
		err   error
	)

	tdata, err = test.LoadData(`testdata/goembed/GoEmbed_test.txt`)
	if err != nil {
		t.Fatal(err)
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
			},
		},
		tag: `default`,
	}}

	var (
		tcase testCase
		fname string
		fpath string
		got   []byte
	)
	for _, tcase = range listCase {
		// Set the output file name based on tag.
		fname = tcase.tag + `.go`
		fpath = filepath.Join(outDir, fname)
		tcase.opts.EmbedOptions.GoFileName = filepath.Join(outDir, fname)

		err = GoEmbed(tcase.opts)
		if err != nil {
			t.Fatal(err)
		}

		got, err = os.ReadFile(fpath)
		if err != nil {
			t.Fatal(err)
		}

		test.Assert(t, tcase.tag, string(tdata.Output[fname]), string(got))
	}
}
