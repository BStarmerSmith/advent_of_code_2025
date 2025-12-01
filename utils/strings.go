package utils

import "strings"

// SplitLines splits a multi-line string into lines
func SplitLines(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n")
}

// SplitByEmptyLines splits text by empty lines (useful for grouped inputs)
func SplitByEmptyLines(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n\n")
}

// Contains checks if a string slice contains a value
func Contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
