package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewRouter() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())

	app.Post("/upload", func(c *fiber.Ctx) error {
		log.Println(c.Body())
		return c.SendStatus(fiber.StatusOK)
	})

	return app
}
