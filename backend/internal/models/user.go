package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                uuid.UUID          `gorm:"column:id;primaryKey"`
	FirebaseUID       string             `gorm:"column:firebase_uid;unique;index"`
	Email             string             `gorm:"column:email"`
	Name              string             `gorm:"column:name"`
	SimulationBalance float64            `gorm:"column:simulation_balance"`
	CreatedAt         time.Time          `gorm:"column:created_at"`
	UpdatedAt         time.Time          `gorm:"column:updated_at"`
	SimulationTrades  []SimulationTrades `gorm:"foreignKey:UserID;references:ID"`
}

func (User) TableName() string {
	return "users"
}
