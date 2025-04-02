package models

import (
	"time"
)

type SimulationPositions struct {
	ID            string    `gorm:"column:id;primaryKey"`
	UserID        string    `gorm:"column:user_id"`
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
