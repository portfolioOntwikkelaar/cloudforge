package database

import "cloudforge/internal/models"

func Migrate() error {
	return DB.AutoMigrate(&models.User{}, &models.Project{})
}
