// SPDX-FileCopyrightText: 2020 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

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
