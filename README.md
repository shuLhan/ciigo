#  Welcome to ciigo

[![PkgGoDev](https://pkg.go.dev/badge/git.sr.ht/~shulhan/ciigo)](https://pkg.go.dev/git.sr.ht/~shulhan/ciigo)

`ciigo` is a library and a program to write static web server with embedded
files using
[asciidoc](https://asciidoctor.org/docs/what-is-asciidoc/)
markup format.


##  ciigo as library

For an up to date documentation of how to use the library see the
[Go documentation page](https://pkg.go.dev/git.sr.ht/~shulhan/ciigo).


##  ciigo as CLI

ciigo as CLI can convert, generate, and/or serve a directory that contains
markup files, as HTML files.

###  Usage

```
$ ciigo [-template <file>] convert <dir>
```

Scan the "dir" recursively to find markup files (.adoc) and convert them into
HTML files.
The template "file" is optional, default to embedded HTML template.

```
$ ciigo [-template <file>] [-out <file>] generate <dir>
```

Convert all markup files inside directory "dir" recursively and then
embed them into ".go" source file.
The output file is optional, default to "ciigo_static.go" in current
directory.

```
$ ciigo [-template <file>] [-address <ip:port>] serve <dir>
```

Serve all files inside directory "dir" using HTTP server, watch
changes on markup files and convert them to HTML files automatically.
If the address is not set, its default to ":8080".


##  Example

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

Create a Go source code in the root repository to generate all markup files
inside the "_contents" directory into HTML and dump all of their contents into
"static.go" file.
Lets named it `generate.go` with the following content,

```
//go:generate go run generate.go

package main

import (
	"git.sr.ht/~shulhan/ciigo"
)

func main() {
	opts := &GenerateOptions{
		Root:           "./_contents",
		HTMLTemplate:   "_contents/html.tmpl",
		GenPackageName: "main",
		GenVarName:     "mysiteFS",
		GenGoFileName:  "cmd/mysite/static.go",
	}
	ciigo.Generate(opts)
}
```

Create the main Go code inside `cmd/mysite`,

```
package main

import (
	"git.sr.ht/~shulhan/ciigo"
)

var mysiteFS *memfs.PathNode

func main() {
	ciigo.Serve(mysiteFS, "./_contents", ":8080", "_contents/html.tmpl")
}
```

Create a new markup file `index.adoc` inside the "_contents" directory.
Each directory, or sub directory, should have `index.adoc` to be able to
accessed by browser,

```
=  Test
:stylesheet: /_contents/index.css

Hello, world!
```

Run `go generate` to convert all files with extension `.adoc`
into HTML and embed it into Go source code at `./cmd/mysite/static.go`

```
$ go generate
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

Create or update any ".adoc" or ",md" files inside "_contents" directory, the
program will automatically generated the HTML file, but you still need to
refresh the web browser to load the new generated file.


###  Deployment

First, we need to make sure that all markup files inside "_contents" are
converted to HTML and regenerate the static Go code,

```
$ go generate
```

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


##  Limitations and Known Bugs

`ciigo` will not handle automatic certificate (e.g. using LetsEncrypt), its
up to administrator how the certificate are gathered or generated.

Using symlink on ".adoc" file inside `content` directory is not supported yet.


##  Resources

The source code for this software can be viewed at
https://git.sr.ht/~shulhan/ciigo
under custom [BSD license](/LICENSE).
