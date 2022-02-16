RELEASES:= \
	_bin/ciigo-linux-amd64 \
	_bin/ciigo-darwin-amd64

.PHONY: all lint test install serve build build-release
.FORCE:

all: test lint

lint:
	golangci-lint run ./...

test:
	go run ./internal/cmd/goembed
	CGO_ENABLED=1 go test -v -race ./...

install:
	go install ./cmd/ciigo

run-example:
	DEBUG=1 go run ./cmd/ciigo-example

build-release: $(RELEASES)

_bin/ciigo-linux-amd64: .FORCE
_bin/ciigo-linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build -o $@ ./cmd/ciigo

_bin/ciigo-darwin-amd64: .FORCE
_bin/ciigo-darwin-amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 \
		go build -o $@ ./cmd/ciigo
