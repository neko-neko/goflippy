.PHONY: build build-api build-admin cross-build cross-build-api cross-build-admin build-docker-api build-docker-admin clean test setup

VERSION      = $(shell git describe --tags --abbrev=0)
HASH         = $(shell git rev-parse --verify HEAD)
BUILD_DATE   = $(shell date '+%Y/%m/%d %H:%M:%S %Z')
PACKAGES     = $(shell go list ./... | grep -v /vendor)
PACKAGE_DIRS = $(shell go list -f {{.Dir}} ./... | grep -v /vendor | grep -v goflippy$$)
LD_FLAGS     = -extldflags '-static'
BUILD_FLAGS  = -X github.com/neko-neko/goflippy.Version=${VERSION} -X github.com/neko-neko/goflippy.Hash=${HASH}

## Build API server and Admin server
build: clean build-api

## Build API server
build-api:
	$(eval OS := $(shell uname -s | tr '[A-Z]' '[a-z]'))
	$(eval ARCH := $(shell arch | sed -e 's/i386/386/' | sed -e 's/x86_64/amd64/'))
	go build -ldflags "${LD_FLAGS} ${BUILD_FLAGS}" -o release/$(OS)/$(ARCH)/goflippy-api github.com/neko-neko/goflippy/cmd/goflippy-api

## Build Admin server
build-admin:
	$(eval OS := $(shell uname -s | tr '[A-Z]' '[a-z]'))
	$(eval ARCH := $(shell arch | sed -e 's/i386/386/' | sed -e 's/x86_64/amd64/'))
	go build -ldflags "${LD_FLAGS} ${BUILD_FLAGS}" -o release/$(OS)/$(ARCH)/goflippy-admin github.com/neko-neko/goflippy/cmd/goflippy-admin

## Build cross platform build API server and Admin server
cross-build: clean cross-build-api cross-build-admin

## Build cross platform build API server
cross-build-api:
	for os in darwin linux windows; do \
		for arch in amd64 386; do \
			GOOS=$$os GOARCH=$$arch go build -ldflags "${LD_FLAGS} ${BUILD_FLAGS}" -o release/$$os/$$arch/goflippy-api github.com/neko-neko/goflippy/cmd/goflippy-api; \
		done; \
	done

## Build cross platform build Admin server
cross-build-admin:
	for os in darwin linux windows; do \
		for arch in amd64 386; do \
			GOOS=$$os GOARCH=$$arch go build -ldflags "${LD_FLAGS} ${BUILD_FLAGS}" -o release/$$os/$$arch/goflippy-admin github.com/neko-neko/goflippy/cmd/goflippy-admin; \
		done; \
	done

build-docker-api:
	docker build -t goflippy-api:$(VERSION) -f cmd/goflippy-api/Dockerfile .

build-docker-admin:
	docker build -t goflippy-admin:$(VERSION) -f cmd/goflippy-admin/Dockerfile .

## Clean artifacts
clean:
	rm -rf release

## Run test
test:
	go test -cover $(PACKAGES)

## Run go generate
generate:
	go generate github.com/neko-neko/goflippy/cmd/goflippy-api/handler/v1

## Install development tools
setup:
	go get -v -u github.com/golang/dep/cmd/dep
	go get -v -u github.com/go-swagger/go-swagger/cmd/swagger
