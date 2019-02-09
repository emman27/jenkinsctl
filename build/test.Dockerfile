FROM golang:1.11-alpine

WORKDIR /go/src/github.com/emman27/jenkinsctl

RUN apk add alpine-sdk dep

COPY Gopkg.lock Gopkg.toml ./
RUN dep ensure --vendor-only

COPY . .

RUN go test ./...
