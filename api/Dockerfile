FROM golang:1.10
EXPOSE 9000

ENV CGO_ENABLED=0
ENV GODEBUG=netdns=go

RUN mkdir -p /go/src/github.com/neko-neko/goflippy/api
COPY . /go/src/github.com/neko-neko/goflippy/api/
WORKDIR /go/src/github.com/neko-neko/goflippy/api/
RUN make build && \
  mv release/linux/amd64/goflippy-api /bin/

CMD ["/bin/goflippy-api"]
