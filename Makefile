## SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
## SPDX-License-Identifier: GPL-3.0-or-later

.PHONY: all lint test install build serve-doc

VERSION:=$(shell git describe --tags)
LDFLAGS:=-ldflags "-s -w -X 'git.sr.ht/~shulhan/ciigo.Version=$(VERSION)'"
DIR_BUILD:=_bin

all: test lint build

lint:
	-golangci-lint run ./...

test:
	go run ./cmd/ciigo-example embed
	CGO_ENABLED=1 go test -failfast -v -race -p=1 ./...

install: build
	mv _bin/ciigo $(GOBIN)

run-example:
	DEBUG=1 go run ./cmd/ciigo-example

build:
	mkdir -p $(DIR_BUILD)
	CGO_ENABLED=0 go build $(LDFLAGS) -o $(DIR_BUILD) ./cmd/...

serve-doc:
	go run ./cmd/ciigo serve _doc

.PHONY: chroot-setup chroot-test

chroot-setup:
	sudo mkdir -p /var/lib/machines/arch.test
	sudo pacstrap -c /var/lib/machines/arch.test base-devel systemd go
	sudo mkdir /var/lib/machines/arch.test/root/ciigo

chroot-test:
	sudo rsync -r . /var/lib/machines/arch.test/root/ciigo/
	sudo systemd-nspawn --bind=/tmp \
		--bind=$${HOME}/go/pkg:/root/go/pkg \
		--bind=$${PWD}/../share:/root/share \
		--bind=$${HOME}/.cache/go-build:/root/.cache/go-build \
		-D /var/lib/machines/arch.test \
		make -C /root/ciigo test
		#sh -c "rm -rf /root/test; mkdir /root/test; stat /root/test; sleep 1; touch /root/test/file; stat /root/test"
