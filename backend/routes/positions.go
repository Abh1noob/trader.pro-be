package routes

import (
	"github.com/Abh1noob/trader.pro-be/api"
	"github.com/gofiber/fiber/v2"
)

func MountPositionRoutes(app *fiber.App, handler *api.PositionHandler) {
	posRoute := app.Group("/api/v1/positions")
	posRoute.Get("/", handler.ListPositions)
	posRoute.Get("/:id", handler.GetPositionByID)
	posRoute.Post("/", handler.CreatePosition)
	posRoute.Put("/:id", handler.UpdatePosition)
	posRoute.Delete("/:id", handler.DeletePosition)
}
