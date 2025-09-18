# Lesson 7: Advanced Consumer Configuration

## 🎯 Learning Objectives
By the end of this lesson, you will understand:
- Advanced consumer configuration options
- Offset management strategies and best practices
- Consumer group tuning and performance optimization
- Error handling and retry mechanisms for consumers
- Consumer lag monitoring and management
- Memory and resource optimization for consumers

## 📚 Theory Section

### 1. **Consumer Configuration Deep Dive**

#### **Core Consumer Configuration Options**

**Basic Consumer Settings:**
```go
config := sarama.NewConfig()
config.Consumer.Return.Errors = true
config.Consumer.Offsets.Initial = sarama.OffsetOldest
config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
config.Consumer.Group.Session.Timeout = 10 * time.Second
config.Consumer.Group.Heartbeat.Interval = 3 * time.Second
```

**Fetch Configuration:**
```go
config.Consumer.Fetch.Min = 1
config.Consumer.Fetch.Default = 1024 * 1024 // 1MB
config.Consumer.Fetch.Max = 10 * 1024 * 1024 // 10MB
config.Consumer.MaxWaitTime = 250 * time.Millisecond
```

### 2. **Offset Management Strategies**

#### **Offset Initialization Options**

**OffsetOldest (Earliest):**
```go
config.Consumer.Offsets.Initial = sarama.OffsetOldest
```
- **Behavior**: Start from the beginning of the topic
- **Use case**: Processing all historical data
- **Risk**: May process duplicate messages

**OffsetNewest (Latest):**
```go
config.Consumer.Offsets.Initial = sarama.OffsetNewest
```
- **Behavior**: Start from the latest messages
- **Use case**: Real-time processing only
- **Risk**: May miss some messages

#### **Offset Commit Strategies**

**Automatic Commit (Default):**
```go
config.Consumer.Offsets.AutoCommit.Enable = true
config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second
```
- **Behavior**: Kafka automatically commits offsets
- **Performance**: Better performance
- **Risk**: May process messages twice if consumer fails

**Manual Commit:**
```go
config.Consumer.Offsets.AutoCommit.Enable = false
// Manually commit after processing
session.MarkMessage(message, "")
session.Commit()
```
- **Behavior**: Application controls when to commit
- **Performance**: Slightly lower performance
- **Benefit**: Exactly-once processing guarantee

#### **Offset Reset Strategies**

**Earliest (Reset to Beginning):**
```go
config.Consumer.Offsets.Initial = sarama.OffsetOldest
```
- **When**: Consumer group starts for the first time
- **Behavior**: Start from the beginning of the topic
- **Use case**: Processing all historical data

**Latest (Reset to End):**
```go
config.Consumer.Offsets.Initial = sarama.OffsetNewest
```
- **When**: Consumer group starts for the first time
- **Behavior**: Start from the latest messages
- **Use case**: Real-time processing only

**None (Fail if no offset):**
```go
config.Consumer.Offsets.Initial = sarama.OffsetNewest
config.Consumer.Offsets.AutoCommit.Enable = false
```
- **When**: Consumer group starts for the first time
- **Behavior**: Fail if no committed offset exists
- **Use case**: Ensuring no data loss

### 3. **Consumer Group Tuning**

#### **Session and Heartbeat Configuration**

**Session Timeout:**
```go
config.Consumer.Group.Session.Timeout = 10 * time.Second
```
- **Purpose**: Maximum time consumer can be silent before being considered dead
- **Trade-off**: Longer timeout = more fault tolerance, slower failure detection
- **Recommendation**: 10-30 seconds

**Heartbeat Interval:**
```go
config.Consumer.Group.Heartbeat.Interval = 3 * time.Second
```
- **Purpose**: How often consumer sends heartbeats
- **Rule**: Must be less than session timeout
- **Recommendation**: 1/3 of session timeout

#### **Rebalancing Configuration**

**Rebalancing Strategy:**
```go
// Round Robin (default)
config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

// Range
config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange

// Sticky
config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
```

**Strategy Comparison:**
| Strategy | Pros | Cons | Use Case |
|----------|------|------|----------|
| Round Robin | Even distribution | May cause more rebalancing | General purpose |
| Range | Simple, predictable | Uneven distribution | Few partitions |
| Sticky | Minimal partition movement | Complex implementation | Large consumer groups |

### 4. **Performance Optimization**

#### **Fetch Configuration Tuning**

**Min Fetch Size:**
```go
config.Consumer.Fetch.Min = 1
```
- **Purpose**: Minimum bytes to fetch per request
- **Impact**: Higher = better throughput, higher latency
- **Recommendation**: 1 byte (default)

**Default Fetch Size:**
```go
config.Consumer.Fetch.Default = 1024 * 1024 // 1MB
```
- **Purpose**: Default bytes to fetch per request
- **Impact**: Higher = better throughput, more memory usage
- **Recommendation**: 1-10MB depending on message size

**Max Fetch Size:**
```go
config.Consumer.Fetch.Max = 10 * 1024 * 1024 // 10MB
```
- **Purpose**: Maximum bytes to fetch per request
- **Impact**: Higher = better throughput, more memory usage
- **Recommendation**: 10-50MB depending on available memory

**Max Wait Time:**
```go
config.Consumer.MaxWaitTime = 250 * time.Millisecond
```
- **Purpose**: Maximum time to wait for data
- **Impact**: Higher = better throughput, higher latency
- **Recommendation**: 100-500ms

#### **Memory Optimization**

**Buffer Size:**
```go
config.ChannelBufferSize = 256
```
- **Purpose**: Size of internal channels
- **Impact**: Higher = better performance, more memory
- **Recommendation**: 256-1024

**Max Processing Time:**
```go
config.Consumer.Group.MaxProcessingTime = 25 * time.Second
```
- **Purpose**: Maximum time to process a batch of messages
- **Impact**: Higher = more fault tolerance, slower rebalancing
- **Recommendation**: 25-30 seconds

### 5. **Error Handling and Retry Mechanisms**

#### **Consumer Error Types**

**Retriable Errors:**
- Network timeouts
- Broker unavailable
- Leader not available
- Not leader for partition

**Non-Retriable Errors:**
- Invalid topic
- Invalid partition
- Authentication failures
- Authorization failures

#### **Error Handling Patterns**

**Retry with Exponential Backoff:**
```go
func handleConsumerError(err error, message *sarama.ConsumerMessage) {
    if retriableError(err) {
        // Exponential backoff retry
        retryCount := getRetryCount(message)
        if retryCount < maxRetries {
            delay := time.Duration(math.Pow(2, float64(retryCount))) * time.Second
            time.Sleep(delay)
            retryMessage(message)
        } else {
            // Send to dead letter queue
            sendToDLQ(message)
        }
    } else {
        // Log and skip non-retriable errors
        log.Printf("Non-retriable error: %v", err)
    }
}
```

**Circuit Breaker Pattern:**
```go
type CircuitBreaker struct {
    failureCount int
    lastFailure  time.Time
    state        string // "closed", "open", "half-open"
}

func (cb *CircuitBreaker) CanExecute() bool {
    if cb.state == "open" {
        if time.Since(cb.lastFailure) > 30*time.Second {
            cb.state = "half-open"
            return true
        }
        return false
    }
    return true
}
```

### 6. **Consumer Lag Monitoring**

#### **What is Consumer Lag?**
Consumer lag is the difference between the latest offset in a partition and the offset that the consumer has processed.

```
Partition: [msg1, msg2, msg3, msg4, msg5, msg6, msg7, msg8, msg9, msg10]
                                    ↑                    ↑
                              Consumer Offset      Latest Offset
                              (msg4)               (msg10)
                              Lag = 6 messages
```

#### **Lag Monitoring Implementation:**
```go
type ConsumerLagMonitor struct {
    consumerGroup string
    topic         string
    partition     int32
    currentOffset int64
    latestOffset  int64
}

func (m *ConsumerLagMonitor) GetLag() int64 {
    return m.latestOffset - m.currentOffset
}

func (m *ConsumerLagMonitor) IsHealthy() bool {
    return m.GetLag() < 1000 // Threshold for healthy lag
}
```

#### **Lag Alerting:**
```go
func (m *ConsumerLagMonitor) CheckLag() {
    lag := m.GetLag()
    if lag > 10000 {
        // Critical lag - send alert
        sendAlert(fmt.Sprintf("High consumer lag: %d messages", lag))
    } else if lag > 1000 {
        // Warning lag - send notification
        sendNotification(fmt.Sprintf("Consumer lag warning: %d messages", lag))
    }
}
```

## 🧪 Hands-on Experiments

### Experiment 1: Offset Management Strategies

**Goal**: Understand different offset initialization and commit strategies

**Steps**:
1. Create consumer with OffsetOldest
2. Create consumer with OffsetNewest
3. Test manual vs automatic commit
4. Observe behavior differences

**Expected Results**:
- OffsetOldest processes all historical messages
- OffsetNewest starts from latest messages
- Manual commit provides better control

### Experiment 2: Consumer Group Performance Tuning

**Goal**: Optimize consumer group performance

**Steps**:
1. Create consumer with default settings
2. Tune fetch configuration
3. Adjust session and heartbeat timeouts
4. Measure performance improvements

**Expected Results**:
- Tuned configuration improves throughput
- Proper timeouts prevent unnecessary rebalancing
- Fetch size affects memory usage

### Experiment 3: Consumer Lag Monitoring

**Goal**: Implement consumer lag monitoring

**Steps**:
1. Create consumer with lag monitoring
2. Send messages faster than consumer can process
3. Monitor lag increase
4. Implement lag-based alerting

**Expected Results**:
- Lag increases when producer is faster than consumer
- Monitoring detects lag spikes
- Alerting works for high lag scenarios

## 📊 Performance Monitoring

### **Key Consumer Metrics:**

#### **Throughput Metrics:**
- **Messages per second**: Rate of message consumption
- **Bytes per second**: Rate of data consumption
- **Fetch requests per second**: Rate of fetch requests

#### **Latency Metrics:**
- **Fetch latency**: Time to fetch messages from broker
- **Processing latency**: Time to process each message
- **Commit latency**: Time to commit offsets

#### **Lag Metrics:**
- **Consumer lag**: Messages behind latest offset
- **Lag trend**: Whether lag is increasing or decreasing
- **Partition lag**: Lag per partition

#### **Error Metrics:**
- **Error rate**: Percentage of failed operations
- **Retry rate**: Percentage of operations that need retry
- **Timeout rate**: Percentage of operations that timeout

### **Monitoring Implementation:**
```go
type ConsumerMetrics struct {
    MessagesConsumed    int64
    BytesConsumed       int64
    FetchLatency        time.Duration
    ProcessingLatency   time.Duration
    CommitLatency       time.Duration
    ConsumerLag         int64
    ErrorCount          int64
    RetryCount          int64
}

func (m *ConsumerMetrics) RecordMessage(latency time.Duration, bytes int) {
    atomic.AddInt64(&m.MessagesConsumed, 1)
    atomic.AddInt64(&m.BytesConsumed, int64(bytes))
    // Update other metrics
}
```

## 🎯 Key Takeaways

1. **Offset management** is crucial for data processing guarantees
2. **Consumer group tuning** affects performance and fault tolerance
3. **Fetch configuration** balances throughput and latency
4. **Error handling** ensures robust consumer applications
5. **Lag monitoring** is essential for production systems
6. **Performance tuning** requires balancing multiple factors

## 📝 Lesson 7 Assessment Questions

1. **What are the different offset initialization strategies?**
2. **When should you use manual vs automatic offset commits?**
3. **How do you choose the right rebalancing strategy?**
4. **What factors affect consumer performance?**
5. **How do you implement consumer lag monitoring?**
6. **What are the key metrics to monitor for consumers?**
7. **How do you handle retriable vs non-retriable errors?**
8. **What is the circuit breaker pattern and when to use it?**

---

## 🔄 Next Lesson Preview: Performance Tuning

**What we'll learn**:
- Throughput optimization strategies
- Latency reduction techniques
- Resource utilization optimization
- Scaling strategies
- Performance testing and benchmarking

**Hands-on experiments**:
- Performance testing with different configurations
- Throughput vs latency optimization
- Resource usage analysis
- Scaling experiments
