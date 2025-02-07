package main

import (
	"fiber-api/config"
	"fiber-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// connect DB
	config.ConnectDB()

	app := fiber.New()

	// Middleware
	app.Use(logger.New(), cors.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome to Fiber API"})
	})

	// routes API modul users
	routes.SetupUserRoutes(app)

	app.Listen(":8000")
}
