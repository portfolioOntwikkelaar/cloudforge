package services

import (
	"cloudforge/internal/models"
	"cloudforge/internal/repositories"
)

func CreateProject(userID uint, req models.CreateProjectRequest) error {
	project := models.Project{
		Name:        req.Name,
		GitURL:      req.GitURL,
		Branch:      req.Branch,
		Description: req.Description,
		UserID:      userID,
	}

	return repositories.CreateProject(&project)
}
