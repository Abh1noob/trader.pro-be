package routes

import (
	"github.com/Abh1noob/trader.pro-be/api"

	"github.com/gofiber/fiber/v2"
)

func MountSimulationRoutes(app *fiber.App, handler *api.SimulationHandler) {

	simRoute := app.Group("/api/v1/simulation")
	simRoute.Post("/trade", handler.CreateTrade)
	simRoute.Get("/trade/:id", handler.GetTradeByID)
	simRoute.Get("/trades", handler.ListTradesByUser)
}
