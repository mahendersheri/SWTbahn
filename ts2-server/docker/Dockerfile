FROM golang:1.15-alpine as builder
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
RUN git clone https://gitlab.rz.uni-bamberg.de/swt/teaching/2020-ws/swt-pr1-2-m/bahnts2/ts2-server
RUN go get github.com/GeertJohan/go.rice/rice
ENV GO111MODULE=on
RUN cd ts2-server && GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/ts2-server
RUN rice append --exec /go/bin/ts2-server -i ./ts2-server/server

FROM alpine:latest
MAINTAINER Nicolas Piganeau <npi@m4x.org>
COPY demo.json /
COPY --from=builder /go/bin/ts2-server /ts2-server
CMD ["/ts2-server", "-loglevel", "debug", "/demo.json"]
