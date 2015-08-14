package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.99.100:5672/")
	Panic(err)
	defer conn.Close()

	ch, err := conn.Channel()
	Panic(err)
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"uploader_received", // name
		"fanout",            // type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	Panic(err)

	err = ch.Publish(
		"uploader_received", // exchange
		"",                  // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("some-image.jpg"),
		})
	Panic(err)

	log.Printf("Sent")
}
