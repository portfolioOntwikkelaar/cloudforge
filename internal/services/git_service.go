package services

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func CloneRepository(gitURL string, targetDir string) error {
	os.RemoveAll(targetDir)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(ctx, "git", "clone", "--depth", "1", gitURL, targetDir)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to clone repository: %v, output: %s", err, string(output))
	}

	return nil

}
