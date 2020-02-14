.PHONY: all lint generate build

all: generate build

lint:
	golangci-lint run --enable-all \
		--disable=wsl --disable=gomnd --disable=funlen ./...

generate:
	go generate

build:
	go install ./cmd/ciigo-example
	go install ./cmd/ciigo
