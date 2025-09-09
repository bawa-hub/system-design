package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	// Get consumer name from command line
	consumerName := "consumer-1"
	if len(os.Args) > 1 {
		consumerName = os.Args[1]
	}

	fmt.Printf("🎧 Consumer Group Demo - %s\n", consumerName)
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("Let's see how consumer groups share work!")
	fmt.Println()

	// Create consumer group
	consumer, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "demo-group", nil)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	topic := "simple-demo"
	handler := &ConsumerGroupHandler{consumerName: consumerName}

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
				fmt.Printf("🛑 %s stopping...\n", consumerName)
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

	fmt.Printf("🎧 %s started! Press Ctrl+C to stop\n", consumerName)
	wg.Wait()
}

type ConsumerGroupHandler struct {
	consumerName string
}

func (h *ConsumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	fmt.Printf("🔧 %s: Consumer group session started\n", h.consumerName)
	return nil
}

func (h *ConsumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	fmt.Printf("🧹 %s: Consumer group session ended\n", h.consumerName)
	return nil
}

func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	fmt.Printf("📋 %s: Assigned to partition %d\n", h.consumerName, claim.Partition())
	
	for {
		select {
		case message := <-claim.Messages():
			if message == nil {
				return nil
			}

			fmt.Printf("📨 %s: Partition %d, Offset %d: %s\n", 
				h.consumerName, message.Partition, message.Offset, string(message.Value))

			// Simulate some processing time
			time.Sleep(1 * time.Second)

			session.MarkMessage(message, "")

		case <-session.Context().Done():
			return nil
		}
	}
}
