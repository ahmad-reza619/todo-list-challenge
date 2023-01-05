package utilities

import (
    "context"
    "log"
    "time"

    amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
    if err != nil {
        log.Panicf("%s: %s", msg, err)
    }
}

func dialMq() *amqp.Connection {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    failOnError(err, "Failed to connect to rabbit mq")
    return conn;
}

func openChannel(conn *amqp.Connection) *amqp.Channel {
    ch, err := conn.Channel()
    failOnError(err, "Failed to open channel")
    return ch;
}

func createQueue(name string, ch *amqp.Channel) amqp.Queue {
    q, err := ch.QueueDeclare(
        name,
        false,
        false,
        false,
        false,
        nil,
    )
    failOnError(err, "Failed to declare a queue")
    return q;
}

func PublishMessageToMqQueue(name string, body amqp.Publishing) {
    conn := dialMq()
    defer conn.Close()
    ch := openChannel(conn)
    defer ch.Close()
    q := createQueue(name, ch)
    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()
    err := ch.PublishWithContext(
        ctx,
        "",
        q.Name,
        false,
        false,
        body,
    );
    failOnError(err, "Failed to publish a message")
    log.Printf("[x] Sent Message to rabbitmq")
}
