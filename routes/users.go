package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber/handlers"
)

func UserGroup(group fiber.Router) {
	group.Get("/", handlers.HandlerGetUser)
	group.Post("/", handlers.HandlerPostUser)
}
