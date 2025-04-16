package positions

import (
	"github.com/Abh1noob/trader.pro-be/internal/models"
	"gorm.io/gorm"
)

type Repository interface {
	ListPositionsByUser(db *gorm.DB, firebaseUID string) ([]models.SimulationPositions, error)
	GetPositionByID(db *gorm.DB, id string) (*models.SimulationPositions, error)
	CreatePosition(db *gorm.DB, position *models.SimulationPositions) error
	UpdatePosition(db *gorm.DB, position *models.SimulationPositions) error
	DeletePosition(db *gorm.DB, id string) error
}
