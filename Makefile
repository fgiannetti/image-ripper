SHELL=/bin/sh
NAME=image-ripper

all: tests build
tests: clean unit-tests

build: $(wildcard *.go) $(wildcard */*.go)
	go build -v .

clean:
	rm -f $(NAME)
	go clean -testcache

unit-tests: $(wildcard *.go) $(wildcard */*.go)
	go test ./pkg/...