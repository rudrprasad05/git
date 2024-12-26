package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Commit represents a single commit in the file's history.
type Commit struct {
	Timestamp time.Time
	Content   string
	Message   string
}

// ReadFile reads the content of a file and returns it as a string.
func ReadFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// WatchFile polls the file for changes and creates commits for each change.
func WatchFile(filePath string, interval time.Duration) {
	var previousContent string
	var history []Commit

	// Read the initial content of the file
	previousContent, err := ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading initial file:", err)
		return
	}
	fmt.Println("Tracking file:", filePath)

	// Add the initial state as the first commit
	history = append(history, Commit{
		Timestamp: time.Now(),
		Content:   previousContent,
		Message:   "Initial commit",
	})

	for {
		// Read the current content of the file
		currentContent, err := ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		// Detect changes and create a new commit if necessary
		if previousContent != currentContent {
			commit := Commit{
				Timestamp: time.Now(),
				Content:   currentContent,
				Message:   fmt.Sprintf("File changed at %v", time.Now()),
			}
			history = append(history, commit)
			fmt.Println("New commit created:")
			fmt.Printf("Timestamp: %v\nMessage: %s\n\n", commit.Timestamp, commit.Message)
		}

		// Update the previous content
		previousContent = currentContent

		// Sleep for the specified interval before checking again
		time.Sleep(interval)
	}
}

// ViewCommitHistory displays the commit history.
func ViewCommitHistory(history []Commit) {
	fmt.Println("Commit History:")
	for i, commit := range history {
		fmt.Printf("Commit %d:\n", i+1)
		fmt.Printf("Timestamp: %v\n", commit.Timestamp)
		fmt.Printf("Message: %s\n", commit.Message)
		fmt.Println("Content:")
		fmt.Println(commit.Content)
		fmt.Println("--------------------")
	}
}

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
func GetAllFiles(ignoredFiles map[string]bool) ([]string, error) {
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
		_, ok := ignoredFiles[dirName]
		if ok {
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

func FindDifference(a, b []string) []string {
	// Create a map to store elements of B for quick lookup
	bMap := make(map[string]bool)
	for _, val := range b {
		bMap[val] = true
	}

	// Find elements in A that are not in B
	var difference []string
	for _, val := range a {
		if !bMap[val] {
			difference = append(difference, val)
		}
	}

	return difference
}

func main() {
	ignoredFiles, err := GetIgnoreFile(".getignore")
	if err != nil {
		fmt.Println("File doesnt exist, skipping", err)
		return
	}

	allFiles, allFileErr := GetAllFiles(ignoredFiles)
	if allFileErr != nil {
		fmt.Println("Error:", allFileErr)
		return
	}

	fmt.Println(allFiles)

	// watchFiles := FindDifference(allFiles, ignoredFiles)

	// for _, s := range watchFiles {
	// 	go WatchFile(s, 1*time.Second)

	// }

	// Start watching the file with an interval of 1 second

	// Keep the main program running
	select {}
}
