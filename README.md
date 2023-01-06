# todo-list-challenge

### How to run it?
First you have to have docker & docker compose, then run the rabbitmq container

```
docker compose up message-broker -d
```

then just run the program

```
go run main.go
```

if you need to run the receiver / worker for rabbitmq, then run

```
go run receiver/receiver.go
```
