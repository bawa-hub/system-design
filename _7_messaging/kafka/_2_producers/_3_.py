"""
Key Producer Operations and Advanced Concepts

1. Asynchronous vs. Synchronous Producer:
        The produce() method is asynchronous by default. It returns immediately, allowing you to continue processing while Kafka handles sending the message in the background.
        You can make the producer synchronous by calling flush() after producing messages. This waits for the acknowledgment before moving on.

2. Message Key: Kafka uses the key to determine which partition a message should go to. If you specify the same key for multiple messages, Kafka ensures that they go to the same partition.
   producer.produce('test-topic', key='user_1', value='Hello User 1!')
   producer.produce('test-topic', key='user_2', value='Hello User 2!')

   Both messages will go to the same partition if they share the same key.

3. Handling Failures and Retries: Kafka provides configurations for retrying failed messages:

    retries: Number of retries on failure.
    retry.backoff.ms: Time to wait before retrying after a failure.
    delivery.timeout.ms: Maximum time to wait for message acknowledgment.
    
    Example:

conf = {
    'bootstrap.servers': 'localhost:9092',
    'acks': 'all',
    'retries': 3,
    'retry.backoff.ms': 500,
    'delivery.timeout.ms': 10000
}
   
   
4. Producer with Key and Value Serialization: Kafka producers can serialize data before sending it to Kafka. By default, Kafka sends byte arrays, but you can use specific serializers for structured data like JSON or Avro.

    Example with JSON serialization:   
        
"""
        
import json
from confluent_kafka import Producer

# Define the JSON serializer function
def json_serializer(obj):
    return json.dumps(obj).encode('utf-8')

# Kafka configuration (removed unsupported key/value serializer configs)
conf = {
    'bootstrap.servers': 'localhost:9092',
}

# Create the producer
producer = Producer(conf)

# Serialize the key and value manually
key = json_serializer({'user_id': 1})
value = json_serializer({'message': 'Hello JSON'})

# Produce the message
producer.produce('test-topic', key=key, value=value)
producer.flush()

     