package main

import (
	"context"
	"fmt"
	"log"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "orders",
		Balancer: &kafka.Hash{}, // ensures same key → same partition
	})
	defer writer.Close()

	ctx := context.Background()

	for i := 1; i <= 5; i++ {
		// key := fmt.Sprintf("user-%d", i%3)
		key := "user-10"
		value := fmt.Sprintf("Order #%d from %s", i, key)

		err := writer.WriteMessages(ctx, kafka.Message{
			Key:   []byte(key),
			Value: []byte(value),
		})
		if err != nil {
			log.Fatal("failed to write:", err)
		}

		log.Printf("Produced: %s", value)
		time.Sleep(500 * time.Millisecond)
	}
}
