RELEASES= _bin/ciigo-linux-amd64 \
	_bin/ciigo-darwin-amd64

.PHONY: all lint install build build-release

all: install

lint:
	golangci-lint run --enable-all \
		--disable=wsl --disable=gomnd --disable=funlen ./...

install:
	go generate
	go install ./cmd/ciigo-example
	go install ./cmd/ciigo

build-release: _bin $(RELEASES)

_bin:
	mkdir -p _bin

_bin/ciigo-linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build -o $@ ./cmd/ciigo

_bin/ciigo-darwin-amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 \
		go build -o $@ ./cmd/ciigo
