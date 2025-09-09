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

// pass argument to make new group name
// go run examples/multiple-consumers/main.go group-1
// go run examples/multiple-consumers/main.go group-2

// What to observe:
// Each consumer group gets ALL messages (broadcast pattern)
// Messages are duplicated across groups
// This is how you can have multiple applications processing the same data

func main() {
	// Get consumer group name from command line or use default
	consumerGroup := "my-consumer-group"
	if len(os.Args) > 1 {
		consumerGroup = os.Args[1]
	}

	fmt.Printf("🚀 Starting consumer in group: %s\n", consumerGroup)

	// Kafka broker configuration
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// Create consumer group
	consumer, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, consumerGroup, config)
	if err != nil {
		log.Fatalf("Error creating consumer group: %v", err)
	}
	defer consumer.Close()

	// Topic to consume from
	topic := "user-events-multi"

	// Create consumer handler
	handler := &ConsumerGroupHandler{consumerGroup: consumerGroup}

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
				fmt.Printf("🛑 Consumer %s stopping...\n", consumerGroup)
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

	fmt.Printf("🎧 Consumer %s started! Listening for messages...\n", consumerGroup)
	fmt.Println("Press Ctrl+C to stop")
	wg.Wait()
}

// ConsumerGroupHandler implements sarama.ConsumerGroupHandler
type ConsumerGroupHandler struct {
	consumerGroup string
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (h *ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	fmt.Printf("🔧 Consumer group %s session started\n", h.consumerGroup)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (h *ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	fmt.Printf("🧹 Consumer group %s session ended\n", h.consumerGroup)
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
			fmt.Printf("📨 [%s] Received message:\n", h.consumerGroup)
			fmt.Printf("   Topic: %s\n", message.Topic)
			fmt.Printf("   Partition: %d\n", message.Partition)
			fmt.Printf("   Offset: %d\n", message.Offset)
			fmt.Printf("   Key: %s\n", string(message.Key))
			fmt.Printf("   Value: %s\n", string(message.Value))
			fmt.Printf("   Timestamp: %s\n", message.Timestamp.Format("15:04:05"))
			fmt.Printf("   %s\n", strings.Repeat("─", 50))

			// Mark message as processed
			session.MarkMessage(message, "")

		case <-session.Context().Done():
			return nil
		}
	}
}
