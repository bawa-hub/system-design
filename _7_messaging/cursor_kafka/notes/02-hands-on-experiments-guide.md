# 🧪 Kafka Hands-on Experiments Guide

## 🎯 Complete Hands-on Learning Path

This guide contains all the hands-on experiments you've completed and can reference for practice.

---

## 🚀 Environment Setup

### Prerequisites
- Docker and Docker Compose installed
- Go 1.21+ installed
- Terminal access

### Start Kafka Environment
```bash
# Navigate to project directory
cd /Users/vikramkumar/CS/_4_backend/_7_messaging/cursor_kafka

# Start Kafka cluster
docker-compose up -d

# Verify Kafka is running
docker exec kafka kafka-topics --list --bootstrap-server localhost:9092
```

### Access Kafka UI
Open http://localhost:8080 in your browser to see:
- Topics and partitions
- Consumer groups
- Message flow in real-time
- Broker status

---

## 📚 Experiment 1: Basic Producer-Consumer

### Objective
Learn the fundamental message flow in Kafka.

### Files Created
- `examples/basic-producer/main.go`
- `examples/basic-consumer/main.go`

### Steps
1. **Create topic:**
   ```bash
   docker exec kafka kafka-topics --create --topic user-events --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
   ```

2. **Run producer:**
   ```bash
   go run examples/basic-producer/main.go
   ```

3. **Run consumer (in another terminal):**
   ```bash
   go run examples/basic-consumer/main.go
   ```

### What You Learned
- Basic message flow: Producer → Topic → Consumer
- Message structure: Key, Value, Timestamp, Offset
- Topic creation and management
- Consumer group basics

### Key Observations
- Messages are sent with keys and values
- Each message gets a unique offset
- Consumer processes messages in real-time
- Messages are stored persistently

---

## 📚 Experiment 2: Multi-Partition Topics

### Objective
Understand how partitioning works and enables parallel processing.

### Files Created
- `examples/simple-partition-demo/main.go`
- `examples/simple-consumer/main.go`

### Steps
1. **Create topic with multiple partitions:**
   ```bash
   docker exec kafka kafka-topics --create --topic simple-demo --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1
   ```

2. **Run partitioning demo:**
   ```bash
   go run examples/simple-partition-demo/main.go
   ```

3. **Run consumer to see messages:**
   ```bash
   go run examples/simple-consumer/main.go
   ```

### What You Learned
- **Key-based partitioning**: Same key always goes to same partition
- **Round-robin partitioning**: No key uses round-robin distribution
- **Partition assignment**: Hash function determines partition
- **Ordering guarantees**: Messages with same key stay in order

### Key Observations
- Different keys go to different partitions
- Same key always goes to same partition
- Partition assignment is deterministic
- Enables parallel processing

---

## 📚 Experiment 3: Consumer Groups

### Objective
Learn how consumer groups distribute work and provide load balancing.

### Files Created
- `examples/consumer-group-demo/main.go`
- `examples/consumer-group-producer/main.go`

### Steps
1. **Run producer to send messages:**
   ```bash
   go run examples/consumer-group-producer/main.go
   ```

2. **Start multiple consumers in same group:**
   ```bash
   # Terminal 1
   go run examples/consumer-group-demo/main.go consumer-1
   
   # Terminal 2
   go run examples/consumer-group-demo/main.go consumer-2
   
   # Terminal 3
   go run examples/consumer-group-demo/main.go consumer-3
   ```

3. **Observe partition assignment and load balancing**

### What You Learned
- **Consumer group behavior**: One consumer per partition
- **Load balancing**: Work distributed across consumers
- **Rebalancing**: Automatic reassignment when consumers join/leave
- **Parallel processing**: Multiple consumers work simultaneously

### Key Observations
- Each consumer gets assigned specific partitions
- Work is distributed evenly across consumers
- Adding/removing consumers triggers rebalancing
- Messages are processed in parallel

---

## 📚 Experiment 4: Replication and Fault Tolerance

### Objective
Understand how replication provides fault tolerance.

### Files Created
- `examples/replication-demo/main.go`
- `examples/replication-consumer/main.go`

### Steps
1. **Create topic with replication:**
   ```bash
   docker exec kafka kafka-topics --create --topic replicated-demo --bootstrap-server localhost:9092 --partitions 2 --replication-factor 1
   ```

2. **Run replication demo:**
   ```bash
   go run examples/replication-demo/main.go
   ```

3. **Run consumer to see messages:**
   ```bash
   go run examples/replication-consumer/main.go
   ```

4. **Check topic details:**
   ```bash
   docker exec kafka kafka-topics --describe --topic replicated-demo --bootstrap-server localhost:9092
   ```

### What You Learned
- **Replication factor**: Number of copies of data
- **Leader and followers**: How replicas work
- **ISR (In-Sync Replicas)**: Up-to-date replicas
- **Fault tolerance levels**: Different replication factors

### Key Observations
- Replication Factor = 1: No fault tolerance
- Leader handles all requests
- ISR ensures data consistency
- Higher replication = better fault tolerance

---

## 🔧 Troubleshooting Guide

### Common Issues and Solutions

#### 1. **Consumer Not Receiving Messages**
**Problem**: Consumer starts but doesn't receive messages
**Solutions**:
- Check if topic exists: `docker exec kafka kafka-topics --list --bootstrap-server localhost:9092`
- Verify consumer group: Check consumer group assignment
- Check offset: Consumer might be reading from latest offset

#### 2. **All Messages Go to Same Partition**
**Problem**: Messages with different keys go to same partition
**Solutions**:
- Check topic partition count: `docker exec kafka kafka-topics --describe --topic <topic> --bootstrap-server localhost:9092`
- Recreate topic with more partitions
- Verify key-based partitioning is working

#### 3. **Consumer Group Rebalancing Issues**
**Problem**: Frequent rebalancing or consumers not getting partitions
**Solutions**:
- Check consumer group status: `docker exec kafka kafka-consumer-groups --bootstrap-server localhost:9092 --describe --group <group>`
- Ensure stable consumer connections
- Check for consumer failures

#### 4. **Connection Issues**
**Problem**: Cannot connect to Kafka
**Solutions**:
- Check if Kafka is running: `docker ps`
- Verify port 9092 is available
- Check Docker logs: `docker logs kafka`

---

## 📊 Performance Testing

### Basic Performance Test

#### 1. **Producer Performance**
```bash
# Send 1000 messages
go run examples/consumer-group-producer/main.go
```

#### 2. **Consumer Performance**
```bash
# Monitor consumer lag
docker exec kafka kafka-consumer-groups --bootstrap-server localhost:9092 --describe --group simple-group --members --verbose
```

#### 3. **Topic Performance**
```bash
# Check topic details
docker exec kafka kafka-topics --describe --topic simple-demo --bootstrap-server localhost:9092
```

### Metrics to Monitor
- **Throughput**: Messages per second
- **Latency**: Time from producer to consumer
- **Consumer Lag**: How far behind consumers are
- **Partition Distribution**: Even distribution across partitions

---

## 🎯 Practice Exercises

### Exercise 1: Custom Partitioning
**Objective**: Create a producer that sends messages to specific partitions
**Requirements**:
- Send messages with different keys
- Verify key-based partitioning
- Test with different partition counts

### Exercise 2: Consumer Group Scaling
**Objective**: Test consumer group behavior with different consumer counts
**Requirements**:
- Start with 1 consumer, then add more
- Observe partition reassignment
- Test consumer failure scenarios

### Exercise 3: Replication Testing
**Objective**: Test different replication factors
**Requirements**:
- Create topics with RF=1, RF=3
- Compare fault tolerance
- Test leader election

### Exercise 4: Performance Optimization
**Objective**: Optimize Kafka performance
**Requirements**:
- Measure baseline performance
- Tune producer settings
- Tune consumer settings
- Compare results

---

## 📝 Experiment Log Template

Use this template to document your experiments:

```markdown
## Experiment: [Name]
**Date**: [Date]
**Objective**: [What you want to learn]

### Setup
- Topic: [Topic name]
- Partitions: [Number]
- Replication Factor: [Number]
- Consumer Group: [Group name]

### Steps
1. [Step 1]
2. [Step 2]
3. [Step 3]

### Observations
- [What you observed]
- [Key findings]
- [Unexpected behavior]

### Learnings
- [What you learned]
- [Key concepts]
- [Best practices]

### Questions
- [Questions that arose]
- [Areas for further study]
```

---

## 🚀 Next Steps

After completing these experiments, you're ready for:
1. **Advanced Configuration**: Tuning producers and consumers
2. **Performance Optimization**: Maximizing throughput and minimizing latency
3. **Error Handling**: Implementing robust error handling
4. **Monitoring**: Setting up comprehensive monitoring
5. **Real-world Applications**: Building production-ready systems

**Keep experimenting and learning!** 🎉
