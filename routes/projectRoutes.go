package routes

import (
	"github.com/GhostbusterJeffrey/dap/handlers"
	"github.com/GhostbusterJeffrey/dap/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupProjectRoutes(app *fiber.App) {
	projectGroup := app.Group("/projects", middleware.JWTMiddleware)

	projectGroup.Post("/", handlers.CreateProject)
	projectGroup.Post("/routes", handlers.CreateAPIRoute)
	projectGroup.Get("/:id", handlers.GetProjectInfo)
}

func SetupDynamicRoutes(app *fiber.App) {
    app.All("/api/:projectID/*", handlers.HandleAPIRoute)
}