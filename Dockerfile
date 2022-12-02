# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o /todo-list-challenge

EXPOSE 8081

CMD [ "/todo-list-challenge" ]