package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		// No arguments provided - run all days
		runAllDays()
		return
	}

	day := os.Args[1]
	dayDir := fmt.Sprintf("day_%s", day)

	// Check if directory exists
	if _, err := os.Stat(dayDir); os.IsNotExist(err) {
		fmt.Printf("Day %s not found. Directory '%s' does not exist.\n", day, dayDir)
		return
	}

	// Run the day's solution
	runDay(day, dayDir)
}

func runDay(day string, dayDir string) {
	cmd := exec.Command("go", "run", ".")
	cmd.Dir = filepath.Join(".", dayDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Running Day %s...\n", day)
	fmt.Println("================")

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running day %s: %v\n", day, err)
	}
	fmt.Println() // Add blank line between days
}

func runAllDays() {
	fmt.Println("Running all available days...")
	fmt.Println("============================")
	fmt.Println()

	// Find all day_* directories
	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	foundDays := false
	for _, entry := range entries {
		if entry.IsDir() && len(entry.Name()) > 4 && entry.Name()[:4] == "day_" {
			foundDays = true
			day := entry.Name()[4:]
			runDay(day, entry.Name())
		}
	}

	if !foundDays {
		fmt.Println("No day directories found.")
	}
}
