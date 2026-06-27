package services

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func BuildProject(projectID uint) error {

	projectDir := filepath.Join(WorkspaceRoot, fmt.Sprintf("project-%d", projectID))

	imageName := fmt.Sprintf("cloudforge-project-%d", projectID)

	cmd := exec.Command("docker", "build", "-t", imageName, projectDir)
	output, err := cmd.CombinedOutput()

	fmt.Println("=============Docker Build==============")
	fmt.Println(string(output))
	fmt.Println("===========================")

	if err != nil {
		return fmt.Errorf("failed to build Docker image: %v", err)

	}
	return nil
}
