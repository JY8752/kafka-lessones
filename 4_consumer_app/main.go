package main

import (
	"fmt"
	"github.com/segmentio/kafka-go"
)

func main() {
	fmt.Println("Consumer Started.")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		GroupID:  "consumer-group-id",
		Topic:    "topic-A",
		MaxBytes: 10e6, // 10MB
	})
}
