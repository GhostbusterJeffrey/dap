package routes

import (
	"github.com/GhostbusterJeffrey/dap/handlers"
	"github.com/GhostbusterJeffrey/dap/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	userGroup := app.Group("/user", middleware.JWTMiddleware)
	userGroup.Get("/name", handlers.GetUserName)
}
