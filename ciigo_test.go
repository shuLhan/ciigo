// SPDX-FileCopyrightText: 2022 M. Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"regexp"
	"sort"
	"testing"

	"github.com/shuLhan/share/lib/test"
)

func TestListFileMarkups(t *testing.T) {
	var (
		dir = "testdata"
	)

	cases := []struct {
		excRegex string
		exp      []string
	}{{
		excRegex: `(ex)/.*`,
		exp: []string{
			"testdata/in/clu/de/file.adoc",
		},
	}, {
		excRegex: `(in|ex)/.*`,
	}, {
		excRegex: `de`,
	}, {
		excRegex: `file$`,
		exp: []string{
			"testdata/ex/clu/de/file.adoc",
			"testdata/in/clu/de/file.adoc",
		},
	}}

	for _, c := range cases {
		excre := regexp.MustCompile(c.excRegex)

		list, err := listFileMarkups(dir, []*regexp.Regexp{excre})
		if err != nil {
			t.Fatal(err)
		}

		got := make([]string, 0, len(list))
		for k := range list {
			got = append(got, k)
		}

		sort.Strings(got)

		test.Assert(t, "list", c.exp, got)
	}
}
