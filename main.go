package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/michaelomh/go-api-server/pkg/configs"
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	routes.NotFoundRoute(app)

	// Define a route for GET /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Start server on port 8080
	app.Listen(":8080")
}
