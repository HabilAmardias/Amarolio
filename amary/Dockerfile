FROM golang:1.25 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
ENV GOPROXY=https://goproxy.io,direct
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o amarolio-amary main.go

FROM alpine:latest AS release-stage

WORKDIR /

COPY --from=build-stage /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build-stage /app/amarolio-amary amarolio-amary
COPY --from=build-stage /app/db ./db

EXPOSE 3000

ENTRYPOINT [ "/bin/sh", "-c", "/amarolio-amary"]