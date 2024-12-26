package p_ackage

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Ensure a subcommand is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: get <command>")
		return
	}

	// Get the subcommand
	command := os.Args[1]

	switch command {
	case "init":
		initializeRepo()
	case "status":
		showStatus()
	default:
		fmt.Printf("get: Unknown command '%s'\n", command)
		fmt.Println("Available commands: init, status")
	}
}

// initializeRepo initializes a new 'get' repository by creating a .get folder
func initializeRepo() {
	// Define the path to the .get directory
	getDir := filepath.Join(".", ".get")

	// Check if .get already exists
	if _, err := os.Stat(getDir); !os.IsNotExist(err) {
		fmt.Println("A 'get' repository already exists in this directory.")
		return
	}

	// Create the .get directory
	err := os.Mkdir(getDir, 0755)
	if err != nil {
		fmt.Println("Error creating .get directory:", err)
		return
	}

	fmt.Println("Initialized empty 'get' repository in", getDir)
}

// showStatus checks and displays the status of the repository
func showStatus() {
	// Define the path to the .get directory
	getDir := filepath.Join(".", ".get")

	// Check if .get exists
	if _, err := os.Stat(getDir); os.IsNotExist(err) {
		fmt.Println("No 'get' repository found. Run 'get init' first.")
		return
	}

	fmt.Println("Repository status:")
	fmt.Println("- .get directory exists")
	// Additional status checks can be added here
}
