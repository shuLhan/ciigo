.PHONY: all generate build

all: generate build

generate:
	go generate

build:
	go install ./cmd/ciigo
