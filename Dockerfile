# BUILD

FROM golang:1.19-alpine AS build-env

LABEL maintainer="ubaidillah<ubed.dev@gmail.com>"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 8083

RUN go build -o /binary

# OPTIMIZE STAGE

FROM alpine:latest

WORKDIR /

COPY --from=build-env /binary /binary

EXPOSE 8083

ENTRYPOINT ["/binary"]