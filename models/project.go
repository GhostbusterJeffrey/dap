package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    Name        string             `bson:"name"`
    Description string             `bson:"description"`
    OwnerID     string             `bson:"ownerID"` // The ID of the user who created the project
    CreatedAt   int64              `bson:"createdAt"`
}

type APIRoute struct {
    Path        string `json:"path"`
    Method      string `json:"method"` // e.g., "GET", "POST"
    Response    string `json:"response"` // The response data
    ProjectID   string `json:"projectID"`
}