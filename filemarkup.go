// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type fileMarkup struct {
	info os.FileInfo // info contains FileInfo of markup file.

	basePath string // basePath contains full path to file without markup extension.
	path     string // path contains full path to markup file.
	pathHtml string // path to HTML file.
}

func newFileMarkup(filePath string, fi os.FileInfo) (fmarkup *fileMarkup, err error) {
	var (
		logp = `newFileMarkup`

		ext string
	)

	if len(filePath) == 0 {
		return nil, fmt.Errorf(`%s: empty path`, logp)
	}
	if fi == nil {
		fi, err = os.Stat(filePath)
		if err != nil {
			return nil, fmt.Errorf(`%s: %w`, logp, err)
		}
	}

	ext = strings.ToLower(filepath.Ext(filePath))

	fmarkup = &fileMarkup{
		path:     filePath,
		info:     fi,
		basePath: strings.TrimSuffix(filePath, ext),
	}

	fmarkup.pathHtml = fmarkup.basePath + `.html`

	return fmarkup, nil
}

// isNewerThanHtml return true if the markup file is newer than HTML file.
func (fm *fileMarkup) isNewerThanHtml() bool {
	var (
		fi os.FileInfo
	)
	fi, _ = os.Stat(fm.pathHtml)
	if fi == nil {
		return true
	}
	return fm.info.ModTime().After(fi.ModTime())
}
