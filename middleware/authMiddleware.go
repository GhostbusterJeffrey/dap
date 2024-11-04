package middleware

import (
    "log"
    "os"
    "strings"

    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(c *fiber.Ctx) error {
    authHeader := c.Get("Authorization")
    if authHeader == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
    }

    // Extract the token
    parts := strings.Split(authHeader, " ")
    if len(parts) != 2 || parts[0] != "Bearer" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Authorization header format"})
    }
    tokenString := parts[1]

    jwtSecret := []byte(os.Getenv("JWT_SECRET"))

    claims := jwt.MapClaims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
        }
        return jwtSecret, nil
    })

    if err != nil || !token.Valid {
        log.Println("JWT parsing error:", err)
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
    }

    userID, ok := claims["userID"].(string)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
    }

    c.Locals("userID", userID)
    return c.Next()
}
