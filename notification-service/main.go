package main

import (
	"fmt"
	"github.com/telman03/go-microservices/notification-service/database"

	"github.com/telman03/go-microservices/notification-service/config"
	"github.com/telman03/go-microservices/notification-service/kafka"

)

func main() {
	// Load environment variables
	config.LoadEnv()
	database.ConnectDB()

	
	fmt.Println("📡 Notification service is starting...")

	// Start Kafka message consumption
	kafka.ConsumeMessages()
}