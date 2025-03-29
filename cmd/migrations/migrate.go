package main

import (
	"log"

	"github.com/Abh1noob/trader.pro-be/config"
	"github.com/Abh1noob/trader.pro-be/internal/models"
)

func main() {

	db := config.InitDB()

	err := db.AutoMigrate(
		&models.Users{},
		&models.SimulationTrades{},
		&models.SimulationPositions{},
	)

	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database migrations completed successfully")
}
