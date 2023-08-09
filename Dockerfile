# syntax=docker/dockerfile:1

FROM golang:1.20.6-alpine3.18 AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN mkdir -p /usr/local/bin
RUN CGO_ENABLED=0 go build -v -o /usr/local/bin/lets-school-central-backend main.go

ADD https://github.com/benbjohnson/litestream/releases/download/v0.3.11/litestream-v0.3.11-linux-amd64.tar.gz /tmp/litestream.tar.gz
RUN tar -C /usr/local/bin -xzf /tmp/litestream.tar.gz

FROM alpine:3.18

ENV RUN_MIGRATIONS=1

COPY .docker/litestream.yml /etc/litestream.yml
COPY .docker/run.sh /scripts/run.sh

COPY --from=builder /usr/local/bin/lets-school-central-backend /usr/local/bin/lets-school-central-backend

RUN chmod +x /scripts/run.sh
RUN chmod +x /usr/local/bin/litestream
RUN chmod +x /usr/local/bin/lets-school-central-backend

RUN apk add bash

RUN mkdir -p /data

ENV PORT=8080
EXPOSE ${PORT}

CMD [ "/scripts/run.sh" ]
