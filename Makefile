.PHONY: build build-api cross-build-api build-admin cross-build-admin clean test setup

PACKAGES = $(shell go list ./... | grep -v /vendor)
LD_FLAGS = -extldflags '-static'

## Build servers
build: build-api build-admin

## Build API server
build-api:
	$(eval OS := $(shell uname -s | tr '[A-Z]' '[a-z]'))
	$(eval ARCH := $(shell arch | sed -e 's/i386/386/' | sed -e 's/x86_64/amd64/'))
	cd api
	go build -ldflags "${LD_FLAGS}" -o release/$(OS)/$(ARCH)/goflippy-api -v github.com/neko-neko/goflippy/api

## Build cross platform build API server
cross-build-api:
	for os in darwin linux windows; do \
		for arch in amd64 386; do \
			GOOS=$$os GOARCH=$$arch go build -ldflags "${LD_FLAGS}" -o release/$$os/$$arch/goflippy-api -v github.com/neko-neko/goflippy/api; \
		done; \
	done

## Build Admin server
build-admin:
	$(eval OS := $(shell uname -s | tr '[A-Z]' '[a-z]'))
	$(eval ARCH := $(shell arch | sed -e 's/i386/386/' | sed -e 's/x86_64/amd64/'))
	cd admin
	go build -ldflags "${LD_FLAGS}" -o release/$(OS)/$(ARCH)/goflippy-admin -v github.com/neko-neko/goflippy/admin

## Build cross platform build Admin server
cross-build-admin:
	for os in darwin linux windows; do \
		for arch in amd64 386; do \
			GOOS=$$os GOARCH=$$arch go build -ldflags "${LD_FLAGS}" -o release/$$os/$$arch/goflippy-admin -v github.com/neko-neko/goflippy/admin; \
		done; \
	done

## Clean artifacts
clean:
	rm -rf api/release
	rm -rf admin/release

## Run test
test:
	@for f in "${PACKAGES}"; do \
		go test -race -cover -covermode=atomic $$f; \
	done

## Run test for CI
test-ci:
	@for f in "${PACKAGES}"; do \
		go test -race -coverprofile=profile.out -covermode=atomic $$f; \
		if [ -f profile.out ]; then \
			cat profile.out >> coverage.txt; \
			rm profile.out; \
		fi; \
	done

## Run go generate
generate:
	go generate github.com/neko-neko/goflippy/api/handler/v1

## Install development tools
setup:
	go get -v -u github.com/go-swagger/go-swagger/cmd/swagger
