package app

import (
	"go-fiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func App() {
	app := fiber.New()

	homeGroup := app.Group("/")
	userGroup := app.Group("/users")

	// middleware
	app.Use(logger.New())

	routes.HomeGroup(homeGroup)
	routes.UserGroup(userGroup)

	app.Listen(":3000")
}
