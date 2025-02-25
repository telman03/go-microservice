package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
	"github.com/telman03/go-microservices/order-service/models"
)

func PublishOrderEvent(order *models.Order) {
	topic := os.Getenv("KAFKA_TOPIC")
	broker := os.Getenv("KAFKA_BROKER")

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker},
		Topic:   topic,
	})

	message, err := json.Marshal(order)
	if err != nil {
		log.Println("❌ Failed to marshal order event:", err)
		return
	}

	err = writer.WriteMessages(context.Background(),
		kafka.Message{Value: message},
	)
	if err != nil {
		log.Println("❌ Failed to send Kafka message:", err)
	} else {
		fmt.Println("✅ Order event published to Kafka:", string(message))
	}
}