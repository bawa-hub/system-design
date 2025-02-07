# Here’s a basic consumer script that connects to Kafka, subscribes to a topic, and consumes messages:

from confluent_kafka import Consumer, KafkaException, KafkaError

# Consumer configuration
conf = {
    'bootstrap.servers': 'localhost:9092',  # Kafka broker address
    'group.id': 'python-consumer-group',  # Consumer group ID
    'auto.offset.reset': 'earliest'  # Start reading from the earliest message
}

# Create Consumer instance
consumer = Consumer(conf)

# Subscribe to a topic
consumer.subscribe(['test-topic'])

# Poll for messages (blocking)
try:
    while True:
        msg = consumer.poll(timeout=1.0)  # Wait for a message for 1 second

        if msg is None:
            # No message received within timeout
            continue
        if msg.error():
            # Handle message error
            if msg.error().code() == KafkaError._PARTITION_EOF:
                # End of partition
                print(f"End of partition reached: {msg.topic()} [{msg.partition()}] at offset {msg.offset()}")
            else:
                raise KafkaException(msg.error())
        else:
            # Message received successfully
            print(f"Consumed message: {msg.value().decode('utf-8')}")

except KeyboardInterrupt:
    print("Consuming interrupted")
finally:
    # Close the consumer
    consumer.close()


"""
Explanation:
    Consumer Configuration: We set the bootstrap.servers to the Kafka broker's address, specify the group.id for the consumer group, and configure auto.offset.reset to earliest, which means the consumer will start from the earliest available message if no offset is committed.
    Subscribing to Topics: The consumer.subscribe() method subscribes to the test-topic. If you want to subscribe to multiple topics, pass a list of topic names.
    Polling for Messages: The consumer.poll() method is used to pull messages from Kafka. It takes a timeout argument in seconds.
    Message Handling: The consumer checks if there was an error with the message. If the message was successfully received, it prints the message's value.
    Graceful Shutdown: We handle a KeyboardInterrupt (Ctrl+C) to close the consumer properly.
"""