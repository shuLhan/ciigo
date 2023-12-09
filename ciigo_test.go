// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"html/template"
	"os"
	"regexp"
	"sort"
	"testing"

	"github.com/shuLhan/share/lib/test"
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
