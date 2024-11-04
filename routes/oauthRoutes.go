package routes

import (
    "context"
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/GhostbusterJeffrey/Dap/config"
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

        return c.JSON(userInfo)
    })
}
