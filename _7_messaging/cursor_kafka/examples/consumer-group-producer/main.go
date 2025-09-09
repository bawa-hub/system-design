package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	fmt.Println("📤 Consumer Group Producer Demo")
	fmt.Println(strings.Repeat("=", 35))
	fmt.Println("Sending messages to test consumer groups...")
	fmt.Println()

	// Create producer
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer producer.Close()

	topic := "simple-demo"

	// Send messages with different keys to different partitions
	keys := []string{"user1", "user2", "user3", "user4", "user5", "user6", "user7", "user8", "user9"}
	
	for i, key := range keys {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(key),
			Value: sarama.StringEncoder(fmt.Sprintf("Message %d from %s", i+1, key)),
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
	fmt.Println("🎉 All messages sent!")
	fmt.Println("💡 Now run multiple consumers to see load balancing!")
}
