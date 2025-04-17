package api

import (
	"strconv"
	"time"

	"github.com/Abh1noob/trader.pro-be/internal/models"
	"github.com/Abh1noob/trader.pro-be/internal/simulation"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type SimulationHandler struct {
	DB   *gorm.DB
	Repo simulation.Repository
}

func NewSimulationHandler(db *gorm.DB) *SimulationHandler {
	return &SimulationHandler{DB: db, Repo: simulation.NewSimulationRepo()}
}

func (h *SimulationHandler) CreateTrade(c *fiber.Ctx) error {
	var trade models.SimulationTrades
	if err := c.BodyParser(&trade); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	firebaseUID := c.Locals("uid").(string)

	var user models.User
	if err := h.DB.Where("firebase_uid = ?", firebaseUID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
	}

	trade.UserID = user.ID
	trade.TotalAmount = float64(trade.Quantity) * trade.Price
	trade.Timestamp = time.Now()

	if err := h.Repo.CreateTrade(h.DB, &trade); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Failed to create trade"})
	}

	return c.Status(fiber.StatusCreated).JSON(trade)
}

func (h *SimulationHandler) GetTradeByID(c *fiber.Ctx) error {
	id := c.Params("id")
	trade, err := h.Repo.GetTradeByID(h.DB, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Trade not found"})
	}

	return c.JSON(trade)
}

func (h *SimulationHandler) ListTradesByUser(c *fiber.Ctx) error {
	userID := c.Locals("uid").(string)
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	trades, err := h.Repo.ListTradesByUser(h.DB, userID, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Failed to retrieve trades"})
	}

	return c.JSON(trades)
}
