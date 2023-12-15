package routes

import (
	"go-fiber/handlers"

	"github.com/gofiber/fiber/v2"
)

func HomeGroup(group fiber.Router)  {
	group.Get("", handlers.HandlerGetHome)
}