package routes

import (
	"github.com/Abh1noob/trader.pro-be/api"
	"github.com/Abh1noob/trader.pro-be/internal/auth"
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(app *fiber.App, authRepo *auth.Repository) {

	authGroup := app.Group("/api/v1/auth")
	authGroup.Post("/login", api.LoginHandler(authRepo))
}
