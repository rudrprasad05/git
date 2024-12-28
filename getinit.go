package main

import (
	"fmt"
	"os"
)

func initializeGet() error {
	// Directories and files to create
	getDir := ".get"
	objectsDir := ".get/objects"
	refsDir := ".get/refs"
	headFile := ".get/HEAD"

	// Create .get directory
	if err := os.MkdirAll(objectsDir, 0755); err != nil {
		return fmt.Errorf("failed to create objects directory: %w", err)
	}

	if err := os.MkdirAll(refsDir, 0755); err != nil {
		return fmt.Errorf("failed to create refs directory: %w", err)
	}

	// Create HEAD file
	headContent := "ref: refs/heads/main\n"
	if err := os.WriteFile(headFile, []byte(headContent), 0644); err != nil {
		return fmt.Errorf("failed to create HEAD file: %w", err)
	}

	fmt.Println("Initialized empty Get repository in", getDir)
	return nil
}
