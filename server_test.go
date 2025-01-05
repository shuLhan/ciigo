// SPDX-FileCopyrightText: 2024 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo_test

import (
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"git.sr.ht/~shulhan/ciigo"
	"git.sr.ht/~shulhan/pakakeh.go/lib/test"
	"git.sr.ht/~shulhan/pakakeh.go/lib/test/httptest"
)

func TestCiigoOnGet(t *testing.T) {
	var tdata *test.Data
	var err error

	tdata, err = test.LoadData(`testdata/onGet_test.txt`)
	if err != nil {
		t.Fatal(err)
	}

	var dirRoot = t.TempDir()

	// Create one adoc, one md, and one html file.
	for _, name := range []string{`one.adoc`, `two.md`, `three.html`} {
		var file = filepath.Join(dirRoot, name)
		err = os.WriteFile(file, tdata.Input[name], 0600)
		if err != nil {
			t.Fatal(err)
		}
	}

	var cigo = ciigo.Ciigo{}

	var serveOpts = ciigo.ServeOptions{
		Address: `127.0.0.1:11083`,
		ConvertOptions: ciigo.ConvertOptions{
			Root:         dirRoot,
			HTMLTemplate: `testdata/onGet_template.gohtml`,
		},
	}
	err = cigo.InitHTTPServer(serveOpts)
	if err != nil {
		t.Fatal(err)
	}

	var redactLastUpdated = regexp.MustCompile(`Last updated.*`)

	var listCase = []struct {
		desc    string
		req     httptest.SimulateRequest
		expBody []byte
	}{{
		desc: `GET /one.html`,
		req: httptest.SimulateRequest{
			Method: http.MethodGet,
			Path:   `/one.html`,
		},
		expBody: tdata.Output[`one.html`],
	}, {
		desc: `GET /two.html`,
		req: httptest.SimulateRequest{
			Method: http.MethodGet,
			Path:   `/two.html`,
		},
		expBody: tdata.Output[`two.html`],
	}, {
		desc: `GET /three.html`,
		req: httptest.SimulateRequest{
			Method: http.MethodGet,
			Path:   `/three.html`,
		},
		expBody: tdata.Output[`three.html`],
	}}
	var result *httptest.SimulateResult
	for _, tcase := range listCase {
		result, err = httptest.Simulate(cigo.HTTPServer.ServeHTTP,
			&tcase.req)
		if err != nil {
			t.Fatal(err)
		}
		var gotBody = redactLastUpdated.ReplaceAll(
			result.ResponseBody, []byte("[REDACTED]"))
		test.Assert(t, tcase.desc, string(tcase.expBody),
			string(gotBody))
	}

	// On next test, we create markup file for three.html.
	// The output from HTML should changes.

	t.Run(`On markup created for HTML`, func(t *testing.T) {
		var name = `three.adoc`
		var file = filepath.Join(dirRoot, name)
		err = os.WriteFile(file, tdata.Input[name], 0600)
		if err != nil {
			t.Fatal(err)
		}

		var req = httptest.SimulateRequest{
			Method: http.MethodGet,
			Path:   `/three.html`,
		}

		var result *httptest.SimulateResult
		result, err = httptest.Simulate(cigo.HTTPServer.ServeHTTP, &req)
		if err != nil {
			t.Fatal(err)
		}

		var expBody = tdata.Output[`new_three.html`]
		var gotBody = redactLastUpdated.ReplaceAll(
			result.ResponseBody, []byte("[REDACTED]"))
		test.Assert(t, `new_three.html`, string(expBody), string(gotBody))
	})

	t.Run(`On markup updated`, func(t *testing.T) {
		var file = filepath.Join(dirRoot, `one.adoc`)
		err = os.WriteFile(file, tdata.Input[`update_one.adoc`], 0600)
		if err != nil {
			t.Fatal(err)
		}

		var req = httptest.SimulateRequest{
			Method: http.MethodGet,
			Path:   `/one.html`,
		}
		var result *httptest.SimulateResult
		result, err = httptest.Simulate(cigo.HTTPServer.ServeHTTP, &req)
		if err != nil {
			t.Fatal(err)
		}

		var expBody = tdata.Output[`update_one.html`]
		var gotBody = redactLastUpdated.ReplaceAll(
			result.ResponseBody, []byte("[REDACTED]"))
		test.Assert(t, `body`, string(expBody), string(gotBody))
	})

	t.Run(`On new directory with markup`, func(t *testing.T) {
		var dirJournal = filepath.Join(dirRoot, `journal`)
		err = os.MkdirAll(dirJournal, 0755)
		if err != nil {
			t.Fatal(err)
		}

		var journalIndex = filepath.Join(dirJournal, `index.adoc`)
		err = os.WriteFile(journalIndex,
			tdata.Input[`/journal/index.adoc`], 0600)
		if err != nil {
			t.Fatal(err)
		}

		var req = httptest.SimulateRequest{
			Method: http.MethodGet,
			Path:   `/journal/`,
		}
		var result *httptest.SimulateResult
		result, err = httptest.Simulate(cigo.HTTPServer.ServeHTTP, &req)
		if err != nil {
			t.Fatal(err)
		}

		var expBody = tdata.Output[`/journal/index.html`]
		var gotBody = redactLastUpdated.ReplaceAll(
			result.ResponseBody, []byte("[REDACTED]"))
		test.Assert(t, `body`, string(expBody), string(gotBody))
	})

	t.Run(`On new directory request without slash`, func(t *testing.T) {
		var dirJournal2 = filepath.Join(dirRoot, `journal2`)
		err = os.MkdirAll(dirJournal2, 0755)
		if err != nil {
			t.Fatal(err)
		}

		var journalIndexAdoc = filepath.Join(dirJournal2, `index.adoc`)
		err = os.WriteFile(journalIndexAdoc,
			tdata.Input[`/journal2/index.adoc`], 0600)
		if err != nil {
			t.Fatal(err)
		}

		var req = httptest.SimulateRequest{
			Method: http.MethodGet,
			Path:   `/journal2`,
		}
		var result *httptest.SimulateResult
		result, err = httptest.Simulate(cigo.HTTPServer.ServeHTTP, &req)
		if err != nil {
			t.Fatal(err)
		}

		var expBody = tdata.Output[`/journal2/index.html`]
		var gotBody = redactLastUpdated.ReplaceAll(
			result.ResponseBody, []byte("[REDACTED]"))
		test.Assert(t, `body`, string(expBody), string(gotBody))
	})
}
