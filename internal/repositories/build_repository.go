package repositories

import (
	"cloudforge/internal/database"
	"cloudforge/internal/models"
)

func GetBuildsByProjectID(projectID uint) ([]models.Build, error) {
	var builds []models.Build
	err := database.DB.Where("project_id = ?", projectID).Order("created_at DESC").Find(&builds).Error
	return builds, err
}

func GetBuildByID(id uint) (*models.Build, error) {
	var build models.Build
	err := database.DB.First(&build, id).Error

	if err != nil {
		return nil, err
	}
	return &build, err
}

func CreateBuild(build *models.Build) error {
	return database.DB.Create(build).Error
}

func UpdateBuild(build *models.Build) error {
	return database.DB.Save(build).Error
}
