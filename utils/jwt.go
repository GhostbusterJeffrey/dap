package utils

import (
    "os"
    "time"
    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // Replace with a secure key, and store it securely in environment variables

// GenerateToken creates a JWT token for the given user ID
func GenerateToken(userID string) (string, error) {
    jwtSecret := []byte(os.Getenv("JWT_SECRET"))
    claims := jwt.MapClaims{
        "userID": userID,
        "exp":    time.Now().Add(time.Hour * 72).Unix(), // Example expiration: 72 hours
    }

    // Make sure to use the correct signing method
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}