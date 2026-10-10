FROM golang:1.26-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -o amarolio-auth main.go

FROM alpine:latest AS release-stage

WORKDIR /

RUN apk add --no-cache ca-certificates tzdata

COPY --from=build-stage /app/amarolio-auth amarolio-auth
COPY --from=build-stage /app/db ./db

EXPOSE 3000

ENTRYPOINT [ "/bin/sh", "-c", "/amarolio-auth"]
