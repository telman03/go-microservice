package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/telman03/go-microservices/order-service/database"
	"github.com/telman03/go-microservices/order-service/kafka"
	"github.com/telman03/go-microservices/order-service/models"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/orders", CreateOrder)
}

func CreateOrder(c *fiber.Ctx) error {
	order := new(models.Order)
	if err := c.BodyParser(order); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	order.Status = "pending"
	database.DB.Create(&order)

	// Publish event to Kafka
	kafka.PublishOrderEvent(order)

	return c.JSON(order)
}