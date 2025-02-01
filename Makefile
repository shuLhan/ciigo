## SPDX-FileCopyrightText: 2019 Shulhan <ms@kilabit.info>
## SPDX-License-Identifier: GPL-3.0-or-later

VERSION:=$(shell git describe --tags)
ENV_GOBIN:=$(shell go env GOBIN)
LDFLAGS:=-ldflags "-s -w -X 'git.sr.ht/~shulhan/ciigo.Version=$(VERSION)'"
DIR_BUILD:=_bin

.PHONY: all
all: lint build test

.PHONY: lint
lint:
	go run ./internal/cmd/gocheck ./...
	go vet ./...

.PHONY: test
test:
	find ./testdata -name "*.html" -delete
	CGO_ENABLED=1 go test -failfast -v -race -p=1 -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o cover.html

.PHONY: install
install: build
	mv _bin/ciigo $(ENV_GOBIN)

.PHONY: run-example
run-example:
	go run ./internal/cmd/ciigo-example

.PHONY: build
build:
	mkdir -p $(DIR_BUILD)
	CGO_ENABLED=0 go build $(LDFLAGS) -o $(DIR_BUILD) ./cmd/...

.PHONY: serve-doc
serve-doc:
	go run ./cmd/ciigo -address=127.0.0.1:20757 serve _doc

.PHONY: chroot-setup
chroot-setup:
	sudo mkdir -p /var/lib/machines/arch.test
	sudo pacstrap -c /var/lib/machines/arch.test base-devel systemd go
	sudo mkdir /var/lib/machines/arch.test/root/ciigo

.PHONY: chroot-test
chroot-test:
	sudo rsync -r . /var/lib/machines/arch.test/root/ciigo/
	sudo systemd-nspawn --bind=/tmp \
		--bind=$${HOME}/go/pkg:/root/go/pkg \
		--bind=$${PWD}/../share:/root/share \
		--bind=$${HOME}/.cache/go-build:/root/.cache/go-build \
		-D /var/lib/machines/arch.test \
		make -C /root/ciigo test
		#sh -c "rm -rf /root/test; mkdir /root/test; stat /root/test; sleep 1; touch /root/test/file; stat /root/test"
