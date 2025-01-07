#  Welcome to ciigo
<!--{{{-->

`ciigo` is a program (as in command line interface, CLI) and a library (as
in Go package) to write static web server with embedded files using
[AsciiDoc](https://asciidoctor.org/docs/what-is-asciidoc/)
and
[Markdown](https://www.markdownguide.org/)
markup format.

As showcase, the following websites are build using ciigo,

* [kilabit.info](https://kilabit.info) -
  [Source code](https://git.sr.ht/~shulhan/kilabit.info).

* [golang-id.org](https://golang-id.org) -
  [Source code](https://git.sr.ht/~shulhan/golang-id-web).


Features,

* Simple. There is only three main commands: `convert`, `embed`, and
  `serve`.

* No themes, built your own template and style.
  Except for default
  [HTML](https://git.sr.ht/~shulhan/ciigo/tree/main/item/template_index_html.go)
  template and embedded
  [CSS](https://git.sr.ht/~shulhan/ciigo/tree/main/item/embedded_css.go),
  there is no other selection of theme that you can "pre-installed".

* No layout, you are free to structure your sites, using normal directory
  and file layout.
  Directory that contains "index.adoc" or "index.md" will be served as is.
  For example "/my/journal/index.adoc" can be accessed as "/my/journal/" or
  "/my/journal/index.html".

<!--}}}-->

##  ciigo as CLI

ciigo as CLI can convert, generate, and/or serve a directory that contains
markup files, as HTML files.

### Installation
<!--{{{-->

The installation require the `git` and Go tools which you can download it
[here](https://go.dev/dl).

First, clone this repository somewhere, let say under your `$HOME/src`
directory,

```
$ mkdir -p $HOME/src/
$ cd $HOME/src
$ git clone https://git.sr.ht/~shulhan/ciigo
```

Run "go install" on ./cmd/ciigo,

```
$ cd $HOME/src/ciigo
$ go install ./cmd/ciigo
```

This will install the `ciigo` into your `$GOBIN`.
If you did not know where `$GOBIN` is, run "go env" and check for `GOBIN`
environment variable.
It usually in `$HOME/go/bin`.

<!--}}}-->
### Usage
<!--{{{-->

```
ciigo [-template <file>] [-exclude <regex>] convert <dir>
```
> Scan the "dir" recursively to find markup files (.adoc or .md) and
  convert them into HTML files.
  The "-template" file is optional, default to embedded HTML template.


```
ciigo [-template <file>] [-exclude <regex>] [-out <file>] \
	[-package-name <string>] [-var-name <string>] \
	embed <dir>
```
> Convert all markup files inside directory "dir" recursively and then embed
  them into ".go" source file.
  The output file is optional, default to "ciigo_static.go" in current
  directory.
  The package name default to main.
  The variable name default to memFS.


```
ciigo help
```
> Print the usage.


```
ciigo [-template <file>]  [-exclude <regex>] [-address <ip:port>] \
    serve <dir>
```
> Serve all files inside directory "dir" using HTTP server, watch
  changes on markup files and convert them to HTML files automatically.
  If the address is not set, its default to ":8080".


```
ciigo version
```
> Print the current ciigo version.

<!--}}}-->
### Example
<!--{{{-->

In this repository, we have an "_example" directory that we can try using
`ciigo`.

```
$ tree _example
_example/
├── custom.css
├── favicon.ico
├── html.tmpl
├── index.adoc
├── index.css
└── sub
    └── index.adoc
```

First, lets convert all AsciiDoc files (.adoc) inside the "_example"
directory recursively,

```
$ ciigo convert ./_example
2025/01/06 19:17:07 convertFileMarkups: converting _example/sub/index.adoc
2025/01/06 19:17:07 convertFileMarkups: converting _example/index.adoc
$
```

Then serve it under HTTP server,

```
$ ciigo serve ./_example/
2025/01/06 19:17:47 ciigo: starting HTTP server at http://127.0.0.1:6320 for "./_example"

```

While the program still running, open the
[http://127.0.0.1:6320](http://127.0.0.1:6320) in your
browser.
Notice that we have "sub" directory.
Lets open
[http://127.0.0.1:6320/sub/](http://127.0.0.1:6320/sub/)
in your browser too.
As we can see, the style of "/sub/" page is different because they use
custom style defined using
```
:stylesheet: /custom.css
```
inside the "/sub/index.adoc" file.
This is one of powerful feature of ciigo, where different page can have
their own style.

Lets convert again but now using new "-template",
```
$ ciigo -template _example/html.tmpl
$
```
Nothing happened, why?
Because none of the markup files is newer than our previous generated HTML.
To force converting all markup files, we can update the modification time of
our template file,

```
$ touch _example/html.tmpl
$
```
or delete all the HTML files,
```
$ find _example/ -name "*.html" -delete
$
```
and then run the convert command again.
```
$ ciigo -template ./_example/html.tmpl convert ./_example/
2025/01/06 19:28:17 convertFileMarkups: converting _example/sub/index.adoc
2025/01/06 19:28:17 convertFileMarkups: converting _example/index.adoc
$
```

Run the `serve` command again to preview our new generated HTML files,
```
$ ciigo serve ./_example/
2025/01/06 19:36:07 ciigo: starting HTTP server at http://127.0.0.1:6320 for "./_example"

```
You can see now that the new template rendered, with "ciigo" in the top-left
as tile and a "Sub" menu at the top-right.

<!--}}}-->
### Writing content and viewing it live
<!--{{{-->

The "ciigo serve" command able to watch for new changes on markup files and
convert it automatically to HTML without need to run "convert" manually.

Lets run the serve command again on _example directory,
```
$ ciigo serve ./_example/
2025/01/06 19:46:54 ciigo: starting HTTP server at http://127.0.0.1:6320 for "./_example"

```

Create new directory "journal" under "_example",
```
$ mkdir -p ./_example/journal/
$
```
Inside the "journal" directory, create new file "index.adoc" with the
following content,

```
= My Journal

Hello world!
```

Go back to the browser and open
[http://127.0.0.1:6320/journal/](http://127.0.0.1:6320/journal/)
and we should see new HTML page generated with the above content.

Each time we refresh the browser, ciigo will check if the markup file is
updated and then convert it automatically to HTML and return it to browser.

Another way to trigger ciigo to rescan for new markup files is by creating
or updating file ".ciigo_rescan" inside the content directory.
In the above example, by creating file "_example/.ciigo_rescan".

Lets try updating the "journal/index.adoc" into
```
= My Journal

Hello ciigo!
```
and create or update the ".ciigo_rescan",
```
$ touch _example/.ciigo_rescan
$
```
and when we refresh the browser, we should see the page being updated.

<!--}}}-->
### Deployment
<!--{{{-->

Once we have write the content as we like, we can deploy the generated
HTML files and other assets files to the server.

For example, using `rsync(1)` we can deploy the "_example/" excluding the
hidden files (start with ".") and files with extension ".adoc", ".md", and
".tmpl"; by issuing the following command,
```
$ rsync --exclude='.*' \
    --exclude='*.adoc' --exclude='*.md' --exclude='*.tmpl' \
    --recursive --archive \
    _example/ \
    user@myserver:/srv/pub/
```

<!--}}}-->

##  ciigo as library
<!--{{{-->

Using ciigo as library means you do not need to install the ciigo program
itself, but you write a Go code by importing the ciigo module into your own
program.

This model gives flexibility.
You can adding custom functions, custom HTTP endpoints, and many more.

For an example on how to use ciigo as library see the code at
[internal/cmd/ciigo-example](https://git.sr.ht/~shulhan/ciigo/tree/main/internal/cmd/ciigo-example).


That's it, happy writing!

<!--}}}-->
##  Links
<!--{{{-->

<https://git.sr.ht/~shulhan/ciigo> - Link to the source code.

<https://lists.sr.ht/~shulhan/ciigo> - Link to development and discussion.

<https://todo.sr.ht/~shulhan/ciigo> - Link to submit an issue, feedback, or
request for new feature.

<https://pkg.go.dev/git.sr.ht/~shulhan/ciigo> - Go module documentation.

[Change log](/CHANGELOG.html) - Log of each releases.

<!--}}}-->
##  License
<!--{{{-->

This software is licensed under GPL 3.0 or later.

Copyright 2022 Shulhan <ms@kilabit.info>

This program is free software: you can redistribute it and/or modify it
under the terms of the GNU General Public License as published by the Free
Software Foundation, either version 3 of the License, or (at your option)
any later version.

This program is distributed in the hope that it will be useful, but WITHOUT
ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
FITNESS FOR A PARTICULAR PURPOSE.
See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with
this program.
If not, see <http://www.gnu.org/licenses/>.

<!--}}}-->
