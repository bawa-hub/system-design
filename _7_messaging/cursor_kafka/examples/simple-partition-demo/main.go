package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	fmt.Println("🎓 Simple Partition Demo")
	fmt.Println(strings.Repeat("=", 30))
	fmt.Println("Let's see how messages go to different partitions!")
	fmt.Println()

	// Create producer
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer producer.Close()

	topic := "simple-demo"

	// Test 1: Send messages with different keys
	fmt.Println("📤 Sending messages with different keys...")
	keys := []string{"alice", "bob", "charlie", "david", "eve"}
	
	for _, key := range keys {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(key),
			Value: sarama.StringEncoder(fmt.Sprintf("Hello from %s!", key)),
		}

		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Printf("Error: %v", err)
		} else {
			fmt.Printf("✅ %s → Partition %d, Offset %d\n", key, partition, offset)
		}
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println()
	fmt.Println("📤 Sending messages with SAME key...")
	
	// Test 2: Send messages with same key
	sameKey := "alice"
	for i := 1; i <= 3; i++ {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(sameKey),
			Value: sarama.StringEncoder(fmt.Sprintf("Message %d from %s", i, sameKey)),
		}

		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Printf("Error: %v", err)
		} else {
			fmt.Printf("✅ %s (msg %d) → Partition %d, Offset %d\n", sameKey, i, partition, offset)
		}
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println()
	fmt.Println("🎯 Key Learning:")
	fmt.Println("• Different keys → Different partitions (parallel processing)")
	fmt.Println("• Same key → Same partition (keeps order)")
	fmt.Println()
	fmt.Println("💡 Next: Run a consumer to see the messages!")
}
