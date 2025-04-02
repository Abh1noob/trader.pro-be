package auth

import (
	"context"
	"errors"
	"time"

	firebaseAuth "firebase.google.com/go/v4/auth"
	"github.com/Abh1noob/trader.pro-be/config"
	"github.com/Abh1noob/trader.pro-be/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	authClient *firebaseAuth.Client
	pgRepo     *PostgresRepo
}

func NewRepository(authClient *firebaseAuth.Client, db *config.DB) *Repository {
	return &Repository{
		authClient: authClient,
		pgRepo:     NewPostgresRepo(db),
	}
}

func (r *Repository) VerifyFirebaseToken(token string) (string, string, string, error) {

	tokenInfo, err := r.authClient.VerifyIDToken(context.Background(), token)
	if err != nil {
		return "", "", "", err
	}

	user, err := r.authClient.GetUser(context.Background(), tokenInfo.UID)
	if err != nil {
		return "", "", "", err
	}

	email, _ := tokenInfo.Claims["email"].(string)
	if email == "" {
		email = user.Email
	}

	name, _ := tokenInfo.Claims["name"].(string)
	if name == "" {
		name = user.DisplayName
	}

	return tokenInfo.UID, email, name, nil
}

func (repo *Repository) DoesUserExist(uid string) (bool, error) {
	var user models.User
	result := repo.pgRepo.DB.Where("firebase_uid = ?", uid).Take(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (repo *Repository) StoreUser(uid, email string, name string) error {

	if name == "" {
		name = "Unnamed User"
	}

	user := models.User{
		ID:                uuid.New(),
		FirebaseUID:       uid,
		Email:             email,
		Name:              name,
		SimulationBalance: 100000,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	return repo.pgRepo.DB.Create(&user).Error
}
