package repositories

import (
	"cloudforge/internal/database"
	"cloudforge/internal/models"
)

func CreateProject(project *models.Project) error {
	return database.DB.Create(project).Error
}

func GetProjectsByUserID(userID uint) ([]models.Project, error) {
	var projects []models.Project
	err := database.DB.Where("user_id = ?", userID).Find(&projects).Error
	return projects, err
}
func GetProjectByID(id uint) (*models.Project, error) {
	var project models.Project
	if err := database.DB.First(&project, id).Error; err != nil {
		return nil, err
	}
	return &project, nil
}
