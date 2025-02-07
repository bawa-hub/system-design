# Consumer Example with Partitions:
# Consumers can read messages from specific partitions or from all partitions of a topic. Here’s an example where we consume messages from my-topic:

from confluent_kafka import Consumer, KafkaError, KafkaException

# Consumer configuration
conf = {
    'bootstrap.servers': 'localhost:9092',
    'group.id': 'python-consumer-group',
    'auto.offset.reset': 'earliest'
}

# Create Consumer instance
consumer = Consumer(conf)

# Subscribe to the topic
consumer.subscribe(['my-topic'])

try:
    while True:
        msg = consumer.poll(timeout=1.0)  # Poll for new messages

        if msg is None:
            continue
        if msg.error():
            if msg.error().code() == KafkaError._PARTITION_EOF:
                print(f"End of partition reached: {msg.topic()} [{msg.partition()}] at offset {msg.offset()}")
            else:
                raise KafkaException(msg.error())
        else:
            # Print consumed message
            print(f"Consumed message: {msg.value().decode('utf-8')} from partition {msg.partition()}")

except KeyboardInterrupt:
    print("Consuming interrupted")
finally:
    consumer.close()


# Explanation:
#     The consumer subscribes to my-topic and consumes messages from all partitions. It will print the partition number for each message consumed.
#     If you have multiple consumers in a group, Kafka will ensure that each partition is consumed by only one consumer in the group.