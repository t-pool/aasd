# Build Gmit in a stock Go builder container
FROM golang:1.10-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /go-mit
RUN cd /go-mit && make gmit

# Pull Gmit into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-mit/build/bin/mit /usr/local/bin/

EXPOSE 8545 8546 9999 9999/udp
ENTRYPOINT ["gmit"]
