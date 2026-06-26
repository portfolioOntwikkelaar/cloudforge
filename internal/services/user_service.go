package services

import (
	"cloudforge/internal/models"
	"cloudforge/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

func Register(email, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Email:    email,
		Password: string(hash),
		Role:     "user",
	}

	return repositories.CreateUser(&user)
}

func Login(email, password string) (*models.User, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}
