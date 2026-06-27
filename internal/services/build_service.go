package services

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"cloudforge/internal/models"
	"cloudforge/internal/repositories"
)

func BuildProject(projectID uint) error {

	projectDir := filepath.Join(WorkspaceRoot, fmt.Sprintf("project-%d", projectID))

	imageName := fmt.Sprintf("cloudforge-project-%d", projectID)

	build := models.Build{
		ProjectID: projectID,
		Status:    "running",
	}

	if err := repositories.CreateBuild(&build); err != nil {
		return err
	}

	cmd := exec.Command("docker", "build", "-t", imageName, projectDir)
	output, err := cmd.CombinedOutput()

	build.Log = string(output)
	if err != nil {
		build.Status = "failed"
	} else {
		build.Status = "success"
	}

	if saveErr := repositories.UpdateBuild(&build); saveErr != nil {
		return saveErr
	}

	fmt.Println("=============Docker Build==============")
	fmt.Println(build.Log)
	fmt.Println("===========================")

	if err != nil {
		return fmt.Errorf("failed to build Docker image: %v", err)

	}
	return nil
}

func GetProjectBuilds(projectID uint) ([]models.Build, error) {
	return repositories.GetBuildsByProjectID(projectID)
}
func GetBuild(id uint) (*models.Build, error) {
	return repositories.GetBuildByID(id)
}
