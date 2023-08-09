# syntax=docker/dockerfile:1

FROM golang:1.20.6-alpine3.18 AS builder

WORKDIR /usr/src/app

ADD https://github.com/benbjohnson/litestream/releases/download/v0.3.11/litestream-v0.3.11-linux-amd64.tar.gz /tmp/litestream.tar.gz
RUN tar -C /usr/local/bin -xzf /tmp/litestream.tar.gz

COPY go.mod go.sum ./

RUN \
    --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . ./

RUN mkdir -p /usr/local/bin
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' -v -o /usr/local/bin/lets-school-central-backend main.go

FROM alpine:3.18

ENV RUN_MIGRATIONS=1

COPY .docker/litestream.yml /etc/litestream.yml
COPY .docker/run.sh /scripts/run.sh

COPY --from=builder /usr/local/bin/litestream /usr/local/bin/litestream
COPY --from=builder /usr/local/bin/lets-school-central-backend /usr/local/bin/lets-school-central-backend

RUN chmod +x /scripts/run.sh
RUN chmod +x /usr/local/bin/litestream
RUN chmod +x /usr/local/bin/lets-school-central-backend

RUN apk add bash

RUN mkdir -p /data

ENV PORT=8080
EXPOSE ${PORT}

CMD [ "/scripts/run.sh" ]
