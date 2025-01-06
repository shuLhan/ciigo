#  Welcome to ciigo

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

* Simple. There is only five commands: convert, embed, help, serve, and
  version.

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


##  ciigo as CLI

ciigo as CLI can convert, generate, and/or serve a directory that contains
markup files, as HTML files.


###  Installation

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


###  Usage

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


### Example

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
our template file and run convert again.

```
$ touch _example/html.tmpl
$ ciigo -template ./_example/html.tmpl convert ./_example/
2025/01/06 19:28:17 convertFileMarkups: converting _example/sub/index.adoc
2025/01/06 19:28:17 convertFileMarkups: converting _example/index.adoc
$
```

Run the `serve` command again,
```
$ ciigo serve ./_example/
2025/01/06 19:36:07 ciigo: starting HTTP server at http://127.0.0.1:6320 for "./_example"

```
You can see now that the new template rendered, with "ciigo" in the top-left
and "Sub" menu at the top-right.


### Writing content and viewing it live

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


### Deployment

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


##  ciigo as library

This section describe step by step instructions on how to build and create
pages to be viewed for local development using `ciigo`.

First, clone the `ciigo` repository.
Let says that we have cloned the `ciigo` repository into
`$HOME/go/src/git.sr.ht/~shulhan/ciigo`.

Create new Go repository for building a website.
For example, in directory `$HOME/go/src/remote.tld/user/mysite`.
Replace "remote.tld/user/mysite" with your private or public repository.

```
$ mkdir -p $HOME/go/src/remote.tld/user/mysite
$ cd $HOME/go/src/remote.tld/user/mysite
```

Initialize the Go module,

```
$ go mod init remote.tld/user/mysite
```

Create directories for storing our content and a package binary.

```
$ mkdir -p cmd/mysite
$ mkdir -p _contents
```

Copy the example of stylesheet and HTML template from `ciigo` repository,

```
$ cp $HOME/go/src/git.sr.ht/~shulhan/ciigo/_example/index.css ./_contents/
$ cp $HOME/go/src/git.sr.ht/~shulhan/ciigo/_example/html.tmpl ./_contents/
```

Create the main Go code inside `cmd/mysite`,

```
package main

import (
	"git.sr.ht/~shulhan/ciigo"
	"git.sr.ht/~shulhan/pakakeh.go/lib/memfs"
)

var mysiteFS *memfs.MemFS

func main() {
	opts := &ciigo.ServeOptions{
		ConvertOptions: ciigo.ConvertOptions{
			Root: "_contents",
			HtmlTemplate: "_contents/html.tmpl",
		},
		Address: ":8080",
		Mfs: mysiteFS,
	}
	err := ciigo.Serve(opts)
	if err != nil {
		log.Fatal(err)
	}
}
```

Create a new markup file `index.adoc` inside the "_contents" directory.
Each directory, or sub directory, should have `index.adoc` to be able to
accessed by browser,

```
=  Test

Hello, world!
```

Now run the `./cmd/mysite` with `DEBUG` environment variable set to non-zero,

```
$ DEBUG=1 go run ./cmd/mysite
```

Any non zero value on `DEBUG` environment signal the running program to watch
changes in ".adoc" files inside "_contents" directory and serve the generated
HTML directly.

Open the web browser at `localhost:8080` to view the generated HTML.
You should see "Hello, world!" as the main page.

Thats it!

Create or update any ".adoc" files inside "_contents" directory, the
program will automatically generated the HTML file.
Refresh the web browser to load the new generated file.


###  Deployment

First, we need to make sure that all markup files inside "_contents" are
converted to HTML and embed it into the static Go code.

Create another Go source code, lets save it in `internal/generate.go` with the
following content,

```
package main

import (
	"log"

	"git.sr.ht/~shulhan/ciigo"
)

func main() {
	opts := &ciigo.EmbedOptions{
		ConvertOptions: ciigo.ConvertOptions{
			Root:           "_contents",
			HtmlTemplate:   "_contents/html.tmpl",
		},
		EmbedOptions: memfs.EmbedOptions{
			PackageName: "main",
			VarName:     "mysiteFS",
			GoFileName:  "cmd/mysite/static.go",
		},
	}
	err := ciigo.GoEmbed(opts)
	if err != nil {
		log.Fatal(err)
	}
}
```

And then run,

```
$ go run ./internal
```

The above command will generate Go source code `cmd/mysite/static.go` that
embed all files inside the "_contents" directory.

Second, build the web server that serve static contents in `static.go`,

```
$ go build cmd/mysite
```

Third, test the web server by running the program and opening `localhost:8080`
on web browser,

```
$ ./mysite
```

Finally, deploy the program to your server.

NOTE: By default, server will listen on address `0.0.0.0` at port `8080`.
If we need to use another port, we can change it at `cmd/mysite/main.go`.

That's it!


##  Links

<https://git.sr.ht/~shulhan/ciigo> - Link to the source code.

<https://lists.sr.ht/~shulhan/ciigo> - Link to development and discussion.

<https://todo.sr.ht/~shulhan/ciigo> - Link to submit an issue, feedback, or
request for new feature.

<https://pkg.go.dev/git.sr.ht/~shulhan/ciigo> - Go module documentation.

[Change log](/CHANGELOG.html) - Log of each releases.


##  License

This software is licensed under GPL 3.0 or later.

Copyright 2022 Shulhan <ms@kilabit.info>

This program is free software: you can redistribute it and/or modify it under
the terms of the GNU General Public License as published by the Free Software
Foundation, either version 3 of the License, or (at your option) any later
version.

This program is distributed in the hope that it will be useful, but WITHOUT
ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
FOR A PARTICULAR PURPOSE.
See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with
this program.
If not, see <http://www.gnu.org/licenses/>.
