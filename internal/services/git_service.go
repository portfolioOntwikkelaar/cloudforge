package services

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const WorkspaceRoot = "/home/damian//go-projects/cloudforge/workspaces"

func CloneRepository(projectID uint, gitURL string) error {

	targetDir := filepath.Join(WorkspaceRoot, fmt.Sprintf("project-%d", projectID))
	if _, err := os.Stat(targetDir); err == nil {
		return nil

	}

	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return err

	}
	fmt.Println("Cloning repository from", gitURL, "to", targetDir)
	cmd := exec.Command("git", "clone", gitURL, targetDir)

	output, err := cmd.CombinedOutput()
	fmt.Println("Git clone output:", string(output))
	if err != nil {
		return fmt.Errorf("%v\n%s", err, string(output))
	}

	return nil
}
