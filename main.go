package main

import (
	"fiber-api/handlers"
	"fiber-api/infrastructure"
	"fiber-api/repositories"
	"fiber-api/routes"
	"fiber-api/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// connect DB
	infrastructure.ConnectDB()

	app := fiber.New()

	// Middleware
	app.Use(logger.New(), cors.New())

	// Dependency Injection
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Register Routes
	routes.SetupUserRoutes(app, userHandler)

	app.Listen(":8000")
}
