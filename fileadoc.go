// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/bytesparadise/libasciidoc"
)

type fileAdoc struct {
	path     string                 // path contains full path to ".adoc" file.
	info     os.FileInfo            // info contains FileInfo of ".adoc" file.
	basePath string                 // basePath contains full path to file without ".adoc" extension.
	metadata map[string]interface{} // metadata contains ".adoc" metadata.
}

func newFileAdoc(path string, fi os.FileInfo) (adoc *fileAdoc, err error) {
	if len(path) == 0 {
		return nil, fmt.Errorf("ciigo: newFileAdoc: empty path")
	}
	if fi == nil {
		fi, err = os.Stat(path)
		if err != nil {
			return nil, fmt.Errorf("ciigo: newFileAdoc: " + err.Error())
		}
	}

	adoc = &fileAdoc{
		path:     path,
		info:     fi,
		basePath: strings.TrimSuffix(path, ".adoc"),
	}

	return adoc, nil
}

//
// isHTMLLatest will return true if generated HTML is exist and its
// modification time is equal or greater than their asciidoc file; otherwise
// it will return false.
//
func (fa *fileAdoc) isHTMLLatest(htmlPath string) bool {
	htmlInfo, err := os.Stat(htmlPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		log.Printf("ciigo: os.Stat(%q): %s\n", htmlPath, err)
		return false
	}
	if htmlInfo == nil {
		return false
	}

	infoTime := fa.info.ModTime()
	htmlTime := htmlInfo.ModTime()

	return htmlTime.Equal(infoTime) || htmlTime.After(infoTime)
}

//
// toHTML convert the asciidoc file to HTML and store its metadata in
//
func (fa *fileAdoc) toHTML(htmlPath string, out io.Writer, force bool) {
	if fa.isHTMLLatest(htmlPath) && !force {
		return
	}

	var (
		ctx = context.Background()
		err error
	)

	fa.metadata, err = libasciidoc.ConvertFileToHTML(ctx, fa.path, out)
	if err != nil {
		log.Fatal(err)
	}
}
