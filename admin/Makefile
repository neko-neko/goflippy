.PHONY: build cross-build clean test setup

VERSION      = $(shell git describe --tags --abbrev=0)
HASH         = $(shell git rev-parse --verify HEAD)
BUILD_DATE   = $(shell date '+%Y/%m/%d %H:%M:%S %Z')
PACKAGES     = $(shell go list ./... | grep -v /vendor)
PACKAGE_DIRS = $(shell go list -f {{.Dir}} ./... | grep -v /vendor | grep -v goflippy$$)
LD_FLAGS     = -extldflags '-static'

## Build Admin server
build:
	$(eval OS := $(shell uname -s | tr '[A-Z]' '[a-z]'))
	$(eval ARCH := $(shell arch | sed -e 's/i386/386/' | sed -e 's/x86_64/amd64/'))
	go build -ldflags "${LD_FLAGS}" -o release/$(OS)/$(ARCH)/goflippy-admin -v github.com/neko-neko/goflippy/admin

## Build cross platform build Admin server
cross-build:
	for os in darwin linux windows; do \
		for arch in amd64 386; do \
			GOOS=$$os GOARCH=$$arch go build -ldflags "${LD_FLAGS}" -o release/$$os/$$arch/goflippy-admin -v github.com/neko-neko/goflippy/admin; \
		done; \
	done

## Clean artifacts
clean:
	rm -rf release

## Run test
test:
	go test -cover $(PACKAGES)

## Run go generate
generate:
	go generate github.com/neko-neko/goflippy/admin/handler

## Install development tools
setup:
	go get -v -u github.com/go-swagger/go-swagger/cmd/swagger
