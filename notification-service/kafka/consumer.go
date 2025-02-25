package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"net/http"

	"github.com/segmentio/kafka-go"
	"github.com/telman03/go-microservices/notification-service/database"
	"github.com/telman03/go-microservices/notification-service/email"
	"github.com/telman03/go-microservices/notification-service/models"
)


func fetchUserEmail(userID uint) (string, error) {
	userServiceURL := fmt.Sprintf("http://localhost:8080/users/%d", userID)

	resp, err := http.Get(userServiceURL)
	if err != nil {
		return "", fmt.Errorf("failed to reach user-service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("user-service returned status: %d", resp.StatusCode)
	}

	var user struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return "", fmt.Errorf("failed to decode user-service response: %w", err)
	}

	return user.Email, nil
}
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
			log.Fatal("❌ Error reading Kafka message:", err)
		}

		// Parse message
		var order struct {
			UserID uint    `json:"user_id"`
			Amount float64 `json:"amount"`
		}
		log.Println("📥 Received Kafka message:", string(msg.Value))

		if err := json.Unmarshal(msg.Value, &order); err != nil {
			log.Println("❌ Error parsing Kafka message:", err)
			continue
		}

		// ✅ Fetch user email from user-service API
		userEmail, err := fetchUserEmail(order.UserID)
		if err != nil {
			log.Println("❌ Failed to fetch user email:", err)
			continue
		}

		// Create notification
		notification := models.Notification{
			UserID:  order.UserID,
			Email:   userEmail,
			Message: fmt.Sprintf("Your order of $%.2f has been placed successfully!", order.Amount),
		}

		// Store notification in database
		database.DB.Create(&notification)
		fmt.Printf("✅ Stored notification for %s\n", userEmail)

		// Send Email
		if userEmail == "" {
			log.Println("❌ Skipping email: No recipient address provided")
			continue
		}

		email.SendEmail(userEmail, "Order Confirmation", notification.Message)
	}
}