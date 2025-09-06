package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

// KafkaProducer handles Kafka message publishing
type KafkaProducer struct {
	writer *kafka.Writer
	logger *logrus.Logger
}

// KafkaConsumer handles Kafka message consumption
type KafkaConsumer struct {
	reader *kafka.Reader
	logger *logrus.Logger
}

// KafkaMessage represents a Kafka message
type KafkaMessage struct {
	Topic     string      `json:"topic"`
	Key       string      `json:"key,omitempty"`
	Value     interface{} `json:"value"`
	Timestamp time.Time   `json:"timestamp"`
}

// NewKafkaProducer creates a new Kafka producer
func NewKafkaProducer(brokers []string, logger *logrus.Logger) *KafkaProducer {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireOne,
		Async:        false,
	}

	return &KafkaProducer{
		writer: writer,
		logger: logger,
	}
}

// PublishMessage publishes a message to Kafka
func (p *KafkaProducer) PublishMessage(ctx context.Context, topic string, key string, value interface{}) error {
	// Serialize the value to JSON
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	// Create Kafka message
	message := kafka.Message{
		Topic: topic,
		Key:   []byte(key),
		Value: jsonValue,
		Time:  time.Now(),
	}

	// Publish message
	if err := p.writer.WriteMessages(ctx, message); err != nil {
		p.logger.WithError(err).Error("Failed to publish message to Kafka")
		return fmt.Errorf("failed to publish message: %w", err)
	}

	p.logger.WithFields(logrus.Fields{
		"topic": topic,
		"key":   key,
	}).Info("Message published to Kafka")

	return nil
}

// Close closes the Kafka producer
func (p *KafkaProducer) Close() error {
	return p.writer.Close()
}

// NewKafkaConsumer creates a new Kafka consumer
func NewKafkaConsumer(brokers []string, topic string, groupID string, logger *logrus.Logger) *KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		Topic:          topic,
		GroupID:        groupID,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset:    kafka.LastOffset,
	})

	return &KafkaConsumer{
		reader: reader,
		logger: logger,
	}
}

// ConsumeMessages consumes messages from Kafka
func (c *KafkaConsumer) ConsumeMessages(ctx context.Context, handler func(KafkaMessage) error) error {
	for {
		select {
		case <-ctx.Done():
			c.logger.Info("Stopping Kafka consumer")
			return ctx.Err()
		default:
			// Read message
			message, err := c.reader.ReadMessage(ctx)
			if err != nil {
				c.logger.WithError(err).Error("Failed to read message from Kafka")
				continue
			}

			// Parse message
			kafkaMsg := KafkaMessage{
				Topic:     message.Topic,
				Key:       string(message.Key),
				Timestamp: message.Time,
			}

			// Unmarshal value
			if err := json.Unmarshal(message.Value, &kafkaMsg.Value); err != nil {
				c.logger.WithError(err).Error("Failed to unmarshal message value")
				continue
			}

			// Handle message
			if err := handler(kafkaMsg); err != nil {
				c.logger.WithError(err).Error("Failed to handle message")
				// Continue processing other messages even if one fails
				continue
			}

			c.logger.WithFields(logrus.Fields{
				"topic": message.Topic,
				"key":   string(message.Key),
			}).Info("Message consumed from Kafka")
		}
	}
}

// Close closes the Kafka consumer
func (c *KafkaConsumer) Close() error {
	return c.reader.Close()
}

// CreateTopic creates a Kafka topic if it doesn't exist
func CreateTopic(brokers []string, topic string, partitions int, replicationFactor int) error {
	conn, err := kafka.DialLeader(context.Background(), "tcp", brokers[0], topic, 0)
	if err != nil {
		return fmt.Errorf("failed to connect to Kafka: %w", err)
	}
	defer conn.Close()

	// Create topic configuration
	topicConfig := kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     partitions,
		ReplicationFactor: replicationFactor,
	}

	// Create topic
	if err := conn.CreateTopics(topicConfig); err != nil {
		return fmt.Errorf("failed to create topic: %w", err)
	}

	return nil
}

// ListTopics lists all Kafka topics
func ListTopics(brokers []string) ([]string, error) {
	conn, err := kafka.Dial("tcp", brokers[0])
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Kafka: %w", err)
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		return nil, fmt.Errorf("failed to read partitions: %w", err)
	}

	topicMap := make(map[string]bool)
	for _, partition := range partitions {
		topicMap[partition.Topic] = true
	}

	var topics []string
	for topic := range topicMap {
		topics = append(topics, topic)
	}

	return topics, nil
}
