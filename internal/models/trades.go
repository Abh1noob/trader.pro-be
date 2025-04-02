package models

import (
	"time"
)

type SimulationTrades struct {
	ID          string    `gorm:"column:id;primaryKey"`
	UserID      string    `gorm:"column:user_id"`
	Symbol      string    `gorm:"column:symbol"`
	TradeType   string    `gorm:"column:trade_type"`
	Quantity    int       `gorm:"column:quantity"`
	Price       float64   `gorm:"column:price"`
	TotalAmount float64   `gorm:"column:total_amount"`
	LimitPrice  float64   `gorm:"column:limit_price"`
	StopLoss    float64   `gorm:"column:stop_loss"`
	Timestamp   time.Time `gorm:"column:timestamp"`
	ExecutedAt  time.Time `gorm:"column:executed_at"`
	Status      string    `gorm:"column:status"`
}

func (SimulationTrades) TableName() string {
	return "simulation_trades"
}
