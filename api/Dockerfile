FROM golang:1.10
MAINTAINER neko-neko <seiichi.nishikata456@gmail.com>

RUN mkdir -p /go/src/github.com/neko-neko/goflippy/api
COPY . /go/src/github.com/neko-neko/goflippy/api/
WORKDIR /go/src/github.com/neko-neko/goflippy/api/
RUN make build && \
  mv release/linux/amd64/goflippy-api /bin/

CMD ["/bin/goflippy-api"]
