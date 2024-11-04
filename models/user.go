package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type User struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    Email     string             `bson:"email"`
    Name      string             `bson:"name"`
    Picture   string             `bson:"picture"`
    CreatedAt time.Time          `bson:"created_at"`
    UpdatedAt time.Time          `bson:"updated_at"`
}
