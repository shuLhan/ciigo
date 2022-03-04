// SPDX-FileCopyrightText: 2020 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import "html/template"

var _embeddedCSS *template.CSS

func embeddedCSS() *template.CSS {
	if _embeddedCSS != nil {
		return _embeddedCSS
	}

	css := template.CSS(`
body {
	margin: 0;
	font-family: Arial, sans-serif;
	background-color: #fff;
	line-height: 1.3;
	text-align: center;
	color: #222;
}
pre,
code {
	font-family: Menlo, monospace;
	font-size: 0.875rem;
}
pre {
	line-height: 1.4;
	overflow-x: auto;
	background: #efefef;
	padding: 0.625rem;
	border-radius: 0.3125rem;
}
a {
	color: #007d9c;
	text-decoration: none;
}
a:hover {
	text-decoration: underline;
}

p,
li {
	max-width: 50rem;
	word-wrap: break-word;
}
li p {
	margin: 2px;
}
p,
pre,
ul,
ol {
	margin: 1.25rem;
}

h1,
h2,
h3,
h4 {
	margin: 1.25rem 0 1.25rem;
	padding: 0;
	color: #007d9c;
	font-weight: bold;
}
h1 {
	font-size: 1.75rem;
	line-height: 1;
}
h1 .text-muted {
	color: #777;
}
h2 {
	clear: right;
	font-size: 1.25rem;
	background: #e0ebf5;
	padding: 0.5rem;
	line-height: 1.25;
	font-weight: normal;
	overflow: auto;
	overflow-wrap: break-word;
}
h2 a {
	font-weight: bold;
}
h3 {
	font-size: 1.25rem;
	line-height: 1.25;
	overflow: auto;
	overflow-wrap: break-word;
}
h3,
h4 {
	margin: 1.25rem 0.3125rem;
}
h4 {
	font-size: 1rem;
}

h2 > span,
h3 > span {
	float: right;
	margin: 0 25px 0 0;
	font-weight: normal;
	color: #5279c7;
}

dl {
	margin: 1.25rem;
}
dd {
	margin: 0 0 0 1.25rem;
}
dl,
dd {
	font-size: 0.875rem;
}

/**
 * Custom classes for pages
 */

.topbar {
	background: #e0ebf5;
	height: 4rem;
	overflow: hidden;
}

.topbar .top-heading,
.topbar .menu {
	padding: 1.313rem 0;
	font-size: 1.25rem;
	font-weight: normal;
}
.topbar .top-heading {
	float: left;
}
.topbar .top-heading a {
	color: #222;
	text-decoration: none;
}

.top-heading .header-logo {
	height: 2rem;
	width: 5.125rem;
}

.topbar .menu {
	float: right;
}
.topbar .menu a {
	margin: 0.625rem 0.125rem;
	padding: 0.625rem;
	color: white;
	background: #007d9c;
	border: 0.0625rem solid #007d9c;
	border-radius: 5px;
}
.topbar .menu form {
	display: inline-block;
}

.page {
	width: 100%;
}

.page > .container,
.topbar > .container,
.footer > .container {
	margin-left: auto;
	margin-right: auto;
	padding: 0 1.25rem;
	max-width: 59.38rem;
}

.page > .container {
	text-align: left;
}

.container .meta {
	font-style: italic;
	margin: 1.25rem;
}

.footer {
	text-align: center;
	color: #666;
	font-size: 0.875rem;
	margin: 2.5rem 0;
}

.ulist li .paragraph {
	margin-bottom: 1em;
}

.ulist li .paragraph {
	margin-bottom: 1em;
}

/** Custom classes */
#toctitle {
	display: none;
}
#toc li {
	list-style: none;
}
#toc ul .sectlevel1 {
	padding: 0px;
	margin: 4px;
}
#toc ul .sectlevel2 {
	margin: 4px;
}

@media screen and (max-width: 992px) {
	#toc {
		all: unset;
	}
}
`)
	_embeddedCSS = &css
	return _embeddedCSS
}
