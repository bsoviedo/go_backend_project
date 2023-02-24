package server

import (
	"github.com/gofiber/fiber/v2"
)

func app() *fiber.App {

	App := fiber.New()

	App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! from golang")
	})

	return App
}
