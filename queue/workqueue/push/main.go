package main

import (
	"github.com/streadway/amqp"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// open connection to rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@192.168.99.100:5672/")
	Panic(err)
	defer conn.Close()

	// open channel
	ch, err := conn.Channel()
	Panic(err)
	defer ch.Close()

	// declare queue
	q, err := ch.QueueDeclare(
		"example_queue", // queue name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	Panic(err)

	err = ch.Publish(
		"",     // exchange
		q.Name, // queue name
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte("hello world"),
		})
	Panic(err)
}
