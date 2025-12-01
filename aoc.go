package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run aoc.go <day>")
		fmt.Println("Example: go run aoc.go 1")
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
	cmd := exec.Command("go", "run", ".")
	cmd.Dir = filepath.Join(".", dayDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Running Day %s...\n", day)
	fmt.Println("================")

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running day %s: %v\n", day, err)
	}
}
