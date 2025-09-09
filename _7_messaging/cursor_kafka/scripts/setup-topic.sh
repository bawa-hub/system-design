#!/bin/bash

echo "🔧 Setting up Kafka topic with multiple partitions..."

# Create topic with 3 partitions
docker exec kafka kafka-topics --create \
  --topic user-events-multi \
  --bootstrap-server localhost:9092 \
  --partitions 3 \
  --replication-factor 1 \
  --if-not-exists

echo "✅ Topic 'user-events-multi' created with 3 partitions"

# List topics to verify
echo "📋 Available topics:"
docker exec kafka kafka-topics --list --bootstrap-server localhost:9092

# Describe the topic to see partition details
echo "📊 Topic details:"
docker exec kafka kafka-topics --describe --topic user-events --bootstrap-server localhost:9092
