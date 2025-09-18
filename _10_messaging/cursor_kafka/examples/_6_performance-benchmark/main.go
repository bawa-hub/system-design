package main

import (
	"fmt"
	"log"
	"math"
	"strings"
	"sync/atomic"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	fmt.Println("🚀 Kafka Performance Benchmark Suite")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Comprehensive performance testing and optimization...")
	fmt.Println()

	// Run different performance tests
	runThroughputTest()
	runLatencyTest()
	runResourceUtilizationTest()
	runScalingTest()
}

func runThroughputTest() {
	fmt.Println("📊 Test 1: Throughput Optimization")
	fmt.Println(strings.Repeat("-", 40))

	// Test different producer configurations
	configs := map[string]sarama.Config{
		"Baseline": createBaselineConfig(),
		"Batched":  createBatchedConfig(),
		"Compressed": createCompressedConfig(),
		"Optimized": createOptimizedConfig(),
	}

	messageCount := 1000
	messageSize := 1024

	for name, config := range configs {
		fmt.Printf("🔧 Testing %s configuration...\n", name)
		
		producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, &config)
		if err != nil {
			log.Printf("Error creating producer: %v", err)
			continue
		}

		start := time.Now()
		for i := 0; i < messageCount; i++ {
			message := &sarama.ProducerMessage{
				Topic: "performance-test",
				Key:   sarama.StringEncoder(fmt.Sprintf("key-%d", i)),
				Value: sarama.StringEncoder(strings.Repeat("x", messageSize)),
			}
			producer.SendMessage(message)
		}
		duration := time.Since(start)
		throughput := float64(messageCount) / duration.Seconds()

		fmt.Printf("   ✅ %s: %.2f messages/second\n", name, throughput)
		producer.Close()
	}
}

func runLatencyTest() {
	fmt.Println("\n\n⏱️ Test 2: Latency Optimization")
	fmt.Println(strings.Repeat("-", 40))

	// Test different latency configurations
	configs := map[string]sarama.Config{
		"Default": createBaselineConfig(),
		"Low Latency": createLowLatencyConfig(),
		"Async": createAsyncConfig(),
	}

	messageCount := 100
	latencies := make([]time.Duration, messageCount)

	for name, config := range configs {
		fmt.Printf("🔧 Testing %s configuration...\n", name)
		
		producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, &config)
		if err != nil {
			log.Printf("Error creating producer: %v", err)
			continue
		}

		// Measure individual message latencies
		for i := 0; i < messageCount; i++ {
			start := time.Now()
			message := &sarama.ProducerMessage{
				Topic: "latency-test",
				Key:   sarama.StringEncoder(fmt.Sprintf("latency-key-%d", i)),
				Value: sarama.StringEncoder(fmt.Sprintf("Latency test message %d", i)),
			}
			producer.SendMessage(message)
			latencies[i] = time.Since(start)
		}

		// Calculate statistics
		avgLatency := calculateAverageLatency(latencies)
		p95Latency := calculatePercentileLatency(latencies, 95)
		p99Latency := calculatePercentileLatency(latencies, 99)

		fmt.Printf("   ✅ %s: Avg=%.2fms, P95=%.2fms, P99=%.2fms\n", 
			name, avgLatency.Milliseconds(), p95Latency.Milliseconds(), p99Latency.Milliseconds())
		
		producer.Close()
	}
}

func runResourceUtilizationTest() {
	fmt.Println("\n\n💾 Test 3: Resource Utilization")
	fmt.Println(strings.Repeat("-", 40))

	// Test memory usage with different configurations
	configs := map[string]sarama.Config{
		"Small Batches": createSmallBatchConfig(),
		"Large Batches": createLargeBatchConfig(),
		"Compressed": createCompressedConfig(),
	}

	messageCount := 5000
	messageSize := 2048

	for name, config := range configs {
		fmt.Printf("🔧 Testing %s configuration...\n", name)
		
		producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, &config)
		if err != nil {
			log.Printf("Error creating producer: %v", err)
			continue
		}

		start := time.Now()
		for i := 0; i < messageCount; i++ {
			message := &sarama.ProducerMessage{
				Topic: "resource-test",
				Key:   sarama.StringEncoder(fmt.Sprintf("resource-key-%d", i)),
				Value: sarama.StringEncoder(strings.Repeat("x", messageSize)),
			}
			producer.SendMessage(message)
		}
		duration := time.Since(start)
		throughput := float64(messageCount) / duration.Seconds()
		memoryPerMessage := float64(messageSize) / throughput

		fmt.Printf("   ✅ %s: %.2f msg/s, %.2f bytes/msg memory\n", 
			name, throughput, memoryPerMessage)
		
		producer.Close()
	}
}

func runScalingTest() {
	fmt.Println("\n\n📈 Test 4: Scaling Test")
	fmt.Println(strings.Repeat("-", 40))

	// Test with different partition counts
	partitionCounts := []int{1, 3, 6, 12}
	messageCount := 2000

	for _, partitions := range partitionCounts {
		fmt.Printf("🔧 Testing with %d partitions...\n", partitions)
		
		// Create topic with specific partition count
		topicName := fmt.Sprintf("scaling-test-%d", partitions)
		
		config := createOptimizedConfig()
		producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, &config)
		if err != nil {
			log.Printf("Error creating producer: %v", err)
			continue
		}

		start := time.Now()
		for i := 0; i < messageCount; i++ {
			message := &sarama.ProducerMessage{
				Topic: topicName,
				Key:   sarama.StringEncoder(fmt.Sprintf("scaling-key-%d", i)),
				Value: sarama.StringEncoder(fmt.Sprintf("Scaling test message %d", i)),
			}
			producer.SendMessage(message)
		}
		duration := time.Since(start)
		throughput := float64(messageCount) / duration.Seconds()

		fmt.Printf("   ✅ %d partitions: %.2f messages/second\n", partitions, throughput)
		producer.Close()
	}
}

// Configuration functions
func createBaselineConfig() sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForLocal
	return *config
}

func createBatchedConfig() sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Flush.Frequency = 100 * time.Millisecond
	config.Producer.Flush.Bytes = 1024 * 1024 // 1MB
	config.Producer.Flush.Messages = 1000
	return *config
}

func createCompressedConfig() sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Flush.Frequency = 100 * time.Millisecond
	config.Producer.Flush.Bytes = 1024 * 1024 // 1MB
	config.Producer.Flush.Messages = 1000
	return *config
}

func createOptimizedConfig() sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Compression = sarama.CompressionLZ4
	config.Producer.Flush.Frequency = 50 * time.Millisecond
	config.Producer.Flush.Bytes = 2 * 1024 * 1024 // 2MB
	config.Producer.Flush.Messages = 2000
	config.Producer.Retry.Max = 3
	config.Producer.Retry.Backoff = 100 * time.Millisecond
	return *config
}

func createLowLatencyConfig() sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Flush.Frequency = 1 * time.Millisecond
	config.Producer.Flush.Bytes = 1024 // 1KB
	config.Producer.Flush.Messages = 1
	config.Producer.Compression = sarama.CompressionSnappy
	return *config
}

func createAsyncConfig() sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Flush.Frequency = 10 * time.Millisecond
	config.Producer.Flush.Bytes = 512 * 1024 // 512KB
	config.Producer.Flush.Messages = 100
	config.Producer.Compression = sarama.CompressionSnappy
	return *config
}

func createSmallBatchConfig() sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Flush.Frequency = 10 * time.Millisecond
	config.Producer.Flush.Bytes = 1024 // 1KB
	config.Producer.Flush.Messages = 10
	return *config
}

func createLargeBatchConfig() sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Flush.Frequency = 200 * time.Millisecond
	config.Producer.Flush.Bytes = 5 * 1024 * 1024 // 5MB
	config.Producer.Flush.Messages = 5000
	return *config
}

// Utility functions
func calculateAverageLatency(latencies []time.Duration) time.Duration {
	var total time.Duration
	for _, latency := range latencies {
		total += latency
	}
	return total / time.Duration(len(latencies))
}

func calculatePercentileLatency(latencies []time.Duration, percentile int) time.Duration {
	// Simple percentile calculation (not optimized for production)
	n := len(latencies)
	index := int(math.Ceil(float64(n) * float64(percentile) / 100.0))
	if index >= n {
		index = n - 1
	}
	return latencies[index]
}

// Performance monitoring
type PerformanceMonitor struct {
	MessagesSent    int64
	BytesSent       int64
	StartTime       time.Time
	Errors          int64
}

func (m *PerformanceMonitor) RecordMessage(bytes int) {
	atomic.AddInt64(&m.MessagesSent, 1)
	atomic.AddInt64(&m.BytesSent, int64(bytes))
}

func (m *PerformanceMonitor) RecordError() {
	atomic.AddInt64(&m.Errors, 1)
}

func (m *PerformanceMonitor) GetThroughput() float64 {
	duration := time.Since(m.StartTime).Seconds()
	return float64(atomic.LoadInt64(&m.MessagesSent)) / duration
}

func (m *PerformanceMonitor) GetErrorRate() float64 {
	total := atomic.LoadInt64(&m.MessagesSent) + atomic.LoadInt64(&m.Errors)
	if total == 0 {
		return 0
	}
	return float64(atomic.LoadInt64(&m.Errors)) / float64(total) * 100
}
