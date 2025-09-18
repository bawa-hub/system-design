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
	// Kafka broker configuration
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// Create consumer group
	consumer, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "my-consumer-group", config)
	if err != nil {
		log.Fatalf("Error creating consumer group: %v", err)
	}
	defer consumer.Close()

	// Topic to consume from
	topic := "user-events"

	// Create consumer handler
	handler := &ConsumerGroupHandler{}

	// Setup signal handling
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	// Start consuming
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-sigchan:
				fmt.Println("🛑 Interrupt signal received, stopping consumer...")
				return
			default:
				ctx := context.Background()
				err := consumer.Consume(ctx, []string{topic}, handler)
				if err != nil {
					log.Printf("Error from consumer: %v", err)
					return
				}
			}
		}
	}()

	// Handle errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Consumer error: %v", err)
		}
	}()

	fmt.Println("🚀 Consumer started! Listening for messages...")
	fmt.Println("Press Ctrl+C to stop")
	wg.Wait()
}

// ConsumerGroupHandler implements sarama.ConsumerGroupHandler
type ConsumerGroupHandler struct{}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (h *ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	fmt.Println("🔧 Consumer group session started")
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (h *ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	fmt.Println("🧹 Consumer group session ended")
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			if message == nil {
				return nil
			}

			// Process the message
			fmt.Printf("📨 Received message:\n")
			fmt.Printf("   Topic: %s\n", message.Topic)
			fmt.Printf("   Partition: %d\n", message.Partition)
			fmt.Printf("   Offset: %d\n", message.Offset)
			fmt.Printf("   Key: %s\n", string(message.Key))
			fmt.Printf("   Value: %s\n", string(message.Value))
			fmt.Printf("   Timestamp: %s\n", message.Timestamp.Format("2006-01-02 15:04:05"))
			fmt.Printf("   %s\n", strings.Repeat("─", 50))

			// Mark message as processed
			session.MarkMessage(message, "")

		case <-session.Context().Done():
			return nil
		}
	}
}
