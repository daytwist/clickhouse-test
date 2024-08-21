FROM golang:1.13-alpine

ENV VERSION v4.7.1

RUN apk add make vim --no-cache git

RUN go get -v -d github.com/golang-migrate/migrate/cli \
  && go get -v -d github.com/lib/pq

WORKDIR /go/src/github.com/golang-migrate/migrate

RUN git checkout ${VERSION} \
  && go build -tags 'clickhouse' -o ./bin/migrate ./cli

WORKDIR /src
