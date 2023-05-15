# syntax=docker/dockerfile:1

## Build
FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /docker-gs-ping

## Deploy
FROM debian:buster-slim

WORKDIR /

COPY --from=build /docker-gs-ping /docker-gs-ping

EXPOSE 8080


ENTRYPOINT ["/docker-gs-ping"]