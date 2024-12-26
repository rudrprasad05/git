package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func GetIgnoreFile(filePath string) (map[string]bool, error) {
	ignoreMap := make(map[string]bool)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newLine := strings.Replace(scanner.Text(), "/", "", 1)
		ignoreMap[newLine] = true
	}

	// Check for errors that may have occurred during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ignoreMap, nil
}

func GetAllFilesAfterIgnoring(ignoredFiles map[string]bool) ([]string, error) {
	var files []string

	// Get the current directory
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// Walk through the directory structure
	err = filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// check if its in ignore file
		dirName := strings.Replace(d.Name(), "/", "", 1)

		if ignoredFiles[dirName] {
			if d.IsDir() {
				// Skip the entire directory
				return filepath.SkipDir
			}
			return nil
		}
		// If it's a file, add it to the list
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}
