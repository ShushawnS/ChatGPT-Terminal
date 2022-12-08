FROM golang:1.19.4-alpine as dev 

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

WORKDIR /go/src/github.com/shushawns/ChatGPT-Terminal


