.PHONY: build clean

build: clean
	cd $(pkg); \
	go build -o ../bin/$(pkg) .

clean:
	rm -rf bin/*

