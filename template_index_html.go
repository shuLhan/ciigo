// Copyright 2020, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

const templateIndexHTML = `<!DOCTYPE html>
<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1" />
		<meta name="theme-color" content="#375EAB" />

		<title>{{.Title}}</title>

		<style>
		{{.EmbeddedCSS}}
		</style>

		{{- range .Styles}}
		<link rel="stylesheet" href="{{.}}" />
		{{- end}}
	</head>
	<body>
		<div class="topbar">
			<div class="container">
				<div class="top-heading">
					<a href="/">ciigo</a>
				</div>
				<div class="menu">
					<form class="item" action="/_internal/search">
						<input type="text" name="q" placeholder="Search" />
					</form>
				</div>
			</div>
		</div>

		<div class="page">
			<div class="container">
				{{.Body}}
			</div>
			<!-- .container -->
		</div>
		<!-- .page -->

		<div class="footer">
			Powered by <a
				href="https://git.sr.ht/~shulhan/ciigo"
			>
				ciigo
			</a>
		</div>
	</body>
</html>`
