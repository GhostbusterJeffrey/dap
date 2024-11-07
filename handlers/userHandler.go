package handlers

import (
	"context"
	"log"

	"github.com/GhostbusterJeffrey/dap/config"
	"github.com/GhostbusterJeffrey/dap/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserName(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	// Convert userID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	// Fetch the user from the database
	collection := config.GetCollection("users")
	var user models.User

	filter := bson.M{"_id": objectID}
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Println("Error finding user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User not found"})
	}

	// Return the user's name
	return c.JSON(fiber.Map{"name": user.Name})
}

func GetUserData(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	// Convert userID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	// Fetch the user from the database
	userCollection := config.GetCollection("users")
	var user models.User

	userFilter := bson.M{"_id": objectID}
	err = userCollection.FindOne(context.Background(), userFilter).Decode(&user)
	if err != nil {
		log.Println("Error finding user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "User not found"})
	}

	// Fetch the user's projects from the database
	projectCollection := config.GetCollection("projects")
	var projects []models.Project

	projectFilter := bson.M{"userID": userID}
	cursor, err := projectCollection.Find(context.Background(), projectFilter)
	if err != nil {
		log.Println("Error finding projects:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error retrieving projects"})
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var project bson.M
		if err := cursor.Decode(&project); err != nil {
			log.Println("Error decoding project:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error decoding project"})
		}

		// Convert the created timestamp to an int64
		project["createdAt"] = int64(project["createdAt"].(primitive.DateTime))

		var p models.Project
		projectBytes, err := bson.Marshal(project)
		if err != nil {
			log.Println("Error marshaling project:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error marshaling project"})
		}
		if err := bson.Unmarshal(projectBytes, &p); err != nil {
			log.Println("Error unmarshaling project:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error unmarshaling project"})
		}
		p.OwnerID = userID
		projects = append(projects, p)
	}

	// Return the user's information including project info
	return c.JSON(fiber.Map{
		"name":      user.Name,
		"email":     user.Email,
		"createdAt": user.CreatedAt,
		"updatedAt": user.UpdatedAt,
		"projects":  projects,
	})
}
