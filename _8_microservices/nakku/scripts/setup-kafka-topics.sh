#!/bin/bash

# Setup Kafka topics for Nakku Quick Commerce

KAFKA_HOST=${KAFKA_HOST:-localhost:9092}

echo "Setting up Kafka topics for Nakku Quick Commerce..."

# Create topics
topics=(
  "user-events"
  "product-events"
  "inventory-events"
  "order-events"
  "cart-events"
  "payment-events"
  "delivery-events"
  "notification-events"
  "location-events"
  "analytics-events"
)

for topic in "${topics[@]}"; do
  echo "Creating topic: $topic"
  docker exec nakku-kafka kafka-topics.sh --create \
    --bootstrap-server localhost:9092 \
    --topic $topic \
    --partitions 3 \
    --replication-factor 1 \
    --if-not-exists
done

echo "Kafka topics setup completed!"

# List all topics
echo "Listing all topics:"
docker exec nakku-kafka kafka-topics.sh --list --bootstrap-server localhost:9092
