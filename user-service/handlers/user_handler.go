package handlers

import (
//	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/telman03/go-microservices/user-service/database"
//	"github.com/telman03/go-microservices/user-service/kafka"
	"github.com/telman03/go-microservices/user-service/models"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(201).JSON(user)
}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}