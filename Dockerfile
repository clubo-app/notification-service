FROM golang:1.18.1-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/cosmtrek/air@latest

COPY ./services/notification ./services/notification
COPY ./packages ./packages

EXPOSE 8080

WORKDIR /app/services/notification

RUN go build -o notification-service

ENTRYPOINT ./notification-service
