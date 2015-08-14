// Log service
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

type AccessLog struct {
	UserID       string    `json:"user_id"`
	Method       string    `json:"method"`
	Host         string    `json:"host"`
	Path         string    `json:"path"`
	IP           string    `json:"ip"`
	ResponseTime float64   `json:"response_time"`
	CreatedAt    time.Time `json:"created_at"`
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

	logConsumer, err := consumer.ConsumePartition("log", 0, sarama.OffsetNewest)
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
			accessLog := &AccessLog{}
			json.Unmarshal(msg.Value, accessLog)
			// you can send log to analysis
			// example influxdb, newrelic, go-metric
			log.Printf("%s %s %s %d",
				accessLog.CreatedAt,
				accessLog.Method,
				accessLog.Path,
				accessLog.ResponseTime,
			)
		case <-signals:
			return
		}
	}
}
