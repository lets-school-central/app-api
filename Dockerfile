# syntax=docker/dockerfile:1

FROM golang:1.20.6-alpine3.18 AS build-stage

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN mkdir -p /usr/local/bin
RUN CGO_ENABLED=0 go build -v -o /usr/local/bin/lets-school-central-backend main.go

FROM alpine:3.18 AS build-release-stage

ENV RUN_MIGRATIONS=1

COPY --from=build-stage /usr/local/bin/lets-school-central-backend /app

EXPOSE 8080

ENTRYPOINT /app serve --http=0.0.0.0:8080
