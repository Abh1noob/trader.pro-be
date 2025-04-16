package models

import (
	"time"

	"github.com/google/uuid"
)

type SimulationTrades struct {
	ID          uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	UserID      uuid.UUID `gorm:"column:user_id"       json:"user_id"`
	Symbol      string    `gorm:"column:symbol"        json:"symbol"`
	TradeType   string    `gorm:"column:trade_type"    json:"trade_type"`
	Quantity    int       `gorm:"column:quantity"      json:"quantity"`
	Price       float64   `gorm:"column:price"         json:"price"`
	TotalAmount float64   `gorm:"column:total_amount"  json:"total_amount"`
	LimitPrice  float64   `gorm:"column:limit_price"   json:"limit_price"`
	StopLoss    float64   `gorm:"column:stop_loss"     json:"stop_loss"`
	Timestamp   time.Time `gorm:"column:timestamp"     json:"timestamp"`
	ExecutedAt  time.Time `gorm:"column:executed_at"   json:"executed_at"`
	Status      string    `gorm:"column:status"        json:"status"`
}

func (SimulationTrades) TableName() string {
	return "simulation_trades"
}
