// SPDX-FileCopyrightText: 2021 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
	"time"

	"git.sr.ht/~shulhan/pakakeh.go/lib/test"
)

func TestWatcher(t *testing.T) {
	var (
		testDir     = t.TempDir()
		convertOpts = ConvertOptions{
			Root: testDir,
		}

		converter *Converter
		err       error
	)

	converter, err = NewConverter(``)
	if err != nil {
		t.Fatal(err)
	}

	err = convertOpts.init()
	if err != nil {
		t.Fatal(err)
	}

	var testWatcher *watcher

	testWatcher, err = newWatcher(converter, convertOpts)
	if err != nil {
		t.Fatal(err)
	}

	testWatcher.watchDir.Delay = 100 * time.Millisecond

	err = testWatcher.start()
	if err != nil {
		t.Fatal(err)
	}

	var tdata *test.Data

	tdata, err = test.LoadData(`testdata/watcher_test.txt`)
	if err != nil {
		t.Fatal(err)
	}

	var pathFileMarkup = filepath.Join(testWatcher.dir, `test.adoc`)

	t.Run(`testAdocCreate`, func(tt *testing.T) {
		testAdocCreate(tt, testWatcher, tdata, pathFileMarkup)
	})
	t.Run(`testAdocUpdate`, func(tt *testing.T) {
		testAdocUpdate(tt, testWatcher, tdata, pathFileMarkup)
	})
	t.Run(`testAdocDelete`, func(tt *testing.T) {
		testAdocDelete(tt, testWatcher, pathFileMarkup)
	})

	pathFileMarkup = filepath.Join(testWatcher.dir, `test.md`)

	t.Run(`testMarkdownCreate`, func(tt *testing.T) {
		testMarkdownCreate(tt, testWatcher, tdata, pathFileMarkup)
	})
	t.Run(`testMarkdownUpdate`, func(tt *testing.T) {
		testMarkdownUpdate(tt, testWatcher, tdata, pathFileMarkup)
	})
	t.Run(`testMarkdownDelete`, func(tt *testing.T) {
		testMarkdownDelete(tt, testWatcher, pathFileMarkup)
	})

	testWatcher.stop()
}

func testAdocCreate(t *testing.T, testWatcher *watcher, tdata *test.Data, pathFile string) {
	var (
		expBody = tdata.Input[`create.adoc`]

		got     *FileMarkup
		err     error
		gotBody []byte
	)

	// Let the OS sync the file system before we create new file,
	// otherwise the modtime for fs.Root does not changes.
	time.Sleep(100 * time.Millisecond)

	err = os.WriteFile(pathFile, expBody, 0600)
	if err != nil {
		t.Fatal(err)
	}

	got = testWatcher.waitChanges()

	test.Assert(t, `New adoc file created`, pathFile, got.path)

	gotBody, err = os.ReadFile(got.pathHTML)
	if err != nil {
		t.Fatal(err)
	}

	gotBody = removeFooter(gotBody, 15)
	expBody = tdata.Output[`create.adoc.html`]

	test.Assert(t, `HTML body`, string(expBody), string(gotBody))
}

func testAdocUpdate(t *testing.T, testWatcher *watcher, tdata *test.Data, pathFile string) {
	var (
		expBody = tdata.Input[`update.adoc`]

		got     *FileMarkup
		err     error
		gotBody []byte
	)

	// Let the OS sync the file system before we create new file,
	// otherwise the modtime for fs.Root does not changes.
	time.Sleep(100 * time.Millisecond)

	err = os.WriteFile(pathFile, expBody, 0600)
	if err != nil {
		t.Fatal(err)
	}

	got = testWatcher.waitChanges()

	test.Assert(t, `adoc file updated`, pathFile, got.path)

	gotBody, err = os.ReadFile(got.pathHTML)
	if err != nil {
		t.Fatal(err)
	}

	gotBody = removeFooter(gotBody, 15)
	expBody = tdata.Output[`update.adoc.html`]

	test.Assert(t, `HTML body`, string(expBody), string(gotBody))
}

func testAdocDelete(t *testing.T, testWatcher *watcher, pathFile string) {
	var (
		err        error
		got        *FileMarkup
		gotIsExist bool
	)

	err = os.Remove(pathFile)
	if err != nil {
		t.Fatal(err)
	}

	got = testWatcher.waitChanges()

	test.Assert(t, `adoc file updated`, pathFile, got.path)

	_, gotIsExist = testWatcher.fileMarkups[pathFile]
	test.Assert(t, `adoc file deleted`, false, gotIsExist)
}

func testMarkdownCreate(t *testing.T, testWatcher *watcher, tdata *test.Data, pathFileMarkdown string) {
	var (
		body = tdata.Input[`create.md`]

		got     *FileMarkup
		err     error
		expBody string
		gotBody []byte
	)

	// Let the OS sync the file system before we create new file,
	// otherwise the modtime for fs.Root does not changes.
	time.Sleep(100 * time.Millisecond)

	err = os.WriteFile(pathFileMarkdown, body, 0600)
	if err != nil {
		t.Fatal(err)
	}

	got = testWatcher.waitChanges()

	test.Assert(t, `New md file created`, pathFileMarkdown, got.path)

	gotBody, err = os.ReadFile(got.pathHTML)
	if err != nil {
		t.Fatal(err)
	}
	gotBody = removeFooter(gotBody, 8)

	expBody = string(tdata.Output[`create.md.html`])
	test.Assert(t, `HTML body`, expBody, string(gotBody))
}

func testMarkdownUpdate(t *testing.T, testWatcher *watcher, tdata *test.Data, pathFileMarkdown string) {
	var (
		body = tdata.Input[`update.md`]

		got     *FileMarkup
		err     error
		expBody string
		gotBody []byte
	)

	// Let the OS sync the file system before we create new file,
	// otherwise the modtime for fs.Root does not changes.
	time.Sleep(100 * time.Millisecond)

	err = os.WriteFile(pathFileMarkdown, body, 0600)
	if err != nil {
		t.Fatal(err)
	}

	got = testWatcher.waitChanges()

	test.Assert(t, `changes path`, pathFileMarkdown, got.path)

	gotBody, err = os.ReadFile(got.pathHTML)
	if err != nil {
		t.Fatal(err)
	}
	gotBody = removeFooter(gotBody, 8)

	expBody = string(tdata.Output[`update.md.html`])
	test.Assert(t, `HTML body`, expBody, string(gotBody))
}

func testMarkdownDelete(t *testing.T, testWatcher *watcher, pathFileMarkdown string) {
	var (
		err        error
		got        *FileMarkup
		gotIsExist bool
	)

	time.Sleep(100 * time.Millisecond)

	err = os.Remove(pathFileMarkdown)
	if err != nil {
		t.Fatal(err)
	}

	got = testWatcher.waitChanges()
	test.Assert(t, `md file updated`, pathFileMarkdown, got.path)

	_, gotIsExist = testWatcher.fileMarkups[pathFileMarkdown]
	test.Assert(t, `md file deleted`, false, gotIsExist)
}

// removeFooter remove the footer from generated HTML since its contains date
// and time that changes during test.
func removeFooter(in []byte, nlast int) (out []byte) {
	var (
		lines = bytes.Split(in, []byte("\n"))
		n     = len(lines)
	)
	if n > nlast {
		lines = lines[:n-nlast]
	}
	out = bytes.Join(lines, []byte("\n"))
	return out
}
