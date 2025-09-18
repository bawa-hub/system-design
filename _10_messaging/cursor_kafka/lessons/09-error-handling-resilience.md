# Lesson 9: Error Handling and Resilience

## 🎯 Learning Objectives
By the end of this lesson, you will understand:
- Error handling patterns and strategies for Kafka applications
- Retry mechanisms and exponential backoff
- Circuit breaker pattern implementation
- Dead letter queues and error recovery
- Monitoring and alerting for errors
- Resilience testing and chaos engineering
- Production-ready error handling best practices

## 📚 Theory Section

### 1. **Error Handling Fundamentals**

#### **Types of Errors in Kafka Applications:**

**Retriable Errors:**
- Network timeouts
- Broker unavailable
- Leader not available
- Not leader for partition
- Request timed out
- Network exceptions

**Non-Retriable Errors:**
- Invalid topic
- Invalid partition
- Authentication failures
- Authorization failures
- Serialization errors
- Message too large

**Temporary Errors:**
- Consumer group rebalancing
- Broker maintenance
- Network partitions
- Resource exhaustion

#### **Error Handling Strategy:**
1. **Classify the error** (retriable vs non-retriable)
2. **Apply appropriate strategy** (retry, skip, or fail)
3. **Log and monitor** error patterns
4. **Implement recovery** mechanisms

### 2. **Retry Mechanisms**

#### **Exponential Backoff:**
```go
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
```

#### **Retry Implementation:**
```go
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
```

#### **Retry with Context:**
```go
func retryWithContext(ctx context.Context, operation func() error, config RetryConfig) error {
    for attempt := 0; attempt <= config.MaxRetries; attempt++ {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
        }
        
        err := operation()
        if err == nil {
            return nil
        }
        
        if !isRetriableError(err) {
            return err
        }
        
        if attempt < config.MaxRetries {
            delay := config.CalculateDelay(attempt)
            select {
            case <-time.After(delay):
            case <-ctx.Done():
                return ctx.Err()
            }
        }
    }
    
    return fmt.Errorf("operation failed after %d retries", config.MaxRetries)
}
```

### 3. **Circuit Breaker Pattern**

#### **Circuit Breaker States:**
- **Closed**: Normal operation, requests pass through
- **Open**: Circuit is open, requests fail fast
- **Half-Open**: Testing if service has recovered

#### **Circuit Breaker Implementation:**
```go
type CircuitBreaker struct {
    name          string
    maxFailures   int
    timeout       time.Duration
    resetTimeout  time.Duration
    
    state         string // "closed", "open", "half-open"
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
            return errors.New("circuit breaker is open")
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
```

### 4. **Dead Letter Queues (DLQ)**

#### **DLQ Implementation:**
```go
type DeadLetterQueue struct {
    producer    sarama.SyncProducer
    dlqTopic    string
    maxRetries  int
}

func (dlq *DeadLetterQueue) SendToDLQ(message *sarama.ConsumerMessage, error error) error {
    dlqMessage := &sarama.ProducerMessage{
        Topic: dlq.dlqTopic,
        Key:   sarama.StringEncoder(fmt.Sprintf("dlq-%s", message.Key)),
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
```

#### **DLQ Consumer:**
```go
func (dlq *DeadLetterQueue) ProcessDLQ() error {
    consumer, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "dlq-processor", nil)
    if err != nil {
        return err
    }
    defer consumer.Close()
    
    handler := &DLQHandler{}
    return consumer.Consume(context.Background(), []string{dlq.dlqTopic}, handler)
}

type DLQHandler struct{}

func (h *DLQHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
    for message := <-claim.Messages() {
        if message == nil {
            return nil
        }
        
        // Process DLQ message
        h.processDLQMessage(message)
        session.MarkMessage(message, "")
    }
    return nil
}
```

### 5. **Error Monitoring and Alerting**

#### **Error Metrics:**
```go
type ErrorMetrics struct {
    TotalErrors      int64
    RetriableErrors  int64
    NonRetriableErrors int64
    CircuitBreakerTrips int64
    DLQMessages      int64
    LastError        time.Time
    ErrorRate        float64
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
```

#### **Alerting System:**
```go
type AlertManager struct {
    errorThreshold    float64
    circuitBreakerThreshold int
    dlqThreshold      int
    alertChannel      chan Alert
}

type Alert struct {
    Type        string
    Message     string
    Severity    string
    Timestamp   time.Time
    Metrics     ErrorMetrics
}

func (am *AlertManager) CheckAlerts(metrics ErrorMetrics) {
    if metrics.ErrorRate > am.errorThreshold {
        am.sendAlert(Alert{
            Type: "HIGH_ERROR_RATE",
            Message: fmt.Sprintf("Error rate %.2f%% exceeds threshold", metrics.ErrorRate),
            Severity: "CRITICAL",
            Timestamp: time.Now(),
            Metrics: metrics,
        })
    }
    
    if metrics.CircuitBreakerTrips > int64(am.circuitBreakerThreshold) {
        am.sendAlert(Alert{
            Type: "CIRCUIT_BREAKER_TRIPS",
            Message: "Circuit breaker has tripped multiple times",
            Severity: "WARNING",
            Timestamp: time.Now(),
            Metrics: metrics,
        })
    }
}
```

### 6. **Resilience Testing**

#### **Chaos Engineering:**
```go
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
        return errors.New("chaos engineering: injected failure")
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
```

#### **Resilience Test Suite:**
```go
func runResilienceTests() {
    tests := []struct {
        name        string
        description string
        testFunc    func() error
    }{
        {
            name: "Network Partition",
            description: "Simulate network partition between producer and broker",
            testFunc: testNetworkPartition,
        },
        {
            name: "Broker Failure",
            description: "Simulate broker failure and recovery",
            testFunc: testBrokerFailure,
        },
        {
            name: "Consumer Group Rebalancing",
            description: "Test consumer group behavior during rebalancing",
            testFunc: testConsumerGroupRebalancing,
        },
        {
            name: "High Error Rate",
            description: "Test system behavior under high error rates",
            testFunc: testHighErrorRate,
        },
    }
    
    for _, test := range tests {
        fmt.Printf("Running test: %s\n", test.name)
        if err := test.testFunc(); err != nil {
            fmt.Printf("Test failed: %v\n", err)
        } else {
            fmt.Printf("Test passed: %s\n", test.name)
        }
    }
}
```

### 7. **Production Best Practices**

#### **Error Handling Checklist:**
- [ ] **Classify errors** (retriable vs non-retriable)
- [ ] **Implement retry logic** with exponential backoff
- [ ] **Use circuit breakers** for external dependencies
- [ ] **Set up dead letter queues** for failed messages
- [ ] **Monitor error rates** and set up alerting
- [ ] **Test failure scenarios** regularly
- [ ] **Document error handling** strategies
- [ ] **Review and update** error handling regularly

#### **Configuration Management:**
```go
type ErrorHandlingConfig struct {
    RetryConfig      RetryConfig
    CircuitBreaker   CircuitBreakerConfig
    DLQConfig        DLQConfig
    MonitoringConfig MonitoringConfig
}

func LoadErrorHandlingConfig(env string) ErrorHandlingConfig {
    switch env {
    case "development":
        return getDevelopmentConfig()
    case "staging":
        return getStagingConfig()
    case "production":
        return getProductionConfig()
    default:
        return getDefaultConfig()
    }
}
```

## 🧪 Hands-on Experiments

### Experiment 1: Retry Mechanism Implementation

**Goal**: Implement robust retry logic with exponential backoff

**Steps**:
1. Create retry configuration
2. Implement exponential backoff
3. Test with different error scenarios
4. Measure retry effectiveness

**Expected Results**:
- Retry logic handles temporary failures
- Exponential backoff prevents overwhelming services
- Jitter prevents thundering herd problems

### Experiment 2: Circuit Breaker Implementation

**Goal**: Implement circuit breaker pattern for fault tolerance

**Steps**:
1. Create circuit breaker implementation
2. Test state transitions
3. Simulate failure scenarios
4. Verify recovery behavior

**Expected Results**:
- Circuit breaker opens on high failure rate
- Fast failure when circuit is open
- Automatic recovery testing

### Experiment 3: Dead Letter Queue Setup

**Goal**: Set up dead letter queue for failed messages

**Steps**:
1. Create DLQ topic and producer
2. Implement DLQ message handling
3. Test message routing to DLQ
4. Set up DLQ monitoring

**Expected Results**:
- Failed messages are routed to DLQ
- DLQ messages contain error information
- DLQ can be processed separately

### Experiment 4: Resilience Testing

**Goal**: Test system resilience under failure conditions

**Steps**:
1. Implement chaos engineering
2. Test different failure scenarios
3. Measure system recovery time
4. Validate error handling

**Expected Results**:
- System handles failures gracefully
- Recovery time is within acceptable limits
- Error handling prevents data loss

## 📊 Error Monitoring Dashboard

### **Key Error Metrics:**
- **Error Rate**: Percentage of failed operations
- **Retry Rate**: Percentage of operations that need retry
- **Circuit Breaker Trips**: Number of circuit breaker activations
- **DLQ Messages**: Number of messages sent to dead letter queue
- **Recovery Time**: Time to recover from failures

### **Alerting Thresholds:**
- **Error Rate > 5%**: Warning
- **Error Rate > 10%**: Critical
- **Circuit Breaker Trips > 3**: Warning
- **DLQ Messages > 100**: Warning
- **Recovery Time > 30s**: Warning

## 🎯 Key Takeaways

1. **Error handling is critical** for production systems
2. **Retry mechanisms** must be implemented with care
3. **Circuit breakers** prevent cascade failures
4. **Dead letter queues** ensure no data loss
5. **Monitoring and alerting** are essential
6. **Resilience testing** validates error handling
7. **Configuration management** enables environment-specific behavior

## 📝 Lesson 9 Assessment Questions

1. **What are the different types of errors in Kafka applications?**
2. **How do you implement exponential backoff for retries?**
3. **What is the circuit breaker pattern and when to use it?**
4. **How do you set up dead letter queues for failed messages?**
5. **What metrics should you monitor for error handling?**
6. **How do you test system resilience?**
7. **What are the best practices for production error handling?**
8. **How do you prevent thundering herd problems?**

---

## 🔄 Next Lesson Preview: Monitoring and Observability

**What we'll learn**:
- Comprehensive monitoring strategies
- Metrics collection and analysis
- Logging and tracing
- Alerting and incident response
- Performance dashboards
- Health checks and status pages

**Hands-on experiments**:
- Set up monitoring infrastructure
- Implement custom metrics
- Create alerting rules
- Build performance dashboards
