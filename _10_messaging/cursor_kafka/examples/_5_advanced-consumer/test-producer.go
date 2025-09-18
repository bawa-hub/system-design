package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	fmt.Println("📤 Consumer Test Producer")
	fmt.Println(strings.Repeat("=", 30))
	fmt.Println("Sending messages for consumer testing...")
	fmt.Println()

	// Create producer
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer producer.Close()

	topic := "advanced-demo"

	// Send messages at different rates to test consumer lag
	fmt.Println("📤 Sending messages at different rates...")
	
	// Phase 1: Normal rate (1 message per second)
	fmt.Println("Phase 1: Normal rate (1 msg/sec for 10 seconds)")
	for i := 0; i < 10; i++ {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(fmt.Sprintf("normal-key-%d", i)),
			Value: sarama.StringEncoder(fmt.Sprintf("Normal rate message %d", i)),
		}
		producer.SendMessage(message)
		fmt.Printf("✅ Sent normal message %d\n", i+1)
		time.Sleep(1 * time.Second)
	}

	// Phase 2: High rate (5 messages per second)
	fmt.Println("\nPhase 2: High rate (5 msg/sec for 10 seconds)")
	for i := 0; i < 50; i++ {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(fmt.Sprintf("high-rate-key-%d", i)),
			Value: sarama.StringEncoder(fmt.Sprintf("High rate message %d", i)),
		}
		producer.SendMessage(message)
		if i%5 == 0 {
			fmt.Printf("✅ Sent high rate message %d\n", i+1)
		}
		time.Sleep(200 * time.Millisecond)
	}

	// Phase 3: Very high rate (10 messages per second)
	fmt.Println("\nPhase 3: Very high rate (10 msg/sec for 10 seconds)")
	for i := 0; i < 100; i++ {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(fmt.Sprintf("very-high-key-%d", i)),
			Value: sarama.StringEncoder(fmt.Sprintf("Very high rate message %d", i)),
		}
		producer.SendMessage(message)
		if i%10 == 0 {
			fmt.Printf("✅ Sent very high rate message %d\n", i+1)
		}
		time.Sleep(100 * time.Millisecond)
	}

	// Phase 4: Back to normal rate
	fmt.Println("\nPhase 4: Back to normal rate (1 msg/sec for 10 seconds)")
	for i := 0; i < 10; i++ {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(fmt.Sprintf("back-normal-key-%d", i)),
			Value: sarama.StringEncoder(fmt.Sprintf("Back to normal message %d", i)),
		}
		producer.SendMessage(message)
		fmt.Printf("✅ Sent back to normal message %d\n", i+1)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("\n🎉 All messages sent!")
	fmt.Println("💡 Watch the consumer lag monitoring in the other terminal!")
}
