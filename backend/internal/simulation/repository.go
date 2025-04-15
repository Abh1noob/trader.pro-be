package simulation

import (
	"github.com/Abh1noob/trader.pro-be/internal/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateTrade(db *gorm.DB, trade *models.SimulationTrades) error
	GetTradeByID(db *gorm.DB, id string) (*models.SimulationTrades, error)
	ListTradesByUser(db *gorm.DB, userID string, limit, offset int) ([]models.SimulationTrades, error)
	UpdateTrade(db *gorm.DB, trade *models.SimulationTrades) error
}
