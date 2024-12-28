package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Commit represents a single commit in the file's history.
type Commit struct {
	Timestamp time.Time
	Content   string
	Message   string
}

func CreateBlob(filePath string) ([]byte, string, error) {
	// Read the content of the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, "", fmt.Errorf("failed to read file: %v", err)
	}

	// Create the "blob" object by prepending the size of the content
	size := len(content)
	header := fmt.Sprintf("blob %d\000", size) // Git's header for blob (size + content)
	blobContent := append([]byte(header), content...)

	// Calculate the SHA-1 hash of the "blob" object
	hash := sha1.New()
	hash.Write(blobContent)
	hashSum := hash.Sum(nil)

	err = os.WriteFile(fmt.Sprintf("%x", hashSum), blobContent, 0644)
	if err != nil {
		fmt.Println("err1")
		log.Fatal(err)
	}

	// Return the SHA-1 hash and the string representation
	return hashSum, fmt.Sprintf("%x", hashSum), nil
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
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("get> ")
		if !scanner.Scan() {
			// Handle EOF (Ctrl+D / Ctrl+Z)
			fmt.Println("\nEOF detected. Use the 'exit' command to quit.")
			continue
		}

		command := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if command == "" {
			// Skip empty commands
			continue
		}

		switch command {
		case "get init":
			if err := initializeGet(); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Repository initialized.")
			}
		case "exit":
			fmt.Println("Exiting Get...")
			return
		default:
			fmt.Println("Unknown command. Supported commands: get init, exit")
		}
	}
}
