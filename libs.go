package main

import (
	"fmt"
	"os"
)

func ReadFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
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
