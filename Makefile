.PHONY: build clean test

CWD :=$(shell pwd)
BINDIR := $(CWD)/bin

build: clean
	@cd $(pkg); \
	go build -o $(BINDIR)/$(pkg) .

test:
	@cd $(pkg); \
		go test -v

clean:
	@rm -rf bin/*

run: build
	bin/$(pkg)

