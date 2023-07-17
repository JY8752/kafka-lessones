package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

const (
	topicName = "ticket-order"
	kafkaAddr = "localhost:29092"
)

func main() {
	fmt.Println("Producer Started.")

	orderId := uuid.New().String()
	userId := "123"
	contentId := "55555"

	eventValue := fmt.Sprintf("order_id=%s, user_id=%s, content_id=%s", orderId, userId, contentId)

	w := &kafka.Writer{
		Addr:     kafka.TCP(kafkaAddr),
		Topic:    topicName,
		Balancer: &kafka.LeastBytes{},
	}

	defer w.Close()

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte(eventValue),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	fmt.Println("Producer Finished.")
}
