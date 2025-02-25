package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/telman03/go-microservices/order-service/database"
	"github.com/telman03/go-microservices/order-service/models"
	"github.com/telman03/go-microservices/order-service/routes"
)

func main() {
	database.ConnectDB()

	// Run migrations
	database.DB.AutoMigrate(&models.Order{})

	app := fiber.New()
	routes.SetupRoutes(app)

	app.Listen(":8082")
}