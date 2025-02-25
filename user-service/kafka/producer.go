package kafka

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

var KafkaWriter *kafka.Writer

func ConnectKafka() {
	KafkaWriter = &kafka.Writer{
		Addr:         kafka.TCP(os.Getenv("KAFKA_BROKER")),
		Topic:        "user_created",
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
		BatchTimeout: 10 * time.Millisecond,
	}

	fmt.Println("Kafka producer connected!")
}

func ProduceMessage(key, value string) {
	err := KafkaWriter.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(value),
		},
	)

	if err != nil {
		log.Println("Failed to send Kafka message:", err)
	} else {
		fmt.Println("Kafka message sent:", value)
	}
}