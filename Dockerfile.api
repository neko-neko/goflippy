FROM golang:1.11
EXPOSE 9001

WORKDIR /go/src/github.com/neko-neko/goflippy
COPY . .
RUN mkdir /app && \
      make build-api && \
      mv release/linux/amd64/goflippy-api /app/

CMD ["/app/goflippy-api"]
