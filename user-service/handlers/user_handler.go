package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/telman03/go-microservices/user-service/database"
	"github.com/telman03/go-microservices/user-service/kafka"
	"github.com/telman03/go-microservices/user-service/models"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	database.DB.Create(&user)

	// Publish event to Kafka
	message, _ := json.Marshal(user)
	kafka.ProduceMessage("user_created", string(message))

	return c.Status(201).JSON(user)
}