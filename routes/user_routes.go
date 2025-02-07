package routes

import (
	"fiber-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	userGroup := app.Group("/users")
	userGroup.Get("/", handlers.GetUsers)
	userGroup.Get("/:id", handlers.GetUser)
	userGroup.Post("/", handlers.CreateUser)
	userGroup.Put("/", handlers.UpdateUser)
	userGroup.Delete("/", handlers.DeleteUser)
}
