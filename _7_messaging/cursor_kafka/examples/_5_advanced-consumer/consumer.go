package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	fmt.Println("🎧 Advanced Consumer Configuration Demo")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Exploring different consumer configurations...")
	fmt.Println()

	// Test different configurations
	testOffsetStrategies()
	testConsumerGroupTuning()
	testConsumerLagMonitoring()
}

func testOffsetStrategies() {
	fmt.Println("📍 Test 1: Offset Management Strategies")
	fmt.Println(strings.Repeat("-", 40))

	// Test OffsetOldest
	fmt.Println("🔧 Testing OffsetOldest (process all historical messages)")
	configOldest := sarama.NewConfig()
	configOldest.Consumer.Offsets.Initial = sarama.OffsetOldest
	configOldest.Consumer.Return.Errors = true
	configOldest.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

	consumerOldest, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "offset-test-oldest", configOldest)
	if err != nil {
		log.Printf("Error creating consumer: %v", err)
		return
	}
	defer consumerOldest.Close()

	// Test OffsetNewest
	fmt.Println("\n🔧 Testing OffsetNewest (process only new messages)")
	configNewest := sarama.NewConfig()
	configNewest.Consumer.Offsets.Initial = sarama.OffsetNewest
	configNewest.Consumer.Return.Errors = true
	configNewest.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

	consumerNewest, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "offset-test-newest", configNewest)
	if err != nil {
		log.Printf("Error creating consumer: %v", err)
		return
	}
	defer consumerNewest.Close()

	fmt.Println("✅ Offset strategies configured successfully")
	fmt.Println("💡 OffsetOldest: Will process all historical messages")
	fmt.Println("💡 OffsetNewest: Will only process new messages")
}

func testConsumerGroupTuning() {
	fmt.Println("\n\n⚙️ Test 2: Consumer Group Tuning")
	fmt.Println(strings.Repeat("-", 40))

	// Default configuration
	fmt.Println("🔧 Default Configuration:")
	configDefault := sarama.NewConfig()
	configDefault.Consumer.Return.Errors = true
	configDefault.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	configDefault.Consumer.Group.Session.Timeout = 10 * time.Second
	configDefault.Consumer.Group.Heartbeat.Interval = 3 * time.Second
	configDefault.Consumer.Fetch.Default = 1024 * 1024 // 1MB
	configDefault.Consumer.MaxWaitTime = 250 * time.Millisecond

	fmt.Printf("   Session Timeout: %v\n", configDefault.Consumer.Group.Session.Timeout)
	fmt.Printf("   Heartbeat Interval: %v\n", configDefault.Consumer.Group.Heartbeat.Interval)
	fmt.Printf("   Fetch Default: %d bytes\n", configDefault.Consumer.Fetch.Default)
	fmt.Printf("   Max Wait Time: %v\n", configDefault.Consumer.MaxWaitTime)

	// Tuned configuration
	fmt.Println("\n🔧 Tuned Configuration (High Throughput):")
	configTuned := sarama.NewConfig()
	configTuned.Consumer.Return.Errors = true
	configTuned.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	configTuned.Consumer.Group.Session.Timeout = 30 * time.Second
	configTuned.Consumer.Group.Heartbeat.Interval = 10 * time.Second
	configTuned.Consumer.Fetch.Default = 10 * 1024 * 1024 // 10MB
	configTuned.Consumer.Fetch.Max = 50 * 1024 * 1024 // 50MB
	configTuned.Consumer.MaxWaitTime = 500 * time.Millisecond
	configTuned.ChannelBufferSize = 1024

	fmt.Printf("   Session Timeout: %v\n", configTuned.Consumer.Group.Session.Timeout)
	fmt.Printf("   Heartbeat Interval: %v\n", configTuned.Consumer.Group.Heartbeat.Interval)
	fmt.Printf("   Fetch Default: %d bytes\n", configTuned.Consumer.Fetch.Default)
	fmt.Printf("   Fetch Max: %d bytes\n", configTuned.Consumer.Fetch.Max)
	fmt.Printf("   Max Wait Time: %v\n", configTuned.Consumer.MaxWaitTime)
	fmt.Printf("   Channel Buffer Size: %d\n", configTuned.ChannelBufferSize)

	fmt.Println("\n💡 Tuned configuration optimizes for:")
	fmt.Println("   • Higher throughput (larger fetch sizes)")
	fmt.Println("   • Better fault tolerance (longer timeouts)")
	fmt.Println("   • Reduced rebalancing (sticky strategy)")
	fmt.Println("   • More memory usage (larger buffers)")
}

func testConsumerLagMonitoring() {
	fmt.Println("\n\n📊 Test 3: Consumer Lag Monitoring")
	fmt.Println(strings.Repeat("-", 40))

	// Create a consumer with lag monitoring
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

	consumer, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "lag-monitor-group", config)
	if err != nil {
		log.Printf("Error creating consumer: %v", err)
		return
	}
	defer consumer.Close()

	// Create lag monitor
	lagMonitor := &ConsumerLagMonitor{
		consumerGroup: "lag-monitor-group",
		topic:         "advanced-demo",
		healthy:       true,
	}

	// Start consumer with lag monitoring
	handler := &LagMonitoringHandler{lagMonitor: lagMonitor}

	// Setup signal handling
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-sigchan:
				fmt.Println("🛑 Stopping consumer...")
				return
			default:
				ctx := context.Background()
				err := consumer.Consume(ctx, []string{"advanced-demo"}, handler)
				if err != nil {
					log.Printf("Error from consumer: %v", err)
					return
				}
			}
		}
	}()

	go func() {
		for err := range consumer.Errors() {
			log.Printf("Consumer error: %v", err)
		}
	}()

	// Start lag monitoring
	go lagMonitor.StartMonitoring()

	fmt.Println("🎧 Consumer with lag monitoring started!")
	fmt.Println("📊 Lag monitoring will check consumer health every 5 seconds")
	fmt.Println("Press Ctrl+C to stop")

	wg.Wait()
}

// ConsumerLagMonitor monitors consumer lag
type ConsumerLagMonitor struct {
	consumerGroup string
	topic         string
	healthy       bool
	lastCheck     time.Time
}

func (m *ConsumerLagMonitor) StartMonitoring() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		m.CheckHealth()
	}
}

func (m *ConsumerLagMonitor) CheckHealth() {
	// Simulate lag check (in real implementation, you'd query Kafka)
	// For demo purposes, we'll simulate different lag scenarios
	
	lag := m.simulateLagCheck()
	
	if lag > 1000 {
		if m.healthy {
			fmt.Printf("🚨 ALERT: High consumer lag detected: %d messages\n", lag)
			m.healthy = false
		}
	} else if lag > 100 {
		if m.healthy {
			fmt.Printf("⚠️  WARNING: Consumer lag increasing: %d messages\n", lag)
		}
	} else {
		if !m.healthy {
			fmt.Printf("✅ Consumer lag back to normal: %d messages\n", lag)
			m.healthy = true
		}
	}
	
	m.lastCheck = time.Now()
}

func (m *ConsumerLagMonitor) simulateLagCheck() int {
	// Simulate different lag scenarios for demo
	now := time.Now()
	seconds := now.Second()
	
	if seconds < 10 {
		return 50 // Normal lag
	} else if seconds < 20 {
		return 500 // Warning lag
	} else if seconds < 30 {
		return 2000 // Critical lag
	} else if seconds < 40 {
		return 1500 // Still high
	} else {
		return 100 // Back to normal
	}
}

// LagMonitoringHandler handles messages with lag monitoring
type LagMonitoringHandler struct {
	lagMonitor *ConsumerLagMonitor
}

func (h *LagMonitoringHandler) Setup(sarama.ConsumerGroupSession) error {
	fmt.Println("🔧 Consumer group session started")
	return nil
}

func (h *LagMonitoringHandler) Cleanup(sarama.ConsumerGroupSession) error {
	fmt.Println("🧹 Consumer group session ended")
	return nil
}

func (h *LagMonitoringHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			if message == nil {
				return nil
			}

			// Process message
			fmt.Printf("📨 Processing message: Partition %d, Offset %d\n", 
				message.Partition, message.Offset)

			// Simulate processing time
			time.Sleep(100 * time.Millisecond)

			// Mark message as processed
			session.MarkMessage(message, "")

		case <-session.Context().Done():
			return nil
		}
	}
}
