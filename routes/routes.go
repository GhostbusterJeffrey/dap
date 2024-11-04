package routes

import (
	"github.com/GhostbusterJeffrey/dap/handlers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/hello", handlers.Hello)
	SetupOAuthRoutes(app)
	SetupUserRoutes(app)
}
