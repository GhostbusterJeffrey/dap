package handlers

import (
    "context"
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/GhostbusterJeffrey/Dap/config"
    "github.com/GhostbusterJeffrey/Dap/models"
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
