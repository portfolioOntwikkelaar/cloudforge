package database

import "cloudforge/internal/models"

func Migrate() error {
	return DB.AutoMigrate(&models.Build{}, &models.User{}, &models.Project{})
}
