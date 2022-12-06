# syntax=docker/dockerfile:1

FROM golang:1.19-buster

# RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . . 

RUN go mod tidy

RUN go build -o todo-list-challenge

EXPOSE 8081

ENTRYPOINT [ "/app/todo-list-challenge" ]