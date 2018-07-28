FROM golang:1.10-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git

ADD . /go/src/github.com/droptheplot/flashcards
WORKDIR /go/src/github.com/droptheplot/flashcards

RUN go get -v ./...
RUN go get github.com/pilu/fresh

RUN go get github.com/golang-migrate/migrate
RUN cd /go/src/github.com/golang-migrate/migrate/cli && go get
RUN go build -tags 'postgres' -o /usr/local/bin/migrate github.com/golang-migrate/migrate/cli
