package utils

import (
	"bufio"
	"fmt"
	"os"
)

// ReadFileLines reads a file and returns its lines as a slice of strings
func ReadFileLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Could not open file %s: %v", filename, err))
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading file %s: %v", filename, err))
	}

	return lines
}
