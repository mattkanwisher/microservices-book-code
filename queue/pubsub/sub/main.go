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

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	Panic(err)

	err = ch.QueueBind(
		q.Name,              // queue name
		"",                  // routing key
		"uploader_received", // exchange
		false,
		nil)
	Panic(err)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // no-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	Panic(err)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Println("Triggered event uploader_received")
			log.Printf("Resizing %s", d.Body)
		}
	}()

	log.Printf("Waiting for event...")
	<-forever
}
