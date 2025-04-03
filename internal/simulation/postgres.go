package simulation

import (
	"github.com/Abh1noob/trader.pro-be/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type simulationRepo struct{}

func NewSimulationRepo() Repository {
	return &simulationRepo{}
}

func (r *simulationRepo) CreateTrade(db *gorm.DB, trade *models.SimulationTrades) error {
	trade.ID = uuid.New()
	return db.Create(trade).Error
}

func (r *simulationRepo) GetTradeByID(db *gorm.DB, id string) (*models.SimulationTrades, error) {
	var trade models.SimulationTrades
	err := db.Where("id = ?", id).First(&trade).Error
	return &trade, err
}

func (r *simulationRepo) ListTradesByUser(db *gorm.DB, userID string, limit, offset int) ([]models.SimulationTrades, error) {
	var userIDStr string

	err := db.Table("users").Select("id").Where("firebase_uid = ?", userID).Scan(&userIDStr).Error
	if err != nil {
		return nil, err
	}

	UserID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, err
	}

	var trades []models.SimulationTrades
	err = db.Where("user_id = ?", UserID).Limit(limit).Offset(offset).Find(&trades).Error
	return trades, err
}

func (r *simulationRepo) UpdateTrade(db *gorm.DB, trade *models.SimulationTrades) error {
	return db.Save(trade).Error
}
