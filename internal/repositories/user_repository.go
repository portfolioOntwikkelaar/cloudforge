package repositories

import (
	"cloudforge/internal/database"
	"cloudforge/internal/models"
)

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := database.DB.Where("email = ?", email).First(&user).Error

	return &user, err
}
