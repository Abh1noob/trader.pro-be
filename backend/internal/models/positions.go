package models

import (
	"time"

	"github.com/google/uuid"
)

type SimulationPositions struct {
	ID            uuid.UUID `gorm:"column:id;primaryKey"`
	UserID        uuid.UUID `gorm:"column:user_id"`
	Symbol        string    `gorm:"column:symbol"`
	Quantity      int       `gorm:"column:quantity"`
	AvgPrice      float64   `gorm:"column:avg_price"`
	CurrentPrice  float64   `gorm:"column:current_price"`
	UnrealizedPnl *float64  `gorm:"column:unrealized_pnl"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}

func (SimulationPositions) TableName() string {
	return "simulation_positions"
}
