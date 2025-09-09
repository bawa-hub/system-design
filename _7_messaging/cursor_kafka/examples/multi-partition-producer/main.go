package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	// Kafka broker configuration
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	// Create producer
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer producer.Close()

	// Topic name
	topic := "user-events-multi"

	// Different user types to demonstrate partitioning
	userTypes := []string{"premium", "standard", "basic"}
	actions := []string{"login", "purchase", "view", "logout", "update_profile"}

	fmt.Println("🚀 Starting multi-partition producer...")
	fmt.Println("📊 Sending messages with different keys to demonstrate partitioning")

	// Send messages with different keys to different partitions
	for i := 1; i <= 20; i++ {
		userType := userTypes[rand.Intn(len(userTypes))]
		action := actions[rand.Intn(len(actions))]
		userID := fmt.Sprintf("%s-user-%d", userType, i)
		
		message := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(userID),
			Value: sarama.StringEncoder(fmt.Sprintf("%s %s at %s", userID, action, time.Now().Format("15:04:05"))),
		}

		// Send message
		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Printf("Failed to send message %d: %v", i, err)
		} else {
			fmt.Printf("✅ %s → Partition %d, Offset %d\n", userID, partition, offset)
		}

		// Small delay to see messages flowing
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Println("🎉 All messages sent successfully!")
	fmt.Println("💡 Notice how different keys go to different partitions!")
}
