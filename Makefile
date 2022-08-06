## SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
## SPDX-License-Identifier: GPL-3.0-or-later

.PHONY: all lint test install build

VERSION:=$(shell git describe --tags)
LDFLAGS:=-ldflags "-s -w -X 'git.sr.ht/~shulhan/ciigo.Version=$(VERSION)'"
DIR_BUILD:=_bin

all: test lint build

lint:
	-golangci-lint run ./...

test:
	go run ./cmd/ciigo-example embed
	CGO_ENABLED=1 go test -v -race ./...

install: build
	mv _bin/ciigo $(GOBIN)

run-example:
	DEBUG=1 go run ./cmd/ciigo-example

build:
	mkdir -p $(DIR_BUILD)
	CGO_ENABLED=0 go build $(LDFLAGS) -o $(DIR_BUILD) ./cmd/...
