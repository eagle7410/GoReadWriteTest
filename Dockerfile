# Do your stuff, build a static binary
FROM golang:alpine as builder
WORKDIR /go/src/app
COPY . .
