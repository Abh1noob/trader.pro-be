package routes

import (
	"github.com/Abh1noob/trader.pro-be/api"
	"github.com/Abh1noob/trader.pro-be/internal/auth"
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(app *fiber.App, authRepo *auth.Repository) {
	app.Post("/api/v1/auth/login", api.LoginHandler(authRepo))
}
