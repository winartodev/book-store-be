FROM golang:1.17-alpine AS builder

RUN apk update && apk upgrade && \
    apk --update add git make curl

RUN mkdir /book-store-be

ADD . /book-store-be

WORKDIR /book-store-be

COPY env.sample .env

RUN go mod download

RUN go build app/main.go

EXPOSE 8080

ENTRYPOINT ["/book-store-be/main"]