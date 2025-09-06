package main

import (
	"context"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "order-processors",
		Topic:    "orders",
		MinBytes: 1,
		MaxBytes: 10e6,
	})
	defer reader.Close()

	ctx := context.Background()

	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}

		log.Printf("Consumed from partition %d offset %d: key=%s value=%s",
			m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
