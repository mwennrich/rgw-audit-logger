FROM golang:1.19 AS build
ENV GO111MODULE=on
ENV CGO_ENABLED=0


COPY . /app
WORKDIR /app

RUN go build rgw-audit-logger.go
RUN strip rgw-audit-logger

FROM scratch

WORKDIR /
USER 0

COPY --from=build /app/rgw-audit-logger /rgw-audit-logger

ENTRYPOINT ["/rgw-audit-logger"]
