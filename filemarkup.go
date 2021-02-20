// Copyright 2019, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ciigo

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

type fileMarkup struct {
	path     string                 // path contains full path to markup file.
	info     os.FileInfo            // info contains FileInfo of markup file.
	basePath string                 // basePath contains full path to file without markup extension.
	metadata map[string]interface{} // metadata contains markup metadata.
}

func newFileMarkup(filePath string, fi os.FileInfo) (fmarkup *fileMarkup, err error) {
	if len(filePath) == 0 {
		return nil, fmt.Errorf("ciigo: newFileMarkup: empty path")
	}
	if fi == nil {
		fi, err = os.Stat(filePath)
		if err != nil {
			return nil, fmt.Errorf("newFileMarkup: " + err.Error())
		}
	}

	ext := strings.ToLower(path.Ext(filePath))

	fmarkup = &fileMarkup{
		path: filePath,
		info: fi,
	}

	fmarkup.basePath = strings.TrimSuffix(filePath, ext)

	return fmarkup, nil
}

//
// isHTMLLatest will return true if generated HTML is exist and its
// modification time is equal or greater than their markup file; otherwise
// it will return false.
//
func (fa *fileMarkup) isHTMLLatest(htmlPath string) bool {
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
