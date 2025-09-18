# Lesson 2: Topics and Partitions Deep Dive

## 🎯 Learning Objectives
By the end of this lesson, you will understand:
- How Kafka partitioning strategy works
- The relationship between topics, partitions, and consumers
- Replication and fault tolerance mechanisms
- Consumer group rebalancing
- Offset management strategies
- Performance implications of partitioning

## 📚 Theory Section

### 1. **Understanding Partitions in Detail**

#### **What is a Partition?**
A partition is a **log file** that stores messages in **sequential order**. Think of it as a **numbered sequence** where each message gets a position (offset).

```
Partition 0: [msg1, msg2, msg3, msg4, msg5, ...]
Partition 1: [msg1, msg2, msg3, msg4, msg5, ...]
Partition 2: [msg1, msg2, msg3, msg4, msg5, ...]
```

#### **Key Properties of Partitions:**
- **Immutable**: Messages are appended, never modified
- **Ordered**: Messages within a partition are always in order
- **Durable**: Stored on disk, not just memory
- **Replicated**: Multiple copies for fault tolerance

### 2. **Partitioning Strategy**

#### **How Kafka Decides Which Partition to Use:**

**Method 1: Key-based Partitioning (Recommended)**
```
Key → Hash Function → Partition Number
```
- **Same key = Same partition** (guarantees ordering)
- **Different keys = Distributed across partitions**
- **Example**: `user-123` always goes to partition 1

**Method 2: Round-robin (No key provided)**
```
Message 1 → Partition 0
Message 2 → Partition 1
Message 3 → Partition 2
Message 4 → Partition 0
```

#### **Partitioning Rules:**
1. **Ordering**: Messages with same key are processed in order
2. **Parallelism**: Different keys can be processed in parallel
3. **Load Balancing**: Work is distributed across partitions

### 3. **Replication and Fault Tolerance**

#### **Replication Factor**
- **Replication Factor = 1**: No replication (single point of failure)
- **Replication Factor = 3**: 3 copies of each partition (recommended for production)

#### **Leader and Followers**
```
Partition 0: Leader (Broker 1) ← Followers (Broker 2, Broker 3)
Partition 1: Leader (Broker 2) ← Followers (Broker 1, Broker 3)
Partition 2: Leader (Broker 3) ← Followers (Broker 1, Broker 2)
```

**How it works:**
- **Leader**: Handles all read/write requests
- **Followers**: Replicate data from leader
- **ISR (In-Sync Replicas)**: Replicas that are up-to-date
- **Failover**: If leader fails, a follower becomes leader

### 4. **Consumer Groups and Partition Assignment**

#### **Consumer Group Behavior:**
- **One consumer per partition**: Each partition is assigned to only one consumer in the group
- **Load balancing**: Work is distributed across consumers
- **Fault tolerance**: If a consumer fails, its partitions are reassigned

#### **Example with 3 Partitions and 2 Consumers:**
```
Consumer 1: Partition 0, Partition 2
Consumer 2: Partition 1
```

#### **Rebalancing:**
When consumers join/leave a group:
1. **Stop**: All consumers stop processing
2. **Reassign**: Partitions are reassigned
3. **Resume**: Consumers resume processing new assignments

### 5. **Offset Management**

#### **What is an Offset?**
- **Offset**: Position of a message in a partition
- **Starts at 0**: First message has offset 0
- **Incremental**: Each new message gets next offset
- **Persistent**: Stored in Kafka, survives restarts

#### **Offset Commit Strategies:**
1. **Automatic**: Kafka commits offsets periodically
2. **Manual**: Application commits offsets after processing
3. **At-least-once**: Process message, then commit offset
4. **Exactly-once**: Use transactions (advanced topic)

## 🧪 Hands-on Experiments

### Experiment 1: Understanding Partitioning Strategy

**Goal**: See how different keys go to different partitions

**Steps**:
1. Create topic with 3 partitions
2. Send messages with different keys
3. Observe partition distribution
4. Send messages with same key
5. Observe they go to same partition

### Experiment 2: Consumer Group Behavior

**Goal**: Understand how consumer groups distribute work

**Steps**:
1. Start 2 consumers in same group
2. Send messages to 3-partition topic
3. Observe partition assignment
4. Add/remove consumers
5. Observe rebalancing

### Experiment 3: Offset Management

**Goal**: Understand how offsets work

**Steps**:
1. Send messages to topic
2. Start consumer and process messages
3. Stop consumer
4. Send more messages
5. Restart consumer
6. Observe where it resumes

## 📊 Performance Implications

### **Partition Count Considerations:**

**Too Few Partitions:**
- ❌ Limited parallelism
- ❌ Consumer bottlenecks
- ❌ Uneven load distribution

**Too Many Partitions:**
- ❌ More overhead
- ❌ More file handles
- ❌ Slower rebalancing

**Optimal Partition Count:**
- ✅ Match consumer count
- ✅ Consider future growth
- ✅ Balance parallelism vs overhead

### **Key Design Principles:**

1. **Partition Count = Consumer Count**: For optimal parallelism
2. **Key-based Partitioning**: For ordering guarantees
3. **Replication Factor = 3**: For fault tolerance
4. **Monitor Rebalancing**: Frequent rebalancing indicates issues

## 🎯 Key Takeaways

1. **Partitions enable parallelism** within a topic
2. **Keys determine partitioning** and ordering guarantees
3. **Consumer groups distribute work** across partitions
4. **Replication provides fault tolerance**
5. **Offset management** tracks processing progress
6. **Partition count** affects performance and scalability

## 📝 Lesson 2 Assessment Questions

1. **How does Kafka decide which partition to send a message to?**
2. **What happens if a consumer in a group fails?**
3. **Why is replication important in Kafka?**
4. **What is the relationship between partition count and consumer count?**
5. **How do offsets help with message processing?**
6. **What are the trade-offs of having too many partitions?**

---

## 🔄 Next Lesson Preview: Consumer Groups and Parallel Processing

**What we'll learn**:
- Consumer group coordination
- Rebalancing strategies
- Offset management best practices
- Error handling and retry mechanisms
- Performance tuning for consumers

**Hands-on experiments**:
- Consumer group rebalancing scenarios
- Offset management strategies
- Error handling patterns
- Performance testing with multiple consumers
