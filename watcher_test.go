// SPDX-FileCopyrightText: 2021 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"bytes"
	"os"
	"path/filepath"
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
		testDir     = `testdata/watcher`
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

	converter, err = NewConverter(`testdata/html.tmpl`)
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

	var tdata *test.Data

	tdata, err = test.LoadData(`testdata/watcher_test.txt`)
	if err != nil {
		t.Fatal(err)
	}

	t.Run(`createAdocFile`, testCreate)
	t.Run(`updateAdocFile`, testUpdate)
	t.Run(`deleteAdocFile`, testDelete)

	var pathFileMarkdown = filepath.Join(testWatcher.dir, `test.md`)

	t.Run(`testMarkdownCreate`, func(tt *testing.T) {
		testMarkdownCreate(tt, tdata, pathFileMarkdown)
	})
	t.Run(`testMarkdownUpdate`, func(tt *testing.T) {
		testMarkdownUpdate(tt, tdata, pathFileMarkdown)
	})
	t.Run(`testMarkdownDelete`, func(tt *testing.T) {
		testMarkdownDelete(tt, pathFileMarkdown)
	})
}

func testCreate(t *testing.T) {
	var (
		got     *FileMarkup
		err     error
		expBody string
		gotBody []byte
	)

	testFileAdoc = filepath.Join(testWatcher.dir, `index.adoc`)

	// Let the OS sync the file system before we create new file,
	// otherwise the modtime for fs.Root does not changes.
	time.Sleep(1 * time.Second)

	testAdocFile, err = os.Create(testFileAdoc)
	if err != nil {
		t.Fatal(err)
	}

	got = waitChanges()

	test.Assert(t, `New adoc file created`, testFileAdoc, got.path)

	expBody = `<!DOCTYPE>
<html>
<head><title></title></head>
<body><div id="header">
</div>
<div id="content">
</div>`

	gotBody, err = os.ReadFile(got.pathHtml)
	if err != nil {
		t.Fatal(err)
	}

	gotBody = removeFooter(gotBody)
	test.Assert(t, `HTML body`, expBody, string(gotBody))
}

func testUpdate(t *testing.T) {
	var (
		err     error
		expBody string
		gotBody []byte
		got     *FileMarkup
	)

	_, err = testAdocFile.WriteString(`= Hello`)
	if err != nil {
		t.Fatal(err)
	}
	err = testAdocFile.Sync()
	if err != nil {
		t.Fatal(err)
	}

	got = waitChanges()
	test.Assert(t, `adoc file updated`, testFileAdoc, got.path)

	expBody = `<!DOCTYPE>
<html>
<head><title>Hello</title></head>
<body><div id="header">
<h1>Hello</h1>
</div>
<div id="content">
</div>`

	gotBody, err = os.ReadFile(got.pathHtml)
	if err != nil {
		t.Fatal(err)
	}

	gotBody = removeFooter(gotBody)

	test.Assert(t, `HTML body`, expBody, string(gotBody))
}

func testDelete(t *testing.T) {
	var (
		err        error
		got        *FileMarkup
		gotIsExist bool
	)

	err = testAdocFile.Close()
	if err != nil {
		t.Fatal(err)
	}

	err = os.Remove(testFileAdoc)
	if err != nil {
		t.Fatal(err)
	}

	got = waitChanges()
	test.Assert(t, `adoc file updated`, testFileAdoc, got.path)

	_, gotIsExist = testWatcher.fileMarkups[testFileAdoc]
	test.Assert(t, `adoc file deleted`, false, gotIsExist)
}

func testMarkdownCreate(t *testing.T, tdata *test.Data, pathFileMarkdown string) {
	var (
		body = tdata.Input[`create.md`]

		got     *FileMarkup
		err     error
		expBody string
		gotBody []byte
	)

	// Let the OS sync the file system before we create new file,
	// otherwise the modtime for fs.Root does not changes.
	time.Sleep(1 * time.Second)

	err = os.WriteFile(pathFileMarkdown, body, 0600)
	if err != nil {
		t.Fatal(err)
	}

	got = waitChanges()

	test.Assert(t, `New md file created`, pathFileMarkdown, got.path)

	gotBody, err = os.ReadFile(got.pathHtml)
	if err != nil {
		t.Fatal(err)
	}
	gotBody = removeFooter(gotBody)

	expBody = string(tdata.Output[`create.md.html`])
	test.Assert(t, `HTML body`, expBody, string(gotBody))
}

func testMarkdownUpdate(t *testing.T, tdata *test.Data, pathFileMarkdown string) {
	var (
		body = tdata.Input[`update.md`]

		got     *FileMarkup
		err     error
		expBody string
		gotBody []byte
	)

	// Let the OS sync the file system before we create new file,
	// otherwise the modtime for fs.Root does not changes.
	time.Sleep(1 * time.Second)

	err = os.WriteFile(pathFileMarkdown, body, 0600)
	if err != nil {
		t.Fatal(err)
	}

	got = waitChanges()

	test.Assert(t, `changes path`, pathFileMarkdown, got.path)

	gotBody, err = os.ReadFile(got.pathHtml)
	if err != nil {
		t.Fatal(err)
	}
	gotBody = removeFooter(gotBody)

	expBody = string(tdata.Output[`update.md.html`])
	test.Assert(t, `HTML body`, expBody, string(gotBody))
}

func testMarkdownDelete(t *testing.T, pathFileMarkdown string) {
	var (
		err        error
		got        *FileMarkup
		gotIsExist bool
	)

	err = os.Remove(pathFileMarkdown)
	if err != nil {
		t.Fatal(err)
	}

	got = waitChanges()
	test.Assert(t, `md file updated`, pathFileMarkdown, got.path)

	_, gotIsExist = testWatcher.fileMarkups[pathFileMarkdown]
	test.Assert(t, `md file deleted`, false, gotIsExist)
}

// removeFooter remove the footer from generated HTML since its contains date
// and time that changes during test.
func removeFooter(in []byte) (out []byte) {
	var (
		lines = bytes.Split(in, []byte("\n"))
		n     = len(lines)
	)
	if n > 7 {
		lines = lines[:n-7]
	}
	out = bytes.Join(lines, []byte("\n"))
	return out
}

func waitChanges() (fmarkup *FileMarkup) {
	var (
		ok bool
	)

	for {
		fmarkup, ok = testWatcher.changes.Pop().(*FileMarkup)
		if ok {
			break
		}
	}
	return fmarkup
}
