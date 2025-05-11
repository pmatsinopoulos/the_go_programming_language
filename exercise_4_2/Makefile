.PHONY: build clean test

test:
	go test -v ./...

build:
	mkdir -p build
	go build -o ./build/the_go_programming_language

clean:
	rm -f -R build/
