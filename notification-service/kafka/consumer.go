package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
	"github.com/telman03/go-microservices/notification-service/database"
	"github.com/telman03/go-microservices/notification-service/email"
	"github.com/telman03/go-microservices/notification-service/models"
)

func ConsumeMessages() {
	topic := os.Getenv("KAFKA_TOPIC")
	broker := os.Getenv("KAFKA_BROKER")
	groupID := os.Getenv("KAFKA_GROUP")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	fmt.Println("📩 Notification service is listening for messages...")

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Error reading Kafka message:", err)
		}

		// Parse message
		var user struct {
			ID    uint   `json:"ID"`
			Name  string `json:"Name"`
			Email string `json:"Email"`
		}

		if err := json.Unmarshal(msg.Value, &user); err != nil {
			log.Println("Error parsing Kafka message:", err)
			continue
		}

		// Create notification
		notification := models.Notification{
			UserID:  user.ID,
			Email:   user.Email,
			Message: fmt.Sprintf("Welcome %s! Your account has been created.", user.Name),
		}

		// Store notification in database
		database.DB.Create(&notification)
		fmt.Printf("✅ Stored notification for %s\n", user.Email)

		// Send Email
		email.SendEmail(user.Email, "Welcome to our service!", notification.Message)
	}
}