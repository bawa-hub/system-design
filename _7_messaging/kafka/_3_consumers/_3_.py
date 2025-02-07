"""
Consumer Group and Offset Management:

    Consumer Group: If multiple consumers are part of a consumer group, Kafka distributes the topic partitions among the consumers in the group. Each consumer will process different partitions concurrently.

    Manual Offset Committing: If you want to manually commit offsets (to ensure messages are only marked as consumed when fully processed), you can disable automatic offset committing and commit offsets explicitly.
"""

consumer = Consumer({
    'bootstrap.servers': 'localhost:9092',
    'group.id': 'python-consumer-group',
    'auto.offset.reset': 'earliest',
    'enable.auto.commit': False  # Disable auto-commit of offsets
})

# Consume messages and commit offset manually
try:
    while True:
        msg = consumer.poll(timeout=1.0)
        if msg is None:
            continue
        if msg.error():
            if msg.error().code() == KafkaError._PARTITION_EOF:
                print(f"End of partition reached: {msg.topic()} [{msg.partition()}] at offset {msg.offset()}")
            else:
                raise KafkaException(msg.error())
        else:
            print(f"Consumed message: {msg.value().decode('utf-8')}")
            # Manually commit the offset
            consumer.commit(message=msg)

except KeyboardInterrupt:
    print("Consuming interrupted")
finally:
    consumer.close()

# In this case, the consumer will only commit the offset after successfully processing the message, ensuring that if there's a failure, the message will be re-consumed.

# Handling Consumer Failures: Kafka provides mechanisms for handling consumer failures, such as retries and dead-letter queues. When a consumer crashes, another consumer in the group can take over processing of the partitions.