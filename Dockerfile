# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

RUN apk update

WORKDIR /app

COPY . . 

RUN go mod tidy

RUN go build -o todo-list-challenge

EXPOSE 3030

ENTRYPOINT [ "/app/todo-list-challenge" ]