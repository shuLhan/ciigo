// SPDX-FileCopyrightText: 2021 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/shuLhan/share/lib/test"
)

var (
	testWatcher  *watcher
	testFileAdoc string
	testAdocFile *os.File
)

func TestWatcher(t *testing.T) {
	var (
		testDir     = "testdata/watcher"
		convertOpts = ConvertOptions{
			Root: testDir,
		}

		converter *Converter
		err       error
	)

	err = os.RemoveAll(testDir)
	if err != nil {
		t.Logf(err.Error())
	}

	err = os.MkdirAll(testDir, 0700)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		os.RemoveAll(testDir)
	})

	converter, err = NewConverter("testdata/html.tmpl")
	if err != nil {
		t.Fatal(err)
	}

	err = convertOpts.init()
	if err != nil {
		t.Fatal(err)
	}

	testWatcher, err = newWatcher(converter, &convertOpts)
	if err != nil {
		t.Fatal(err)
	}

	err = testWatcher.start()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("createAdocFile", testCreate)
	t.Run("updateAdocFile", testUpdate)
	t.Run("deleteAdocFile", testDelete)
}

func testCreate(t *testing.T) {
	var (
		err error
	)

	testFileAdoc = filepath.Join(testWatcher.dir, "index.adoc")
	testAdocFile, err = os.Create(testFileAdoc)
	if err != nil {
		t.Fatal(err)
	}

	got := waitChanges()

	test.Assert(t, "New adoc file created", testFileAdoc, got.path)

	expBody := `
<div id="header">
<div class="details">
</div>
</div>
<div id="content">
<div id="preamble">
<div class="sectionbody">
</div>
</div>
</div>`
	gotBody := removeFooter(string(got.fhtml.Body))
	test.Assert(t, "HTML body", expBody, gotBody)
}

func testUpdate(t *testing.T) {
	_, err := testAdocFile.WriteString("= Hello")
	if err != nil {
		t.Fatal(err)
	}
	err = testAdocFile.Sync()
	if err != nil {
		t.Fatal(err)
	}

	got := waitChanges()
	test.Assert(t, "adoc file updated", testFileAdoc, got.path)

	expBody := `
<div id="header">
<h1>Hello</h1>
<div class="details">
</div>
</div>
<div id="content">
<div id="preamble">
<div class="sectionbody">
</div>
</div>
</div>`
	gotBody := removeFooter(string(got.fhtml.Body))
	test.Assert(t, "HTML body", expBody, gotBody)
}

func testDelete(t *testing.T) {
	err := testAdocFile.Close()
	if err != nil {
		t.Fatal(err)
	}

	err = os.Remove(testFileAdoc)
	if err != nil {
		t.Fatal(err)
	}

	got := waitChanges()
	test.Assert(t, "adoc file updated", testFileAdoc, got.path)

	_, gotIsExist := testWatcher.fileMarkups[testFileAdoc]
	test.Assert(t, "adoc file deleted", false, gotIsExist)
}

// removeFooter remove the footer from generated HTML. The footer is 4 lines
// at the bottom.
func removeFooter(in string) string {
	lines := strings.Split(in, "\n")
	n := len(lines)
	if n > 5 {
		lines = lines[:n-5]
	}
	return strings.Join(lines, "\n")
}

func waitChanges() (fmarkup *fileMarkup) {
	var (
		ok bool
	)

	for {
		time.Sleep(5000)
		fmarkup, ok = testWatcher.changes.Pop().(*fileMarkup)
		if ok {
			break
		}
	}
	return fmarkup
}
