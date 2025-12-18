package routes

import (
	"go-user-api/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, userHandler *handler.UserHandler) {
	app.Post("/users", userHandler.CreateUser)
	app.Get("/users/:id", userHandler.GetUser)
	app.Get("/users", userHandler.ListUsers)
	app.Put("/users/:id", userHandler.UpdateUser)
	app.Delete("/users/:id", userHandler.DeleteUser)
}
