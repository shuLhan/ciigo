<!--
// SPDX-FileCopyrightText: 2020 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later
--!>
#  Welcome to ciigo

[![PkgGoDev](https://pkg.go.dev/badge/git.sr.ht/~shulhan/ciigo)](https://pkg.go.dev/git.sr.ht/~shulhan/ciigo)

`ciigo` is a library and a program to write static web server with embedded
files using
[asciidoc](https://asciidoctor.org/docs/what-is-asciidoc/)
markup format.


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
	"git.sr.ht/~shulhan/ciigo"
)

func main() {
	opts := &ciigo.EmbedOptions{
		ConvertOptions: ciigo.ConvertOptions{
			Root:           "./_contents",
			HtmlTemplate:   "_contents/html.tmpl",
		},
		EmbedOptions: memfs.EmbedOptions{
			PackageName: "main",
			VarName:     "mysiteFS",
			GoFileName:  "cmd/mysite/static.go",
		},
	}
	ciigo.GoEmbed(opts)
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

*NOTE:* By default, server will listen on address `0.0.0.0` at port `8080`.
If you need to use another port, you can change it at `cmd/mysite/main.go`.

That's it!

##  Limitations and Known Bugs

`ciigo` will not handle automatic certificate (e.g. using LetsEncrypt), its
up to the user how the certificate are gathered, generated, or served.

Using symlink on ".adoc" file inside Root directory that reference file
outside of Root is not supported, yet.


## License

```
Copyright 2019-2021, M. Shulhan (ms@kilabit.info).
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of copyright holder nor the names of its contributors may
   be used to endorse or promote products derived from this software without
   specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDERS OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```
