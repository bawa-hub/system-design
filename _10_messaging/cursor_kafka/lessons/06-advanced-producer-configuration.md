# Lesson 6: Advanced Producer Configuration

## 🎯 Learning Objectives
By the end of this lesson, you will understand:
- Advanced producer configuration options
- Batching and compression strategies
- Acknowledgment modes and durability guarantees
- Idempotence and exactly-once semantics
- Performance tuning for producers
- Error handling and retry mechanisms

## 📚 Theory Section

### 1. **Producer Configuration Deep Dive**

#### **Core Configuration Options**

**Bootstrap Servers:**
```go
config := sarama.NewConfig()
config.Net.SASL.Enable = false
config.Net.TLS.Enable = false
config.Net.DialTimeout = 30 * time.Second
config.Net.ReadTimeout = 30 * time.Second
config.Net.WriteTimeout = 30 * time.Second
```

**Producer-Specific Settings:**
```go
config.Producer.RequiredAcks = sarama.WaitForAll
config.Producer.Retry.Max = 3
config.Producer.Return.Successes = true
config.Producer.Return.Errors = true
config.Producer.Timeout = 10 * time.Second
```

### 2. **Batching and Compression**

#### **Batching Configuration:**
```go
config.Producer.Flush.Frequency = 100 * time.Millisecond
config.Producer.Flush.Bytes = 1024 * 1024 // 1MB
config.Producer.Flush.Messages = 1000
```

**Benefits of Batching:**
- **Higher throughput**: Send multiple messages in one request
- **Reduced network overhead**: Fewer network round trips
- **Better compression**: More data to compress together

**Trade-offs:**
- **Increased latency**: Messages wait for batch to fill
- **Memory usage**: Messages stored in memory until sent
- **Risk of data loss**: Messages in memory lost if producer fails

#### **Compression Configuration:**
```go
config.Producer.Compression = sarama.CompressionSnappy
// Options: None, GZIP, Snappy, LZ4, ZSTD
```

**Compression Comparison:**
| Algorithm | CPU Usage | Compression Ratio | Speed |
|-----------|-----------|-------------------|-------|
| None | Low | 1:1 | Fastest |
| GZIP | High | 3:1 | Slowest |
| Snappy | Medium | 2:1 | Fast |
| LZ4 | Low | 2.5:1 | Very Fast |
| ZSTD | Medium | 3:1 | Medium |

### 3. **Acknowledgment Modes**

#### **Acks = 0 (Fire and Forget)**
```go
config.Producer.RequiredAcks = sarama.NoResponse
```
- **Durability**: No guarantee (messages can be lost)
- **Performance**: Highest throughput, lowest latency
- **Use case**: Logs, metrics, non-critical data

#### **Acks = 1 (Leader Acknowledgment)**
```go
config.Producer.RequiredAcks = sarama.WaitForLocal
```
- **Durability**: Leader confirms receipt (can lose data if leader fails)
- **Performance**: Good balance
- **Use case**: Most applications

#### **Acks = All (ISR Acknowledgment)**
```go
config.Producer.RequiredAcks = sarama.WaitForAll
```
- **Durability**: All ISR replicas confirm receipt (highest durability)
- **Performance**: Lower throughput, higher latency
- **Use case**: Critical data requiring high durability

### 4. **Idempotence and Exactly-Once Semantics**

#### **Idempotent Producer:**
```go
config.Producer.Idempotent = true
config.Producer.RequiredAcks = sarama.WaitForAll
config.Producer.Retry.Max = 3
config.Producer.Retry.RetryBackoff = 100 * time.Millisecond
```

**Benefits:**
- **Duplicate prevention**: Same message sent twice won't be duplicated
- **Retry safety**: Safe to retry failed messages
- **Exactly-once semantics**: Achieve exactly-once delivery

**Requirements:**
- **Acks = All**: Must use WaitForAll
- **Retry enabled**: Must enable retries
- **Max in-flight = 1**: Only one message in flight per partition

### 5. **Performance Tuning Strategies**

#### **Throughput Optimization:**
```go
// Increase batch size
config.Producer.Flush.Bytes = 1024 * 1024 // 1MB
config.Producer.Flush.Messages = 1000

// Increase compression
config.Producer.Compression = sarama.CompressionSnappy

// Use async producer
producer, err := sarama.NewAsyncProducer(brokers, config)
```

#### **Latency Optimization:**
```go
// Reduce batch size
config.Producer.Flush.Bytes = 1024 // 1KB
config.Producer.Flush.Messages = 10

// Reduce flush frequency
config.Producer.Flush.Frequency = 10 * time.Millisecond

// Use sync producer for immediate feedback
producer, err := sarama.NewSyncProducer(brokers, config)
```

#### **Memory Optimization:**
```go
// Limit batch size
config.Producer.Flush.Bytes = 512 * 1024 // 512KB
config.Producer.Flush.Messages = 500

// Set buffer size
config.Producer.Flush.BufferSize = 256
```

### 6. **Error Handling and Retry Mechanisms**

#### **Retry Configuration:**
```go
config.Producer.Retry.Max = 3
config.Producer.Retry.RetryBackoff = 100 * time.Millisecond
config.Producer.Retry.MaxRetryTime = 5 * time.Second
```

#### **Error Handling Patterns:**
```go
// Sync Producer Error Handling
partition, offset, err := producer.SendMessage(message)
if err != nil {
    if retriableError(err) {
        // Retry logic
        retryMessage(message)
    } else {
        // Log and handle non-retriable error
        log.Printf("Non-retriable error: %v", err)
    }
}

// Async Producer Error Handling
go func() {
    for err := range producer.Errors() {
        if retriableError(err.Err) {
            // Retry logic
            retryMessage(err.Msg)
        } else {
            // Handle non-retriable error
            log.Printf("Non-retriable error: %v", err.Err)
        }
    }
}()
```

## 🧪 Hands-on Experiments

### Experiment 1: Batching and Compression

**Goal**: Understand the impact of batching and compression

**Steps**:
1. Create producer with no batching, no compression
2. Create producer with batching, no compression
3. Create producer with batching and compression
4. Measure throughput and latency for each

**Expected Results**:
- Batching increases throughput but increases latency
- Compression reduces network usage but increases CPU usage
- Optimal configuration depends on use case

### Experiment 2: Acknowledgment Modes

**Goal**: Understand different durability guarantees

**Steps**:
1. Create producer with acks=0
2. Create producer with acks=1
3. Create producer with acks=all
4. Test durability by stopping broker during send

**Expected Results**:
- acks=0: Fastest but can lose data
- acks=1: Good balance, can lose data if leader fails
- acks=all: Slowest but highest durability

### Experiment 3: Idempotent Producer

**Goal**: Understand exactly-once semantics

**Steps**:
1. Create idempotent producer
2. Send same message multiple times
3. Verify no duplicates in consumer
4. Test retry behavior

**Expected Results**:
- Duplicate messages are not stored
- Retries are safe and don't create duplicates
- Exactly-once semantics achieved

## 📊 Performance Monitoring

### **Key Metrics to Monitor:**

#### **Throughput Metrics:**
- **Messages per second**: Rate of message production
- **Bytes per second**: Rate of data production
- **Batch size**: Average messages per batch

#### **Latency Metrics:**
- **Send latency**: Time from send to acknowledgment
- **Batch latency**: Time messages wait in batch
- **Network latency**: Time for network round trip

#### **Error Metrics:**
- **Retry rate**: Percentage of messages that need retry
- **Error rate**: Percentage of messages that fail
- **Timeout rate**: Percentage of messages that timeout

### **Monitoring Implementation:**
```go
// Custom metrics collection
type ProducerMetrics struct {
    MessagesSent    int64
    BytesSent       int64
    SendLatency     time.Duration
    RetryCount      int64
    ErrorCount      int64
}

func (p *ProducerMetrics) RecordSend(latency time.Duration, bytes int) {
    atomic.AddInt64(&p.MessagesSent, 1)
    atomic.AddInt64(&p.BytesSent, int64(bytes))
    // Update latency, retry, error counts
}
```

## 🎯 Key Takeaways

1. **Batching improves throughput** but increases latency
2. **Compression reduces network usage** but increases CPU usage
3. **Acknowledgment modes** trade off between performance and durability
4. **Idempotence enables exactly-once semantics** but requires specific configuration
5. **Performance tuning** requires balancing throughput, latency, and durability
6. **Error handling** is crucial for production systems

## 📝 Lesson 6 Assessment Questions

1. **What are the trade-offs between different acknowledgment modes?**
2. **How does batching affect throughput and latency?**
3. **What is idempotence and why is it important?**
4. **How do you choose the right compression algorithm?**
5. **What are the key metrics to monitor for producers?**
6. **How do you handle retriable vs non-retriable errors?**
7. **What configuration changes improve throughput?**
8. **What configuration changes improve latency?**

---

## 🔄 Next Lesson Preview: Advanced Consumer Configuration

**What we'll learn**:
- Consumer configuration options
- Offset management strategies
- Consumer group tuning
- Error handling patterns
- Performance optimization

**Hands-on experiments**:
- Consumer configuration tuning
- Offset management strategies
- Consumer group performance testing
- Error handling implementation
