package auth

import (
	"errors"
	"time"

	"github.com/Abh1noob/trader.pro-be/config"
	"github.com/Abh1noob/trader.pro-be/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostgresRepo struct {
	DB *gorm.DB
}

func NewPostgresRepo(db *config.DB) *PostgresRepo {
	return &PostgresRepo{DB: db.DB}
}

func (r *PostgresRepo) StoreUserIfNotExists(firebaseUID, email string) error {
	var user models.User
	result := r.DB.Where("firebase_uid = ?", firebaseUID).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		newUser := models.User{
			ID:                uuid.New(),
			FirebaseUID:       firebaseUID,
			Email:             email,
			Name:              "Default Name",
			SimulationBalance: 100000.0,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		}
		return r.DB.Create(&newUser).Error
	}

	if result.Error != nil {
		return result.Error
	}

	if user.Email != email {
		user.Email = email
		user.UpdatedAt = time.Now()
		return r.DB.Save(&user).Error
	}

	return nil
}

func (r *PostgresRepo) GetUserByID(firebaseUID string) (*models.User, error) {
	var user models.User
	result := r.DB.Where("firebase_uid = ?", firebaseUID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
