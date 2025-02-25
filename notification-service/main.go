package main

import (
	"fmt"
	"log"

	"github.com/telman03/go-microservices/notification-service/config"
	"github.com/telman03/go-microservices/notification-service/kafka"
	"github.com/telman03/go-microservices/notification-service/email"

	"github.com/IBM/sarama"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Kafka consumer setup
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	brokers := []string{"localhost:9092"}
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatal("Error creating Kafka consumer:", err)
	}
	defer consumer.Close()

	// Subscribe to Kafka topic
	topic := "order_created"
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("Error creating partition consumer:", err)
	}
	defer partitionConsumer.Close()

	fmt.Println("📡 Listening for new orders...")

	// Listen for messages
	for msg := range partitionConsumer.Messages() {
		fmt.Println("🔔 New Order Received:", string(msg.Value))

		// Send email notification
		email.HandleMessage(msg.Value)
	}

	// Start Kafka message consumption
	kafka.ConsumeMessages()
}
