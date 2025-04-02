package middlewares

import (
	"context"
	"log"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
)

func FirebaseAuthMiddleware(app *firebase.App) fiber.Handler {
	return func(c *fiber.Ctx) error {

		if strings.HasPrefix(c.Path(), "/public/") {
			return c.Next()
		}

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Authorization header format"})
		}

		client, err := app.Auth(context.Background())
		if err != nil {
			log.Println("Error getting Firebase Auth client:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Auth client error"})
		}

		token, err := client.VerifyIDToken(context.Background(), tokenString)
		if err != nil {
			log.Println("Invalid Firebase token:", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
		}

		c.Locals("uid", token.UID)
		c.Locals("email", token.Claims["email"])
		c.Locals("name", token.Claims["name"])

		return c.Next()
	}
}
