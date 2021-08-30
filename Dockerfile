FROM golang:1.17 AS build
ENV GO111MODULE=on
ENV CGO_ENABLED=0


COPY . /app
WORKDIR /app

RUN go build rgw-audit-logger.go
RUN strip rgw-audit-logger

FROM alpine:3.14

WORKDIR /
USER 0

COPY --from=build /app .

ENTRYPOINT ./rgw-audit-logger
