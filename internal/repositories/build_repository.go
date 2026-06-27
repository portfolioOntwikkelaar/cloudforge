package repositories

import (
	"cloudforge/internal/database"
	"cloudforge/internal/models"
)

func CreateBuild(build *models.Build) error {
	return database.DB.Create(build).Error
}

func UpdateBuild(build *models.Build) error {
	return database.DB.Save(build).Error
}
