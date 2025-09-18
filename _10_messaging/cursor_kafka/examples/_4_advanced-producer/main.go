package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	fmt.Println("🚀 Advanced Producer Configuration Demo")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Exploring different producer configurations...")
	fmt.Println()

	// Test different configurations
	testBatchingAndCompression()
	testAcknowledgmentModes()
	testIdempotentProducer()
}

func testBatchingAndCompression() {
	fmt.Println("📦 Test 1: Batching and Compression")
	fmt.Println(strings.Repeat("-", 40))

	// Configuration 1: No batching, no compression
	fmt.Println("🔧 Configuration 1: No batching, no compression")
	config1 := sarama.NewConfig()
	config1.Producer.Return.Successes = true
	config1.Producer.Flush.Frequency = 0 // No batching
	config1.Producer.Compression = sarama.CompressionNone

	producer1, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config1)
	if err != nil {
		log.Printf("Error creating producer 1: %v", err)
		return
	}
	defer producer1.Close()

	start := time.Now()
	for i := 0; i < 100; i++ {
		message := &sarama.ProducerMessage{
			Topic: "advanced-demo",
			Key:   sarama.StringEncoder(fmt.Sprintf("key-%d", i)),
			Value: sarama.StringEncoder(fmt.Sprintf("Message %d with no batching and no compression", i)),
		}
		producer1.SendMessage(message)
	}
	duration1 := time.Since(start)
	fmt.Printf("✅ Sent 100 messages in %v (no batching, no compression)\n", duration1)

	// Configuration 2: With batching, no compression
	fmt.Println("\n🔧 Configuration 2: With batching, no compression")
	config2 := sarama.NewConfig()
	config2.Producer.Return.Successes = true
	config2.Producer.Flush.Frequency = 100 * time.Millisecond
	config2.Producer.Flush.Messages = 10
	config2.Producer.Compression = sarama.CompressionNone

	producer2, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config2)
	if err != nil {
		log.Printf("Error creating producer 2: %v", err)
		return
	}
	defer producer2.Close()

	start = time.Now()
	for i := 0; i < 100; i++ {
		message := &sarama.ProducerMessage{
			Topic: "advanced-demo",
			Key:   sarama.StringEncoder(fmt.Sprintf("key-%d", i)),
			Value: sarama.StringEncoder(fmt.Sprintf("Message %d with batching and no compression", i)),
		}
		producer2.SendMessage(message)
	}
	duration2 := time.Since(start)
	fmt.Printf("✅ Sent 100 messages in %v (with batching, no compression)\n", duration2)

	// Configuration 3: With batching and compression
	fmt.Println("\n🔧 Configuration 3: With batching and compression")
	config3 := sarama.NewConfig()
	config3.Producer.Return.Successes = true
	config3.Producer.Flush.Frequency = 100 * time.Millisecond
	config3.Producer.Flush.Messages = 10
	config3.Producer.Compression = sarama.CompressionSnappy

	producer3, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config3)
	if err != nil {
		log.Printf("Error creating producer 3: %v", err)
		return
	}
	defer producer3.Close()

	start = time.Now()
	for i := 0; i < 100; i++ {
		message := &sarama.ProducerMessage{
			Topic: "advanced-demo",
			Key:   sarama.StringEncoder(fmt.Sprintf("key-%d", i)),
			Value: sarama.StringEncoder(fmt.Sprintf("Message %d with batching and Snappy compression", i)),
		}
		producer3.SendMessage(message)
	}
	duration3 := time.Since(start)
	fmt.Printf("✅ Sent 100 messages in %v (with batching and compression)\n", duration3)

	fmt.Printf("\n📊 Performance Comparison:\n")
	fmt.Printf("No batching, no compression: %v\n", duration1)
	fmt.Printf("With batching, no compression: %v\n", duration2)
	fmt.Printf("With batching and compression: %v\n", duration3)
}

func testAcknowledgmentModes() {
	fmt.Println("\n\n📨 Test 2: Acknowledgment Modes")
	fmt.Println(strings.Repeat("-", 40))

	// Test acks=0 (fire and forget)
	fmt.Println("🔧 Testing acks=0 (fire and forget)")
	config0 := sarama.NewConfig()
	config0.Producer.RequiredAcks = sarama.NoResponse
	config0.Producer.Return.Successes = true

	producer0, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config0)
	if err != nil {
		log.Printf("Error creating producer: %v", err)
		return
	}
	defer producer0.Close()

	start := time.Now()
	for i := 0; i < 50; i++ {
		message := &sarama.ProducerMessage{
			Topic: "advanced-demo",
			Key:   sarama.StringEncoder(fmt.Sprintf("acks0-key-%d", i)),
			Value: sarama.StringEncoder(fmt.Sprintf("Message %d with acks=0", i)),
		}
		producer0.SendMessage(message)
	}
	duration0 := time.Since(start)
	fmt.Printf("✅ Sent 50 messages in %v (acks=0)\n", duration0)

	// Test acks=1 (leader acknowledgment)
	fmt.Println("\n🔧 Testing acks=1 (leader acknowledgment)")
	config1 := sarama.NewConfig()
	config1.Producer.RequiredAcks = sarama.WaitForLocal
	config1.Producer.Return.Successes = true

	producer1, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config1)
	if err != nil {
		log.Printf("Error creating producer: %v", err)
		return
	}
	defer producer1.Close()

	start = time.Now()
	for i := 0; i < 50; i++ {
		message := &sarama.ProducerMessage{
			Topic: "advanced-demo",
			Key:   sarama.StringEncoder(fmt.Sprintf("acks1-key-%d", i)),
			Value: sarama.StringEncoder(fmt.Sprintf("Message %d with acks=1", i)),
		}
		producer1.SendMessage(message)
	}
	duration1 := time.Since(start)
	fmt.Printf("✅ Sent 50 messages in %v (acks=1)\n", duration1)

	// Test acks=all (ISR acknowledgment)
	fmt.Println("\n🔧 Testing acks=all (ISR acknowledgment)")
	configAll := sarama.NewConfig()
	configAll.Producer.RequiredAcks = sarama.WaitForAll
	configAll.Producer.Return.Successes = true

	producerAll, err := sarama.NewSyncProducer([]string{"localhost:9092"}, configAll)
	if err != nil {
		log.Printf("Error creating producer: %v", err)
		return
	}
	defer producerAll.Close()

	start = time.Now()
	for i := 0; i < 50; i++ {
		message := &sarama.ProducerMessage{
			Topic: "advanced-demo",
			Key:   sarama.StringEncoder(fmt.Sprintf("acksAll-key-%d", i)),
			Value: sarama.StringEncoder(fmt.Sprintf("Message %d with acks=all", i)),
		}
		producerAll.SendMessage(message)
	}
	durationAll := time.Since(start)
	fmt.Printf("✅ Sent 50 messages in %v (acks=all)\n", durationAll)

	fmt.Printf("\n📊 Acknowledgment Mode Comparison:\n")
	fmt.Printf("acks=0 (fire and forget): %v (fastest, no durability guarantee)\n", duration0)
	fmt.Printf("acks=1 (leader ack): %v (balanced)\n", duration1)
	fmt.Printf("acks=all (ISR ack): %v (slowest, highest durability)\n", durationAll)
}

func testIdempotentProducer() {
	fmt.Println("\n\n🔄 Test 3: Idempotent Producer")
	fmt.Println(strings.Repeat("-", 40))

	// Create idempotent producer
	config := sarama.NewConfig()
	config.Producer.Idempotent = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 3
	config.Producer.Retry.Backoff = 100 * time.Millisecond
	config.Producer.Return.Successes = true
	config.Net.MaxOpenRequests = 1 // Required for idempotent producer

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Printf("Error creating idempotent producer: %v", err)
		return
	}
	defer producer.Close()

	fmt.Println("🔧 Testing idempotent producer with duplicate messages...")

	// Send same message multiple times
	message := &sarama.ProducerMessage{
		Topic: "advanced-demo",
		Key:   sarama.StringEncoder("idempotent-key"),
		Value: sarama.StringEncoder("This message should only appear once despite multiple sends"),
	}

	for i := 0; i < 3; i++ {
		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Printf("Error sending message: %v", err)
		} else {
			fmt.Printf("✅ Send %d: Partition %d, Offset %d\n", i+1, partition, offset)
		}
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("\n💡 Key Learning Points:")
	fmt.Println("• Batching improves throughput but increases latency")
	fmt.Println("• Compression reduces network usage but increases CPU usage")
	fmt.Println("• acks=0: Fastest but no durability guarantee")
	fmt.Println("• acks=1: Balanced performance and durability")
	fmt.Println("• acks=all: Slowest but highest durability")
	fmt.Println("• Idempotent producer prevents duplicate messages")
}
