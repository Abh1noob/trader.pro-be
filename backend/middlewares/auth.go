package middlewares

import (
	"context"
	"fmt"
	"log"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
)

func FirebaseAuthMiddleware(app *firebase.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Skip middleware for public routes and login endpoint
		if strings.HasPrefix(c.Path(), "/public/") {
			return c.Next()
		}

		var tokenString string

		// Try to get token from Authorization header first
		authHeader := c.Get("Authorization")
		fmt.Print("Auth Header:", authHeader)
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			// If not in header, try to get from cookie
			tokenString = c.Cookies("auth_token")
			// Remove Bearer prefix if present in cookie
			if strings.HasPrefix(tokenString, "Bearer ") {
				tokenString = strings.TrimPrefix(tokenString, "Bearer ")
			}
		}

		fmt.Println("Auth Token:", tokenString)

		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{"error": "Missing Authorization header or auth_token cookie"})
		}

		client, err := app.Auth(context.Background())
		if err != nil {
			log.Println("Error getting Firebase Auth client:", err)
			return c.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{"error": "Auth client error"})
		}

		token, err := client.VerifyIDToken(context.Background(), tokenString)
		if err != nil {
			log.Println("Invalid Firebase token:", err)
			return c.Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{"error": "Invalid or expired token"})
		}

		c.Locals("uid", token.UID)
		c.Locals("email", token.Claims["email"])
		c.Locals("name", token.Claims["name"])

		return c.Next()
	}
}
