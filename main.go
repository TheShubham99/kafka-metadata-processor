package main

import (
	"atlan-lily/consumer"
	"atlan-lily/producer"
	"atlan-lily/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go consumer.Consume("metadata_topic")

	go server.StartServer()

	// Produce a test message to Kafka - Outbound
	if err := producer.Produce("metadata_topic", map[string]interface{}{"12345": "This is an outbound event."}); err != nil {
		log.Fatalf("Failed to produce message: %v", err)
	}

	// Wait for exit signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
