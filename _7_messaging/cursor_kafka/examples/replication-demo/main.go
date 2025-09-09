package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	fmt.Println("🔄 Replication and Fault Tolerance Demo")
	fmt.Println(strings.Repeat("=", 45))
	fmt.Println("Understanding how Kafka handles failures...")
	fmt.Println()

	// Create producer
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer producer.Close()

	topic := "replicated-demo"

	// Send some messages
	fmt.Println("📤 Sending messages to replicated topic...")
	messages := []string{
		"Important message 1",
		"Important message 2", 
		"Important message 3",
		"Important message 4",
		"Important message 5",
	}

	for i, msg := range messages {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(fmt.Sprintf("key-%d", i+1)),
			Value: sarama.StringEncoder(msg),
		}

		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Printf("Error: %v", err)
		} else {
			fmt.Printf("✅ %s → Partition %d, Offset %d\n", msg, partition, offset)
		}
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println()
	fmt.Println("🎯 Key Concepts:")
	fmt.Println("• Replication Factor = 1: No replication (single point of failure)")
	fmt.Println("• Leader: Broker 1 handles all read/write requests")
	fmt.Println("• ISR (In-Sync Replicas): [1] - only one replica")
	fmt.Println("• If Broker 1 fails, data is lost!")
	fmt.Println()
	fmt.Println("💡 In production, use Replication Factor = 3 for fault tolerance!")
}
