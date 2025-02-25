package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/telman03/go-microservices/user-service/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/users", handlers.CreateUser)
	app.Get("/users/:id", handlers.GetUserByID)
}