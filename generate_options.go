package ciigo

import "github.com/shuLhan/share/lib/memfs"

const (
	// DefaultRoot define default Root value for GenerateOptions.
	DefaultRoot = "."
)

//
// GenerateOptions define the options for calling Generate function.
//
type GenerateOptions struct {
	// Root directory where its content will be embedded into Go source
	// code.
	// Default to DefaultRoot if its empty.
	Root string

	// HTMLTemplate the HTML template to be used when converting asciidoc
	// file into HTML.
	// If empty it will default to use embedded HTML template.
	// See template_index_html.go for template format.
	HTMLTemplate string

	// GenPackageName the name of package in Go generated source code.
	// Default to memfs.DefaultGenPackageName if its empty.
	GenPackageName string

	// GenVarName the name of variable where all files in Root will be
	// stored.
	// Default to memfs.DefaultGenVarName if its empty.
	GenVarName string

	// GenGoFileName the file name of Go source code will be written.
	// Default to memfs.DefaultGenGoFileName if its empty.
	GenGoFileName string
}

func (opts *GenerateOptions) init() {
	if len(opts.Root) == 0 {
		opts.Root = DefaultRoot
	}
	if len(opts.GenPackageName) == 0 {
		opts.GenPackageName = memfs.DefaultGenPackageName
	}
	if len(opts.GenVarName) == 0 {
		opts.GenVarName = memfs.DefaultGenVarName
	}
	if len(opts.GenGoFileName) == 0 {
		opts.GenGoFileName = memfs.DefaultGenGoFileName
	}
}
