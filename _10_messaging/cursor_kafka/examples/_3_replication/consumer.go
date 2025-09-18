package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"

	"github.com/IBM/sarama"
)

func main() {
	fmt.Println("🎧 Replication Consumer Demo")
	fmt.Println(strings.Repeat("=", 30))
	fmt.Println("Reading messages from replicated topic...")
	fmt.Println()

	// Create consumer group
	consumer, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "replication-group", nil)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	topic := "replicated-demo"
	handler := &ReplicationHandler{}

	// Setup signal handling
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-sigchan:
				fmt.Println("🛑 Stopping consumer...")
				return
			default:
				ctx := context.Background()
				err := consumer.Consume(ctx, []string{topic}, handler)
				if err != nil {
					log.Printf("Error: %v", err)
					return
				}
			}
		}
	}()

	go func() {
		for err := range consumer.Errors() {
			log.Printf("Consumer error: %v", err)
		}
	}()

	fmt.Println("🎧 Consumer started! Press Ctrl+C to stop")
	wg.Wait()
}

type ReplicationHandler struct{}

func (h *ReplicationHandler) Setup(sarama.ConsumerGroupSession) error {
	fmt.Println("🔧 Consumer group started")
	return nil
}

func (h *ReplicationHandler) Cleanup(sarama.ConsumerGroupSession) error {
	fmt.Println("🧹 Consumer group ended")
	return nil
}

func (h *ReplicationHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			if message == nil {
				return nil
			}

			fmt.Printf("📨 Partition %d, Offset %d: %s\n", 
				message.Partition, message.Offset, string(message.Value))

			session.MarkMessage(message, "")

		case <-session.Context().Done():
			return nil
		}
	}
}
