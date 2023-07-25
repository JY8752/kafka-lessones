package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

const (
	topicName = "ticket-order"
	kafkaAddr = "localhost:29092"
)

func main() {
	fmt.Println("Consumer Started.")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaAddr},
		Topic:    topicName,
		MaxBytes: 10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Println(string(m.Value))
	}

	if err := r.Close(); err != nil {
		fmt.Println("failed to close reader:", err)
	}
}
