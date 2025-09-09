package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	fmt.Println("🎓 Lesson 2: Topics and Partitions Deep Dive")
	fmt.Println(strings.Repeat("=", 50))

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
	topic := "lesson2-demo"

	fmt.Println("📚 Experiment 1: Understanding Partitioning Strategy")
	fmt.Println("Sending messages with different keys to see partition distribution...")
	fmt.Println()

	// Test 1: Different keys (should go to different partitions)
	fmt.Println("🔍 Test 1: Different keys")
	keys := []string{"user-1", "user-2", "user-3", "user-4", "user-5"}
	for i, key := range keys {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(key),
			Value: sarama.StringEncoder(fmt.Sprintf("Message %d with key %s", i+1, key)),
		}

		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		} else {
			fmt.Printf("✅ Key: %s → Partition: %d, Offset: %d\n", key, partition, offset)
		}
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println()
	fmt.Println("🔍 Test 2: Same keys (should go to same partition)")
	
	// Test 2: Same keys (should go to same partition)
	sameKey := "user-same"
	for i := 1; i <= 3; i++ {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(sameKey),
			Value: sarama.StringEncoder(fmt.Sprintf("Message %d with same key %s", i, sameKey)),
		}

		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		} else {
			fmt.Printf("✅ Key: %s → Partition: %d, Offset: %d\n", sameKey, partition, offset)
		}
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println()
	fmt.Println("🔍 Test 3: No key (round-robin distribution)")
	
	// Test 3: No key (round-robin)
	for i := 1; i <= 6; i++ {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(fmt.Sprintf("Message %d with no key", i)),
		}

		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		} else {
			fmt.Printf("✅ No key → Partition: %d, Offset: %d\n", partition, offset)
		}
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println()
	fmt.Println("🎯 Key Learning Points:")
	fmt.Println("1. Same key always goes to same partition (ordering guarantee)")
	fmt.Println("2. Different keys are distributed across partitions (parallelism)")
	fmt.Println("3. No key uses round-robin distribution")
	fmt.Println("4. Partition assignment is deterministic based on key hash")
	fmt.Println()
	fmt.Println("💡 Next: Run the consumer to see how messages are processed!")
}
