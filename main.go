package main

import (
	"log"

	"github.com/GhostbusterJeffrey/dap/config"
	"github.com/GhostbusterJeffrey/dap/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Initialize Google OAuth
	config.InitGoogleOAuth()

	// Initialize Fiber
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Connect to MongoDB
	config.ConnectDB()

	// Set up routes
	routes.Setup(app)

	// Start the server
	log.Fatal(app.Listen(":8080"))
}
