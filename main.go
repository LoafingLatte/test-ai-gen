package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// GET route for root path that returns "Hello world"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	// Listen on port 3000
	app.Listen(":3000")
}
