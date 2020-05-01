// Copyright 2020, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

const templateSearch = `
<h3> Search result </h3>
{{range $result := .}}
<h4>
<a href="{{$result.Path}}">{{$result.Path}}</a>
</h4>
	{{range $result.Snippets}}
	<p>... {{.}} ...</p>
	{{end}}
{{end}}`
