package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/GhostbusterJeffrey/Dap/handlers"
)

func Setup(app *fiber.App) {
    app.Get("/hello", handlers.Hello)
    SetupOAuthRoutes(app)
    SetupUserRoutes(app)
}
