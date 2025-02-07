package routes

import (
	"fiber-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, userHandler *handlers.UserHandler) {
	userGroup := app.Group("/users")
	userGroup.Get("/", userHandler.GetUsers)
	userGroup.Get("/:id", userHandler.GetUser)
	userGroup.Post("/", userHandler.CreateUser)
	userGroup.Put("/", userHandler.UpdateUser)
	userGroup.Delete("/", userHandler.DeleteUser)
}
