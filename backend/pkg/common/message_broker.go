package common

import (
	"os"
	"github.com/segmentio/kafka-go"
	// Uncomment for RabbitMQ: "github.com/rabbitmq/amqp091-go"
)

func NewKafkaWriter(topic string) *kafka.Writer {
	brokers := os.Getenv("KAFKA_BROKERS")
	return &kafka.Writer{
		Addr:     kafka.TCP(brokers),
		Topic:    topic,
	}
}

// For RabbitMQ, implement similar connection logic as needed.
