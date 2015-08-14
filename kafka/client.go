// Buy service
package main

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
)

const kafkaAddr = "128.199.223.6:9092"

var SyncProducer sarama.SyncProducer
var AsyncProducer sarama.AsyncProducer

type AccessLog struct {
	UserID       string    `json:"user_id"`
	Method       string    `json:"method"`
	Host         string    `json:"host"`
	Path         string    `json:"path"`
	IP           string    `json:"ip"`
	ResponseTime float64   `json:"response_time"`
	CreatedAt    time.Time `json:"created_at"`
}

type Order struct {
	UserID    string    `json:"user_id"`
	OrderID   string    `json:"order"`
	CreatedAt time.Time `json:"created_at"`
}

func getUserID(r *http.Request) string {
	r.ParseForm()
	return r.Form.Get("user_id")
}

func BuyHandler(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)

	order := &Order{
		UserID:    userID,
		OrderID:   uuid.NewV1().String(),
		CreatedAt: time.Now().UTC(),
	}

	orderJson, _ := json.Marshal(order)

	pmsg := &sarama.ProducerMessage{
		Partition: 0,
		Topic:     "buy",
		Value:     sarama.ByteEncoder(orderJson),
	}

	if _, _, err := SyncProducer.SendMessage(pmsg); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(orderJson)
}

func LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		started := time.Now()
		userID := getUserID(r)
		next(w, r)

		accessLog := &AccessLog{
			UserID:       userID,
			Method:       r.Method,
			Host:         r.Host,
			Path:         r.RequestURI,
			IP:           r.RemoteAddr,
			ResponseTime: float64(time.Since(started)) / float64(time.Second),
		}

		logJson, _ := json.Marshal(accessLog)

		AsyncProducer.Input() <- &sarama.ProducerMessage{
			Partition: 0,
			Topic:     "log",
			Value:     sarama.ByteEncoder(logJson),
		}
	})
}

func main() {
	var err error
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewManualPartitioner
	SyncProducer, err = sarama.NewSyncProducer([]string{kafkaAddr}, config)
	if err != nil {
		panic(err)
	}
	defer SyncProducer.Close()

	AsyncProducer, err = sarama.NewAsyncProducer([]string{kafkaAddr}, config)
	if err != nil {
		panic(err)
	}
	defer AsyncProducer.Close()

	http.HandleFunc("/buy", LogMiddleware(BuyHandler))

	http.ListenAndServe(":3000", nil)
}
