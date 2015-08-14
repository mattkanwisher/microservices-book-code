package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
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

	err = ch.Qos(
		3,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	Panic(err)

	// register a consumer
	msgs, err := ch.Consume(
		"example_queue", // queue name
		"",              // consumer
		false,           // no-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	Panic(err)

	forever := make(chan bool)

	go func() {
		// receive message channel
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			//long runing task
			time.Sleep(time.Second * 10)
			// tell rabbitmq task is done,
			// delete this message from queue
			d.Ack(false)
		}
	}()

	log.Printf("Waiting for messages...")
	<-forever
}
