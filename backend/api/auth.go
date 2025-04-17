package api

import (
	"time"

	"github.com/Abh1noob/trader.pro-be/internal/auth"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(authRepo *auth.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {

		uid := c.Locals("uid")
		email := c.Locals("email")
		name := c.Locals("name")

		if uid == nil || email == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
		}

		uidStr := uid.(string)
		emailStr := email.(string)
		nameStr := ""
		if name != nil {
			nameStr = name.(string)
		}

		exists, err := authRepo.DoesUserExist(uidStr)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "failed to check user existence",
				"details": err.Error(),
			})
		}

		if !exists {
			if err := authRepo.StoreUser(uidStr, emailStr, nameStr); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   "failed to store user",
					"details": err.Error(),
				})
			}
		}

		c.Cookie(&fiber.Cookie{
			Name:     "auth_token",
			Value:    c.Get("Authorization"),
			Expires:  time.Now().Add(24 * time.Hour),
			HTTPOnly: true,
			Secure:   true,
			SameSite: "None",
		})

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Logged in successfully",
			"uid":     uidStr,
			"email":   emailStr,
		})
	}
}
