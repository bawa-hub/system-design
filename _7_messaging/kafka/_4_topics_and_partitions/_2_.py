# Producer Example with Partitions:

# Creating Topics (via Kafka CLI)
# Use the kafka-topics.sh tool to create a topic with partitions and replication:
# kafka-topics.sh --create --topic my-topic --partitions 3 --replication-factor 2 --bootstrap-server localhost:9092

# A producer can send messages to different partitions of a topic, either by specifying a key or letting Kafka handle the partitioning automatically.

# Here’s a Python producer example that sends messages to my-topic:

from confluent_kafka import Producer

# Callback function to report message delivery
def delivery_report(err, msg):
    if err is not None:
        print(f"Message delivery failed: {err}")
    else:
        print(f"Message delivered to {msg.topic()} [{msg.partition()}]")

# Producer configuration
conf = {
    'bootstrap.servers': 'localhost:9092',
    'client.id': 'python-producer'
}

# Create Producer instance
producer = Producer(conf)

# Sending messages with keys (which will determine partition)
for i in range(10):
    key = f"user_{i % 3}"  # Example key, rotating between 3 users
    producer.produce('my-topic', key=key, value=f"Message {i}", callback=delivery_report)

# Wait for all messages to be sent
producer.flush()

# Explanation:
#     Messages are sent to my-topic, and the key is used to determine the partition.
#     Kafka will hash the key and use the result to decide which partition the message will be placed in. Since we have 3 keys (user_0, user_1, user_2), the messages will be distributed across 3 partitions.