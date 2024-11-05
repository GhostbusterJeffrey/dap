package handlers

import (
	"context"
	"strings"
	"time"

	"github.com/GhostbusterJeffrey/dap/config"
	"github.com/GhostbusterJeffrey/dap/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProject(c *fiber.Ctx) error {
    userID := c.Locals("userID").(string)

    // Define a struct to hold the incoming project data
    var project struct {
        Name        string `json:"name"`
        Description string `json:"description"`
    }

    // Parse the JSON body into the project struct
    if err := c.BodyParser(&project); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    // Check if the name is empty
    if project.Name == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Project name is required"})
    }

    // Prepare the project data to insert into MongoDB
    collection := config.GetCollection("projects")
    newProject := bson.M{
        "name":        project.Name,
        "description": project.Description,
        "userID":      userID,
        "createdAt":   time.Now(), // Use time.Now() for a proper date object
    }

    // Insert the new project into the database
    result, err := collection.InsertOne(context.Background(), newProject)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create project"})
    }

    // Return a success response with the project ID
    return c.JSON(fiber.Map{"message": "Project created", "projectID": result.InsertedID})
}

// List of valid HTTP methods
var validHTTPMethods = map[string]bool{
    "GET":     true,
    "POST":    true,
    "PUT":     true,
    "DELETE":  true,
    "PATCH":   true,
    "HEAD":    true,
    "OPTIONS": true,
}

func CreateAPIRoute(c *fiber.Ctx) error {
    userID := c.Locals("userID").(string)

    // Parse and validate the request body
    var request struct {
        Path      string `json:"path"`
        Method    string `json:"method"`
        Response  string `json:"response"`
        ProjectID string `json:"projectID"`
    }
    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
    }

    // Ensure all fields are provided
    if request.Path == "" || request.Method == "" || request.Response == "" || request.ProjectID == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "All fields (path, method, response, projectID) are required"})
    }

    // Validate the HTTP method
    request.Method = strings.ToUpper(request.Method)
    if !validHTTPMethods[request.Method] {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid HTTP method"})
    }

    // Validate the ProjectID and check if the user owns it
    projectObjectID, err := primitive.ObjectIDFromHex(request.ProjectID)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ProjectID format"})
    }

    projectCollection := config.GetCollection("projects")
    var project bson.M
    err = projectCollection.FindOne(context.Background(), bson.M{"_id": projectObjectID, "userID": userID}).Decode(&project)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Project not found or you do not have permission to access this project"})
    }

    // Check for duplicate path in the project
    routeCollection := config.GetCollection("api_routes")
    var existingRoute bson.M
    err = routeCollection.FindOne(context.Background(), bson.M{"path": request.Path, "projectid": request.ProjectID}).Decode(&existingRoute)
    if err == nil {
        return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "An API route with this path already exists for the project"})
    } else if err.Error() != "mongo: no documents in result" {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error checking for duplicate path"})
    }

    // Save the API route in your database
    route := models.APIRoute{
        Path:      request.Path,
        Method:    request.Method,
        Response:  request.Response,
        ProjectID: request.ProjectID,
    }

    _, err = routeCollection.InsertOne(c.Context(), route)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create API route"})
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "API route created successfully"})
}

func GetProjectInfo(c *fiber.Ctx) error {
    // Extract the project ID from the URL parameters
    projectID := c.Params("id")

    // Convert the projectID to a MongoDB ObjectID
    objectID, err := primitive.ObjectIDFromHex(projectID)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid project ID format"})
    }

    // Fetch the project from the database
    collection := config.GetCollection("projects")
    var project bson.M
    err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&project)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Project not found"})
    }

    return c.JSON(project)
}
