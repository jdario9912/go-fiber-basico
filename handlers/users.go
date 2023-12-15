package handlers

import (
	"go-fiber/types"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func HandlerGetUser(c *fiber.Ctx) error {
	user := types.User{
		Firstname: "Joe",
		Lastname:  "Doe",
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func HandlerPostUser(c *fiber.Ctx) error {
	user := types.User{}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	user.Id = uuid.New().String()

	return c.Status(fiber.StatusOK).JSON(user)
}