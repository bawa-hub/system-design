package main

import (
	"fmt"
	"log"
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
	topic := "user-events"

	// Send messages
	for i := 1; i <= 10; i++ {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(fmt.Sprintf("user-%d", i)),
			Value: sarama.StringEncoder(fmt.Sprintf("User %d performed an action at %s", i, time.Now().Format(time.RFC3339))),
		}

		// Send message
		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Printf("Failed to send message %d: %v", i, err)
		} else {
			fmt.Printf("✅ Message %d sent successfully to partition %d at offset %d\n", i, partition, offset)
		}

		// Small delay to see messages flowing
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("🎉 All messages sent successfully!")
}
