# Kafka Producer Example using Python

# First, ensure you have the necessary Python libraries for Kafka. 
# We will use the confluent-kafka library, which is a Python client for Kafka that provides a high-performance producer and consumer.

# Install the confluent-kafka library: To interact with Kafka from Python, you need to install the confluent-kafka library. Run the following command:
# pip install confluent-kafka

from confluent_kafka import Producer

# Callback to confirm message delivery
def delivery_report(err, msg):
    if err is not None:
        print(f"Message delivery failed: {err}")
    else:
        print(f"Message delivered to {msg.topic()} [{msg.partition()}]")

# Producer configuration
conf = {
    'bootstrap.servers': 'localhost:9092',  # Kafka broker address
    'client.id': 'python-producer'
}

# Create Producer instance
producer = Producer(conf)

# Produce a message
producer.produce('test-topic', key='key1', value='Hello Kafka Again!', callback=delivery_report)

# Wait for any outstanding messages to be delivered
producer.flush()


# Producer Configuration: We specify the Kafka broker's address using bootstrap.servers and set the client ID for the producer.
# produce(): This method is used to send a message to a Kafka topic. Here, we're sending the message Hello Kafka! to the test-topic with a key key1. The key helps Kafka determine which partition the message goes to.
# Callback: The delivery_report function is used as a callback to confirm if the message was successfully delivered or if there was an error.
# flush(): It ensures that all messages are sent before the program exits.

"""
Producer Configuration Options:
acks: Set the acknowledgment level.
'acks': 'all'  # Ensures the message is acknowledged by all replicas
batch.size: Control the size of batches of messages.
'batch.size': 16384  # Size of the batch in bytes
compression.type: Use compression for message batches to improve throughput.
'compression.type': 'gzip'  # You can also use 'snappy' or 'lz4'
"""
