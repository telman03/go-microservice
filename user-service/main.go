package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/telman03/go-microservices/user-service/config"
	"github.com/telman03/go-microservices/user-service/database"
	"github.com/telman03/go-microservices/user-service/kafka"
	"github.com/telman03/go-microservices/user-service/routes"
)

func main() {
	config.LoadEnv()
	database.ConnectDB()
	kafka.ConnectKafka()

	app := fiber.New()
	routes.SetupRoutes(app)

	app.Listen(":8080")
//	yswo yqwq tsyk kher
}