# 🚀 Kafka Learning Journey: From Zero to Expert

Welcome to your comprehensive Kafka learning experience! This repository contains everything you need to master Apache Kafka from absolute beginner to expert level.

## 🎯 Learning Path Overview

### Phase 1: Foundation (Weeks 1-2)
- [x] Environment Setup
- [ ] Core Concepts
- [ ] Basic Producer/Consumer
- [ ] Topics and Partitions
- [ ] Consumer Groups

### Phase 2: Intermediate (Weeks 3-4)
- [ ] Advanced Configuration
- [ ] Performance Tuning
- [ ] Error Handling
- [ ] Monitoring

### Phase 3: Advanced (Weeks 5-6)
- [ ] Stream Processing
- [ ] Schema Registry
- [ ] Security
- [ ] Multi-cluster

### Phase 4: Expert (Weeks 7-8)
- [ ] Custom Connectors
- [ ] Performance Optimization
- [ ] Troubleshooting
- [ ] Real-world Architectures

## 🛠️ Quick Start

1. **Start Kafka Environment:**
   ```bash
   docker-compose up -d
   ```

2. **Access Kafka UI:**
   Open http://localhost:8080 in your browser

3. **Run Examples:**
   ```bash
   # Go examples
   go run examples/basic-producer/main.go
   go run examples/basic-consumer/main.go
   ```

## 📚 Learning Materials

- `examples/` - Hands-on code examples
- `lessons/` - Detailed lesson content
- `exercises/` - Practice exercises
- `projects/` - Real-world projects

## 🎓 Current Lesson: Understanding Kafka Fundamentals

### What is Kafka?
Apache Kafka is a distributed streaming platform that can:
- **Publish and subscribe** to streams of records
- **Store** streams of records in a fault-tolerant way
- **Process** streams of records as they occur

### Key Concepts:
1. **Producer** - Applications that send data
2. **Consumer** - Applications that read data
3. **Topic** - Categories for organizing messages
4. **Partition** - Topics are split into partitions for parallel processing
5. **Broker** - Kafka servers that store and serve data
6. **Consumer Group** - Multiple consumers working together

Let's start with our first hands-on example!
