package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/GhostbusterJeffrey/Dap/handlers"
    "github.com/GhostbusterJeffrey/Dap/middleware"
)

func SetupUserRoutes(app *fiber.App) {
    userGroup := app.Group("/user", middleware.JWTMiddleware)
    userGroup.Get("/name", handlers.GetUserName)
}
