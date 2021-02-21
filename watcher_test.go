// Copyright 2021, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

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
	testDir := "testdata/watcher"

	err := os.MkdirAll(testDir, 0700)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		os.RemoveAll(testDir)
	})

	htmlg, err := newHTMLGenerator(nil, "testdata/html.tmpl", true)
	if err != nil {
		t.Fatal(err)
	}

	testWatcher, err = newWatcher(htmlg, testDir)
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

	test.Assert(t, "New adoc file created", testFileAdoc, got.path, true)

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
	test.Assert(t, "HTML body", expBody, gotBody, true)
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
	test.Assert(t, "adoc file updated", testFileAdoc, got.path, true)

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
	test.Assert(t, "HTML body", expBody, gotBody, true)
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
	test.Assert(t, "adoc file updated", testFileAdoc, got.path, true)

	_, gotIsExist := testWatcher.fileMarkups[testFileAdoc]
	test.Assert(t, "adoc file deleted", false, gotIsExist, true)
}

//
// removeFooter remove the footer from generated HTML. The footer is 4 lines
// at the bottom.
//
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
