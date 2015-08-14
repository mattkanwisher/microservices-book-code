package main

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"time"
)

const kafkaAddr = "128.199.223.6:9092"

type Order struct {
	UserID    string    `json:"user_id"`
	OrderID   string    `json:"order"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	config := sarama.NewConfig()
	// Handle errors manually
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{kafkaAddr}, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	logConsumer, err := consumer.ConsumePartition("buy", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer logConsumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	for {
		select {
		case err := <-logConsumer.Errors():
			log.Println(err)
		case msg := <-logConsumer.Messages():
			order := &Order{}
			json.Unmarshal(msg.Value, order)
			log.Printf("notification to %s with order %s", order.UserID, order.OrderID)
		case <-signals:
			return
		}
	}
}
