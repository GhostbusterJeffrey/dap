package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/GhostbusterJeffrey/dap/handlers"
    "github.com/GhostbusterJeffrey/dap/middleware"
)

func SetupProjectRoutes(app *fiber.App) {
    projectGroup := app.Group("/projects", middleware.JWTMiddleware)

    projectGroup.Post("/", handlers.CreateProject)
    projectGroup.Post("/routes", handlers.CreateAPIRoute)
    projectGroup.Get("/:id", handlers.GetProjectInfo)
}
