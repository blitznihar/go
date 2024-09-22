package queue

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaProducerConnection func for connection to Kafka producer.
func KafkaProducerConnection() (*kafka.Producer, error) {
	// Define database connection settings.

	opts := &kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BROKER"),
		"client.id":         hostname(),

		"acks": "all",
	}
	topic := os.Getenv("KAFKA_PRODUCERTOPIC")
	if topic == "" {
		return nil, fmt.Errorf("KAFKA_PRODUCERTOPIC is not set")
	}
	clientID := os.Getenv("KAFKA_CLIENTID")
	if clientID == "" {
		return nil, fmt.Errorf("KAFKA_CLIENTID is not set")
	}
	// Define database connection for Mysql.
	p, err := kafka.NewProducer(opts)

	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}
	defer p.Close()
	return p, nil
}

func KafkaConsumerConnection() (*kafka.Consumer, error) {
	// Define database connection settings.

	opts := &kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BROKER"),
		"group.id":          hostname(),
		"auto.offset.reset": "smallest",
	}

	// Define database connection for Mysql.
	c, err := kafka.NewConsumer(opts)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to kafka, %w", err)
	}

	return c, nil
}

// Use a helper function to get hostname
func hostname() string {
	hostname, _ := os.Hostname()
	return hostname
}
