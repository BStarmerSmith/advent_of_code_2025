package utils

import (
	"strconv"
	"strings"
)

// ParseInts parses a slice of strings into integers
func ParseInts(strs []string) []int {
	nums := make([]int, len(strs))
	for i, s := range strs {
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			panic(err)
		}
		nums[i] = n
	}
	return nums
}

// ParseIntsSplit parses a single string with space-separated integers
func ParseIntsSplit(str string) []int {
	parts := strings.Fields(str)
	return ParseInts(parts)
}

// MustAtoi converts string to int, panics on error
func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

// Sum returns the sum of a slice of integers
func Sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Min returns the minimum of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Abs returns the absolute value
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
