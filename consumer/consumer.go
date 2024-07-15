package consumer

import (
	"log"

	"github.com/Shopify/sarama"
)

func Consume(topic string) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	// reffered to official docs of the sarama library,

	brokers := []string{"kafka:9092"}
	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Printf("Failed to start Kafka consumer: %v", err)
		return
	}
	defer master.Close()

	consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Printf("Failed to start Kafka partition consumer: %v", err)
		return
	}
	defer consumer.Close()

	for {
		select {
		case msg := <-consumer.Messages():
			log.Printf("Consumed message: %s\n", string(msg.Value))
		case err := <-consumer.Errors():
			log.Printf("Error consuming message: %v\n", err)
		}
	}
}
