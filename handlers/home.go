package handlers

import "github.com/gofiber/fiber/v2"

func HandlerGetHome(c *fiber.Ctx) error {
	return c.SendString("hola Fiber")
}