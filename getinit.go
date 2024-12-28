package main

import (
	"fmt"
	"os"
)

func GetInit() error {
	// Directories and files to create
	getDir := ".get"
	directories := []string{
		".get/objects",
		".get/refs/heads",
		".get/refs/tags",
		".get/hooks",
		".get/info",
		".get/logs",
		".get/logs/refs/heads",
		".get/logs/refs/tags",
	}
	files := map[string]string{
		".get/HEAD":                    "ref: refs/heads/main\n", // Points to the main branch
		".get/config":                  "[core]\n\trepositoryformatversion = 0\n\tfilemode = true\n\tbare = false\n",
		".get/description":             "Unnamed repository; edit this file to name the repository.\n",
		".get/info/exclude":            "# Ignore patterns for untracked files\n",
		".get/packed-refs":             "", // Empty initially
		".get/hooks/pre-commit.sample": "# Sample pre-commit hook\n",
	}

	// Create all required directories
	for _, dir := range directories {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// Create all required files with default content
	for filePath, content := range files {
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to create file %s: %w", filePath, err)
		}
	}

	fmt.Println("Initialized empty Get repository in", getDir)
	return nil
}
