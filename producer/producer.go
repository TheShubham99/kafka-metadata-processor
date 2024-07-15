package producer

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func Produce(topic string, message map[string]interface{}) error {
	// Convert the map to a JSON string
	msgBytes, err := json.Marshal(message)
	if err != nil {
		log.Printf("Failed to marshal message to JSON: %v", err)
		return fmt.Errorf("failed to marshal message to JSON: %v", err)
	}

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	brokers := []string{"kafka:9092"}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Printf("Failed to start Kafka producer: %v", err)
		return fmt.Errorf("failed to start Kafka producer: %v", err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msgBytes),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send message to Kafka: %v", err)
		return fmt.Errorf("failed to send message to Kafka: %v", err)
	}

	log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	return nil
}
