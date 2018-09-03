[![CircleCI](https://circleci.com/gh/neko-neko/goflippy/tree/master.svg?style=shield&circle-token=b4e94e627c67fd9ab598b4b5124e98a65fb816ea)](https://circleci.com/gh/neko-neko/goflippy/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/neko-neko/goflippy)](https://goreportcard.com/report/github.com/neko-neko/goflippy)
[![codecov](https://codecov.io/gh/neko-neko/goflippy/branch/master/graph/badge.svg)](https://codecov.io/gh/neko-neko/goflippy)
[![Docker Automated build](https://img.shields.io/docker/automated/nekoneko/goflippy-api.svg)](https://hub.docker.com/r/nekoneko/goflippy-api/)
[![Docker Automated build](https://img.shields.io/docker/automated/nekoneko/goflippy-admin.svg)](https://hub.docker.com/r/nekoneko/goflippy-admin/)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

# TODO

- [ ] Management console
- [ ] Test

# Overview

goflippy is a platform of managing `Feature Toggles`.

# Using goflippy

TODO

# Client SDK

- Java [goflippy-java](https://github.com/neko-neko/goflippy-java)

# Setting Up goflippy For Development
## Run servers

goflippy support for `docker-compose`.
Run this:

```bash
$ docker-compose up
```

## Run tests

Run this:

```bash
$ docker-compose run --rm api make test
```

## Generate Swagger Document

goflippy is support swagger generator.  
Automatically generate swagger.json from documents in code.  
Run this:

```bash
$ make generate
```
