package positions

import (
	"github.com/Abh1noob/trader.pro-be/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type positionRepo struct{}

func NewPositionRepo() Repository {
	return &positionRepo{}
}

func (r *positionRepo) ListPositionsByUser(
	db *gorm.DB,
	firebaseUID string,
) ([]models.SimulationPositions, error) {
	var user models.User
	if err := db.Where("firebase_uid = ?", firebaseUID).First(&user).Error; err != nil {
		return nil, err
	}
	var positions []models.SimulationPositions
	if err := db.Where("user_id = ?", user.ID).Find(&positions).Error; err != nil {
		return nil, err
	}
	return positions, nil
}

func (r *positionRepo) GetPositionByID(
	db *gorm.DB,
	id string,
) (*models.SimulationPositions, error) {
	var position models.SimulationPositions
	if err := db.Where("id = ?", id).First(&position).Error; err != nil {
		return nil, err
	}
	return &position, nil
}

func (r *positionRepo) CreatePosition(db *gorm.DB, position *models.SimulationPositions) error {
	position.ID = uuid.New()
	return db.Create(position).Error
}

func (r *positionRepo) UpdatePosition(db *gorm.DB, position *models.SimulationPositions) error {
	return db.Save(position).Error
}

func (r *positionRepo) DeletePosition(db *gorm.DB, id string) error {
	return db.Delete(&models.SimulationPositions{}, "id = ?", id).Error
}
