package main

import (
	"fmt"
	"os"
	"time"
)

// ReadFile reads the content of a file and returns it as a string.
func ReadFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// CompareFileContent compares the previous and current content of the file.
func CompareFileContent(previousContent, currentContent string) {
	if previousContent != currentContent {
		fmt.Println("File content has changed!")
		fmt.Println("Previous content:\n", previousContent)
		fmt.Println("Current content:\n", currentContent)
	} else {
		fmt.Println("No changes detected.")
	}
}

// WatchFile polls the file for changes by checking its content at a regular interval.
func WatchFile(filePath string, interval time.Duration) {
	var previousContent string

	// Read the initial content of the file
	previousContent, err := ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading initial file:", err)
		return
	}
	fmt.Println("Tracking file:", filePath)

	for {
		// Read the current content of the file
		currentContent, err := ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		// Compare the previous content with the current content to detect changes
		CompareFileContent(previousContent, currentContent)

		// Update the previous content
		previousContent = currentContent

		// Sleep for the specified interval before checking again
		time.Sleep(interval)
	}
}

func main() {
	// Path to the file being tracked
	filePath := "example.txt"

	// Start watching the file with an interval of 1 second
	go WatchFile(filePath, 1*time.Second)

	// Keep the main program running
	select {}
}
