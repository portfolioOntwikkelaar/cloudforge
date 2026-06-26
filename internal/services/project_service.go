package services

import (
	"cloudforge/internal/models"
	"cloudforge/internal/repositories"
	"fmt"
)

func CreateProject(userID uint, req models.CreateProjectRequest) error {

	fmt.Println("=== CreateProject gestart ===")
	project := models.Project{
		Name:        req.Name,
		GitURL:      req.GitURL,
		Branch:      req.Branch,
		Description: req.Description,
		UserID:      userID,
	}

	if err := repositories.CreateProject(&project); err != nil {
		return err
	}

	fmt.Println("Project created with ID:", project.ID)
	fmt.Println("Project GitURL:", project.GitURL)
	if project.GitURL != "" {
		if err := CloneRepository(project.ID, project.GitURL); err != nil {
			fmt.Println("failed to clone repository:", err)
			return err
		}
		fmt.Println("Repository cloned successfully")
	}

	return nil
}

func GetProjects(userID uint) ([]models.Project, error) {
	return repositories.GetProjectsByUserID(userID)
}

//fmt.Println("Project created with ID:", project.ID)
//fmt.Println("Project GitURL:", project.GitURL)

//fmt.Println("CloneRepository wordt aangeroepen")
