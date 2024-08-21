FROM golang:1.23-alpine

ENV VERSION v4.17.1

RUN apk add --no-cache make vim git

WORKDIR /go/src/github.com/golang-migrate/migrate

RUN git clone https://github.com/golang-migrate/migrate.git . \
  && git checkout ${VERSION} \
  && go mod download

RUN go build -tags 'clickhouse' -o ./bin/migrate ./cli

WORKDIR /src
