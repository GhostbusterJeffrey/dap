package routes

import (
	"context"
	"log"
	"time"

	"github.com/GhostbusterJeffrey/Dap/config"
	"github.com/GhostbusterJeffrey/Dap/models"
	"github.com/GhostbusterJeffrey/Dap/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func SetupOAuthRoutes(app *fiber.App) {
	app.Get("/auth/login", func(c *fiber.Ctx) error {
		url := config.GoogleOAuthConfig.AuthCodeURL("f8b6271e-6854-4109-a451-3e21067f9bcb")
		return c.Redirect(url)
	})

	app.Get("/auth/callback", func(c *fiber.Ctx) error {
		code := c.Query("code")
		if code == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Code not found in the URL")
		}

		token, err := config.GoogleOAuthConfig.Exchange(context.Background(), code)
		if err != nil {
			log.Println("Error exchanging code for token:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange token")
		}

		client := config.GoogleOAuthConfig.Client(context.Background(), token)
		service, err := oauth2.NewService(context.Background(), option.WithHTTPClient(client))
		if err != nil {
			log.Println("Error creating OAuth2 service:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to create OAuth2 service")
		}

		userInfo, err := service.Userinfo.Get().Do()
		if err != nil {
			log.Println("Error getting user info:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user info")
		}

		// Connect to the users collection in MongoDB
		collection := config.GetCollection("users")
		var user models.User

		// Check if the user already exists
		filter := bson.M{"email": userInfo.Email}
		err = collection.FindOne(context.Background(), filter).Decode(&user)

		if err != nil { // User does not exist, create a new one
			user = models.User{
				ID:        primitive.NewObjectID(),
				Email:     userInfo.Email,
				Name:      userInfo.Name,
				Picture:   userInfo.Picture,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			_, err = collection.InsertOne(context.Background(), user)
			if err != nil {
				log.Println("Error creating new user:", err)
				return c.Status(fiber.StatusInternalServerError).SendString("Failed to create user")
			}
		}

		// Generate a JWT token for the user
		authToken, err := utils.GenerateToken(user.ID.Hex())
		if err != nil {
			log.Println("Error generating token:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate auth token")
		}

		// Return the auth token and user info
		return c.JSON(fiber.Map{
			"authToken": authToken,
			"user":      user,
		})
	})
}
