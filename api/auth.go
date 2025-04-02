package api

import (
	"strings"
	"time"

	"github.com/Abh1noob/trader.pro-be/internal/auth"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(authRepo *auth.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing or invalid Authorization header"})
		}

		idToken := strings.TrimPrefix(token, "Bearer ")
		uid, email, name, err := authRepo.VerifyFirebaseToken(idToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}

		exists, err := authRepo.DoesUserExist(uid)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "failed to check user existence",
				"details": err.Error(),
			})
		}

		if !exists {
			if err := authRepo.StoreUser(uid, email, name); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   "failed to store user",
					"details": err.Error(),
				})
			}
		}

		c.Cookie(&fiber.Cookie{
			Name:     "auth_token",
			Value:    idToken,
			Expires:  time.Now().Add(24 * time.Hour),
			HTTPOnly: true,
			Secure:   true,
			SameSite: "Lax",
		})

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Logged in successfully",
			"uid":     uid,
			"email":   email,
		})
	}
}
