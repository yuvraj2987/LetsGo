.PHONY: build clean test

build: clean
	@cd $(pkg); \
	go build -o ../../bin/$(pkg) .

test:
	@cd $(pkg); \
		go test -v

clean:
	@rm -rf bin/*

run: build
	bin/$(pkg)

