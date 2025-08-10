# Makefile

# Makefile for mkinput project

.PHONY: build clean test coverage

build:
	go build -o mkinput ./cmd/mkinput
	go build -o track ./cmd/track

clean:
	rm -f mkinput track

# Run tests and report coverage

test:
	go test ./...

coverage:
	go test -cover ./...
