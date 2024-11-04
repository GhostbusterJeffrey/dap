package main

import (
	"log"

	"github.com/GhostbusterJeffrey/dap/config"
	"github.com/GhostbusterJeffrey/dap/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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

	// Connect to MongoDB
	config.ConnectDB()

	// Set up routes
	routes.Setup(app)

	// Start the server
	log.Fatal(app.Listen(":8080"))
}
