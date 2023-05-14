#  Welcome to ciigo

`ciigo` is a library and a program to write static web server with embedded
files using
[AsciiDoc](https://asciidoctor.org/docs/what-is-asciidoc/)
and
[Markdown](https://www.markdownguide.org/)
markup format.

As showcase, the following websites are build using ciigo,

* [kilabit.info](https://kilabit.info)
* [golang-id.org](https://golang-id.org)


##  ciigo as CLI

ciigo as CLI can convert, generate, and/or serve a directory that contains
markup files, as HTML files.

###  Usage

```
$ ciigo [-template <file>] [-exclude <regex>] convert <dir>
```

Scan the "dir" recursively to find markup files (.adoc) and convert them into
HTML files.
The template "file" is optional, default to embedded HTML template.

```
$ ciigo [-template <file>] [-exclude <regex>] [-out <file>] generate <dir>
```

Convert all markup files inside directory "dir" recursively and then
embed them into ".go" source file.
The output file is optional, default to "ciigo_static.go" in current
directory.

```
$ ciigo [-template <file>] [-exclude <regex>] [-address <ip:port>] serve <dir>
```

Serve all files inside directory "dir" using HTTP server, watch
changes on markup files and convert them to HTML files automatically.
If the address is not set, its default to ":8080".


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
	"github.com/shuLhan/share/lib/memfs"
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
If you need to use another port, you can change it at `cmd/mysite/main.go`.

That's it!


##  Limitations and known bugs

Using symlink on ".adoc" file inside Root directory that reference file
outside of Root is not supported, yet.


##  Links

[Ciigo repository](https://git.sr.ht/shulhan/ciigo).

[Go module
documentation](https://pkg.go.dev/git.sr.ht/https://git.sr.ht/shulhan/ciigoshulhan/ciigo).


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
