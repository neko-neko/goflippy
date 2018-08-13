FROM golang:1.10
MAINTAINER neko-neko <seiichi.nishikata456@gmail.com>

RUN mkdir -p /go/src/github.com/neko-neko/goflippy
COPY . /go/src/github.com/neko-neko/goflippy/
WORKDIR /go/src/github.com/neko-neko/goflippy/
RUN make build-admin && \
  mv release/linux/amd64/goflippy-admin /bin/

CMD ["/bin/goflippy-admin"]
