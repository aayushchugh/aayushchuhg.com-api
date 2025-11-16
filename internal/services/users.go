package services

import (
	"github.com/aayushchugh/ayushchugh.com-api/internal/database"
	"github.com/aayushchugh/ayushchugh.com-api/internal/models"
)

type UserService struct{}

func (r *UserService) Create(user *models.User) error {
	err := database.DB.Create(user).Error
	return err
}

func (r *UserService) GetById(id string) (*models.User, error) {
	var user models.User
	err := database.DB.First(&user, id).Error
	return &user, err
}

func (r *UserService) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
