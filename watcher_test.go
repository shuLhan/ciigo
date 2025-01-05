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
	testWatcher.opts.FileWatcherOptions.Interval = 50 * time.Millisecond

	err = testWatcher.start()
	if err != nil {
		t.Fatal(err)
	}

	var tdata *test.Data

	tdata, err = test.LoadData(`testdata/watcher_test.txt`)
	if err != nil {
		t.Fatal(err)
	}

	var pathFileMarkup = filepath.Join(testWatcher.opts.Root, `test.adoc`)

	t.Run(`testAdocCreate`, func(tt *testing.T) {
		testAdocCreate(tt, testWatcher, tdata, pathFileMarkup)
	})
	t.Run(`testAdocUpdate`, func(tt *testing.T) {
		testAdocUpdate(tt, testWatcher, tdata, pathFileMarkup)
	})
	t.Run(`testAdocDelete`, func(tt *testing.T) {
		testAdocDelete(tt, testWatcher, pathFileMarkup)
	})

	pathFileMarkup = filepath.Join(testWatcher.opts.Root, `test.md`)

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

	err = os.WriteFile(pathFile, expBody, 0600)
	if err != nil {
		t.Fatal(err)
	}
	testWatcher.watchDir.ForceRescan()

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

	err = os.WriteFile(pathFile, expBody, 0600)
	if err != nil {
		t.Fatal(err)
	}
	testWatcher.watchDir.ForceRescan()

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
	testWatcher.watchDir.ForceRescan()

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

	err = os.WriteFile(pathFileMarkdown, body, 0600)
	if err != nil {
		t.Fatal(err)
	}
	testWatcher.watchDir.ForceRescan()

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

	err = os.WriteFile(pathFileMarkdown, body, 0600)
	if err != nil {
		t.Fatal(err)
	}
	testWatcher.watchDir.ForceRescan()

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

	err = os.Remove(pathFileMarkdown)
	if err != nil {
		t.Fatal(err)
	}
	testWatcher.watchDir.ForceRescan()

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

func TestWatcherGetFileMarkupByHTML(t *testing.T) {
	var w = watcher{
		fileMarkups: map[string]*FileMarkup{
			`/markup/with/adoc/file.adoc`: &FileMarkup{
				kind: markupKindAdoc,
			},
			`/markup/with/md/file.md`: &FileMarkup{
				kind: markupKindMarkdown,
			},
		},
	}

	var listCase = []struct {
		expFileMarkup *FileMarkup
		fileHTML      string
		expIsNew      bool
	}{{
		fileHTML: `/notexist.html`,
	}, {
		fileHTML:      `/markup/with/adoc/file.html`,
		expFileMarkup: w.fileMarkups[`/markup/with/adoc/file.adoc`],
	}, {
		fileHTML:      `/markup/with/adoc/file.HTML`,
		expFileMarkup: w.fileMarkups[`/markup/with/adoc/file.adoc`],
	}, {
		fileHTML:      `/markup/with/md/file.HTML`,
		expFileMarkup: w.fileMarkups[`/markup/with/md/file.md`],
	}}

	var (
		gotFileMarkup *FileMarkup
		gotIsNew      bool
	)
	for _, tcase := range listCase {
		gotFileMarkup, gotIsNew = w.getFileMarkupByHTML(tcase.fileHTML)
		test.Assert(t, tcase.fileHTML, tcase.expFileMarkup, gotFileMarkup)
		test.Assert(t, tcase.fileHTML+` isNew`, tcase.expIsNew, gotIsNew)
	}
}
