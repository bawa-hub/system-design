package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	fmt.Println("🛡️ Error Handling and Resilience Demo")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Building robust, fault-tolerant Kafka applications...")
	fmt.Println()

	// Run different error handling demonstrations
	demonstrateRetryMechanism()
	demonstrateCircuitBreaker()
	demonstrateDeadLetterQueue()
	demonstrateResilienceTesting()
}

func demonstrateRetryMechanism() {
	fmt.Println("🔄 Retry Mechanism Demonstration")
	fmt.Println(strings.Repeat("-", 40))

	// Create retry configuration
	retryConfig := RetryConfig{
		MaxRetries:   3,
		InitialDelay: 100 * time.Millisecond,
		MaxDelay:     5 * time.Second,
		Multiplier:   2.0,
		Jitter:       true,
	}

	// Test retry with different error scenarios
	scenarios := []struct {
		name        string
		errorType   string
		shouldRetry bool
	}{
		{"Network Timeout", "timeout", true},
		{"Broker Unavailable", "broker_unavailable", true},
		{"Invalid Topic", "invalid_topic", false},
		{"Authentication Failed", "auth_failed", false},
	}

	for _, scenario := range scenarios {
		fmt.Printf("🔧 Testing %s scenario...\n", scenario.name)
		
		attempts := 0
		err := retryOperation(func() error {
			attempts++
			fmt.Printf("   Attempt %d: ", attempts)
			
			if scenario.shouldRetry && attempts < 3 {
				fmt.Printf("Simulated %s error (will retry)\n", scenario.errorType)
				return fmt.Errorf("simulated %s error", scenario.errorType)
			} else if scenario.shouldRetry {
				fmt.Printf("Success after retries\n")
				return nil
			} else {
				fmt.Printf("Non-retriable %s error\n", scenario.errorType)
				return fmt.Errorf("non-retriable %s error", scenario.errorType)
			}
		}, retryConfig)
		
		if err != nil {
			fmt.Printf("   ❌ Final result: %v\n", err)
		} else {
			fmt.Printf("   ✅ Final result: Success\n")
		}
		fmt.Println()
	}
}

func demonstrateCircuitBreaker() {
	fmt.Println("⚡ Circuit Breaker Demonstration")
	fmt.Println(strings.Repeat("-", 40))

	// Create circuit breaker
	cb := &CircuitBreaker{
		name:          "kafka-producer",
		maxFailures:   3,
		timeout:       10 * time.Second,
		resetTimeout:  30 * time.Second,
		state:         "closed",
		failureCount:  0,
	}

	// Simulate operations with different success rates
	operations := []struct {
		name        string
		successRate float64
		count       int
	}{
		{"Normal Operations", 0.9, 10},
		{"Degrading Service", 0.3, 10},
		{"Service Recovery", 0.8, 10},
	}

	for _, op := range operations {
		fmt.Printf("🔧 Testing %s (%.0f%% success rate)...\n", op.name, op.successRate*100)
		
		for i := 0; i < op.count; i++ {
			err := cb.Execute(func() error {
				if rand.Float64() < op.successRate {
					return nil
				}
				return fmt.Errorf("operation failed")
			})
			
			state := cb.GetState()
			if err != nil {
				fmt.Printf("   Operation %d: Failed (State: %s)\n", i+1, state)
			} else {
				fmt.Printf("   Operation %d: Success (State: %s)\n", i+1, state)
			}
			
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println()
	}
}

func demonstrateDeadLetterQueue() {
	fmt.Println("💀 Dead Letter Queue Demonstration")
	fmt.Println(strings.Repeat("-", 40))

	// Create DLQ producer
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Printf("Error creating DLQ producer: %v", err)
		return
	}
	defer producer.Close()

	dlq := &DeadLetterQueue{
		producer:   producer,
		dlqTopic:   "dead-letter-queue",
		maxRetries: 3,
	}

	// Simulate failed messages
	failedMessages := []struct {
		topic     string
		partition int32
		offset    int64
		key       string
		value     string
		error     string
	}{
		{"user-events", 0, 100, "user-1", "Invalid JSON data", "JSON parsing error"},
		{"order-events", 1, 200, "order-2", "Corrupted message", "Deserialization error"},
		{"payment-events", 2, 300, "payment-3", "Missing required field", "Validation error"},
	}

	fmt.Println("🔧 Sending failed messages to DLQ...")
	for i, msg := range failedMessages {
		consumerMessage := &sarama.ConsumerMessage{
			Topic:     msg.topic,
			Partition: msg.partition,
			Offset:    msg.offset,
			Key:       []byte(msg.key),
			Value:     []byte(msg.value),
		}
		
		err := dlq.SendToDLQ(consumerMessage, fmt.Errorf(msg.error))
		if err != nil {
			fmt.Printf("   ❌ Failed to send message %d to DLQ: %v\n", i+1, err)
		} else {
			fmt.Printf("   ✅ Message %d sent to DLQ successfully\n", i+1)
		}
	}
	fmt.Println()
}

func demonstrateResilienceTesting() {
	fmt.Println("🧪 Resilience Testing Demonstration")
	fmt.Println(strings.Repeat("-", 40))

	// Create chaos engine
	chaosEngine := &ChaosEngine{
		enabled:     false,
		failureRate: 0.0,
		latency:     0,
	}

	// Test normal operation
	fmt.Println("🔧 Testing normal operation...")
	for i := 0; i < 5; i++ {
		err := chaosEngine.InjectFailure()
		if err != nil {
			fmt.Printf("   Operation %d: Failed - %v\n", i+1, err)
		} else {
			fmt.Printf("   Operation %d: Success\n", i+1)
		}
	}

	// Enable chaos engineering
	fmt.Println("\n🔧 Enabling chaos engineering (30% failure rate)...")
	chaosEngine.EnableChaos(0.3, 100*time.Millisecond)

	// Test with chaos
	for i := 0; i < 10; i++ {
		err := chaosEngine.InjectFailure()
		if err != nil {
			fmt.Printf("   Operation %d: Failed - %v\n", i+1, err)
		} else {
			fmt.Printf("   Operation %d: Success\n", i+1)
		}
		time.Sleep(50 * time.Millisecond)
	}

	// Disable chaos
	fmt.Println("\n🔧 Disabling chaos engineering...")
	chaosEngine.DisableChaos()

	// Test recovery
	fmt.Println("🔧 Testing recovery...")
	for i := 0; i < 5; i++ {
		err := chaosEngine.InjectFailure()
		if err != nil {
			fmt.Printf("   Operation %d: Failed - %v\n", i+1, err)
		} else {
			fmt.Printf("   Operation %d: Success\n", i+1)
		}
	}
	fmt.Println()
}

// Retry configuration and implementation
type RetryConfig struct {
	MaxRetries    int
	InitialDelay  time.Duration
	MaxDelay      time.Duration
	Multiplier    float64
	Jitter        bool
}

func (r *RetryConfig) CalculateDelay(attempt int) time.Duration {
	delay := float64(r.InitialDelay) * math.Pow(r.Multiplier, float64(attempt))
	if delay > float64(r.MaxDelay) {
		delay = float64(r.MaxDelay)
	}
	
	if r.Jitter {
		// Add random jitter to prevent thundering herd
		jitter := delay * 0.1 * (rand.Float64() - 0.5)
		delay += jitter
	}
	
	return time.Duration(delay)
}

func retryOperation(operation func() error, config RetryConfig) error {
	var lastErr error
	
	for attempt := 0; attempt <= config.MaxRetries; attempt++ {
		err := operation()
		if err == nil {
			return nil
		}
		
		lastErr = err
		
		if !isRetriableError(err) {
			return err
		}
		
		if attempt < config.MaxRetries {
			delay := config.CalculateDelay(attempt)
			time.Sleep(delay)
		}
	}
	
	return fmt.Errorf("operation failed after %d retries: %w", config.MaxRetries, lastErr)
}

func isRetriableError(err error) bool {
	if err == nil {
		return false
	}
	
	errorStr := err.Error()
	retriableErrors := []string{
		"timeout",
		"broker_unavailable",
		"leader_not_available",
		"network",
		"connection",
	}
	
	for _, retriableError := range retriableErrors {
		if strings.Contains(strings.ToLower(errorStr), retriableError) {
			return true
		}
	}
	
	return false
}

// Circuit breaker implementation
type CircuitBreaker struct {
	name          string
	maxFailures   int
	timeout       time.Duration
	resetTimeout  time.Duration
	
	state         string
	failureCount  int
	lastFailure   time.Time
	nextAttempt   time.Time
	
	mutex         sync.RWMutex
}

func (cb *CircuitBreaker) Execute(operation func() error) error {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	
	if cb.state == "open" {
		if time.Now().Before(cb.nextAttempt) {
			return fmt.Errorf("circuit breaker is open")
		}
		cb.state = "half-open"
	}
	
	err := operation()
	
	if err != nil {
		cb.recordFailure()
		return err
	}
	
	cb.recordSuccess()
	return nil
}

func (cb *CircuitBreaker) recordFailure() {
	cb.failureCount++
	cb.lastFailure = time.Now()
	
	if cb.failureCount >= cb.maxFailures {
		cb.state = "open"
		cb.nextAttempt = time.Now().Add(cb.resetTimeout)
	}
}

func (cb *CircuitBreaker) recordSuccess() {
	cb.failureCount = 0
	cb.state = "closed"
}

func (cb *CircuitBreaker) GetState() string {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()
	return cb.state
}

// Dead letter queue implementation
type DeadLetterQueue struct {
	producer   sarama.SyncProducer
	dlqTopic   string
	maxRetries int
}

func (dlq *DeadLetterQueue) SendToDLQ(message *sarama.ConsumerMessage, error error) error {
	dlqMessage := &sarama.ProducerMessage{
		Topic: dlq.dlqTopic,
		Key:   sarama.StringEncoder(fmt.Sprintf("dlq-%s", string(message.Key))),
		Value: sarama.StringEncoder(fmt.Sprintf("Original: %s, Error: %s", 
			string(message.Value), error.Error())),
		Headers: []sarama.RecordHeader{
			{Key: []byte("original-topic"), Value: []byte(message.Topic)},
			{Key: []byte("original-partition"), Value: []byte(fmt.Sprintf("%d", message.Partition))},
			{Key: []byte("original-offset"), Value: []byte(fmt.Sprintf("%d", message.Offset))},
			{Key: []byte("error-message"), Value: []byte(error.Error())},
			{Key: []byte("timestamp"), Value: []byte(time.Now().Format(time.RFC3339))},
		},
	}
	
	_, _, err := dlq.producer.SendMessage(dlqMessage)
	return err
}

// Chaos engine implementation
type ChaosEngine struct {
	enabled     bool
	failureRate float64
	latency     time.Duration
	mutex       sync.RWMutex
}

func (ce *ChaosEngine) InjectFailure() error {
	ce.mutex.RLock()
	defer ce.mutex.RUnlock()
	
	if !ce.enabled {
		return nil
	}
	
	if rand.Float64() < ce.failureRate {
		return fmt.Errorf("chaos engineering: injected failure")
	}
	
	if ce.latency > 0 {
		time.Sleep(ce.latency)
	}
	
	return nil
}

func (ce *ChaosEngine) EnableChaos(failureRate float64, latency time.Duration) {
	ce.mutex.Lock()
	defer ce.mutex.Unlock()
	
	ce.enabled = true
	ce.failureRate = failureRate
	ce.latency = latency
}

func (ce *ChaosEngine) DisableChaos() {
	ce.mutex.Lock()
	defer ce.mutex.Unlock()
	
	ce.enabled = false
	ce.failureRate = 0.0
	ce.latency = 0
}

// Error metrics implementation
type ErrorMetrics struct {
	TotalErrors        int64
	RetriableErrors    int64
	NonRetriableErrors int64
	CircuitBreakerTrips int64
	DLQMessages        int64
	LastError          time.Time
	ErrorRate          float64
}

func (em *ErrorMetrics) RecordError(err error, isRetriable bool) {
	atomic.AddInt64(&em.TotalErrors, 1)
	if isRetriable {
		atomic.AddInt64(&em.RetriableErrors, 1)
	} else {
		atomic.AddInt64(&em.NonRetriableErrors, 1)
	}
	em.LastError = time.Now()
}

func (em *ErrorMetrics) RecordCircuitBreakerTrip() {
	atomic.AddInt64(&em.CircuitBreakerTrips, 1)
}

func (em *ErrorMetrics) RecordDLQMessage() {
	atomic.AddInt64(&em.DLQMessages, 1)
}

func (em *ErrorMetrics) GetErrorRate() float64 {
	total := atomic.LoadInt64(&em.TotalErrors)
	if total == 0 {
		return 0
	}
	return float64(total) / float64(total+1000) * 100 // Simplified calculation
}
