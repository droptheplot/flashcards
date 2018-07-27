FROM golang:1.10-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git

ADD . /go/src/github.com/droptheplot/flashcards
WORKDIR /go/src/github.com/droptheplot/flashcards

RUN go get -v ./...
RUN go get github.com/pilu/fresh
