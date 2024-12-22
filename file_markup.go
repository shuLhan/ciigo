// SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
// SPDX-License-Identifier: GPL-3.0-or-later

package ciigo

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// List of markup kind.
const (
	markupKindAdoc     = 1
	markupKindMarkdown = 2
)

// FileMarkup contains the markup path and its kind.
type FileMarkup struct {
	info os.FileInfo // info contains FileInfo of markup file.

	path     string // Full path to markup file.
	pathHTML string // path to HTML file.

	kind int
}

// NewFileMarkup create new FileMarkup instance form file in "filePath".
// The "fi" option is optional, if its nil it will Stat-ed manually.
func NewFileMarkup(filePath string, fi os.FileInfo) (fmarkup *FileMarkup, err error) {
	var logp = `NewFileMarkup`

	if len(filePath) == 0 {
		return nil, fmt.Errorf(`%s: empty path`, logp)
	}
	if fi == nil {
		fi, err = os.Stat(filePath)
		if err != nil {
			return nil, fmt.Errorf(`%s: %w`, logp, err)
		}
	}

	var ext = strings.ToLower(filepath.Ext(filePath))
	var basePath = strings.TrimSuffix(filePath, ext)

	fmarkup = &FileMarkup{
		info:     fi,
		path:     filePath,
		pathHTML: basePath + `.html`,
		kind:     markupKind(ext),
	}

	return fmarkup, nil
}

func markupKind(ext string) int {
	switch ext {
	case extAsciidoc:
		return markupKindAdoc
	case extMarkdown:
		return markupKindMarkdown
	}
	return 0
}
