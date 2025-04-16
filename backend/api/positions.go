package api

import (
	"github.com/Abh1noob/trader.pro-be/internal/models"
	"github.com/Abh1noob/trader.pro-be/internal/positions"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PositionHandler struct {
	DB   *gorm.DB
	Repo positions.Repository
}

func NewPositionHandler(db *gorm.DB, repo positions.Repository) *PositionHandler {
	return &PositionHandler{DB: db, Repo: repo}
}

func (h *PositionHandler) ListPositions(c *fiber.Ctx) error {
	firebaseUID, ok := c.Locals("uid").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	positions, err := h.Repo.ListPositionsByUser(h.DB, firebaseUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Failed to fetch positions"})
	}
	return c.JSON(positions)
}

func (h *PositionHandler) GetPositionByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if _, err := uuid.Parse(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid position ID"})
	}
	position, err := h.Repo.GetPositionByID(h.DB, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Position not found"})
	}
	return c.JSON(position)
}

func (h *PositionHandler) CreatePosition(c *fiber.Ctx) error {
	var position models.SimulationPositions
	if err := c.BodyParser(&position); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	firebaseUID, ok := c.Locals("uid").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var user models.User
	if err := h.DB.Where("firebase_uid = ?", firebaseUID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
	}

	position.ID = uuid.New()
	position.UserID = user.ID

	if err := h.Repo.CreatePosition(h.DB, &position); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Failed to create position"})
	}
	return c.Status(fiber.StatusCreated).JSON(position)
}

func (h *PositionHandler) UpdatePosition(c *fiber.Ctx) error {
	id := c.Params("id")
	if _, err := uuid.Parse(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid position ID"})
	}
	position, err := h.Repo.GetPositionByID(h.DB, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Position not found"})
	}
	if err := c.BodyParser(position); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.Repo.UpdatePosition(h.DB, position); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Failed to update position"})
	}
	return c.JSON(position)
}

func (h *PositionHandler) DeletePosition(c *fiber.Ctx) error {
	id := c.Params("id")
	if _, err := uuid.Parse(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid position ID"})
	}
	if err := h.Repo.DeletePosition(h.DB, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Failed to delete position"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
