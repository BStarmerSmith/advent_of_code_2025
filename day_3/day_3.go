package main

import (
	"advent_of_code_2025/utils"
	"fmt"
	"strconv"
	"strings"
)

func Day3() {
	lines := utils.ReadFileLines("input.txt")
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	totalJoltage := 0
	for _, line := range lines {
		char_array := strings.Split(line, "")
		joltage := getJoltage(char_array)
		totalJoltage += joltage
	}
	fmt.Println("Part 1:")
	fmt.Println("Total Joltage: " + strconv.Itoa(totalJoltage))
}

func part2(lines []string) {
	totalJoltage := 0
	for _, line := range lines {
		char_array := strings.Split(line, "")
		joltage := getJoltageN(char_array, 12, []int{})
		totalJoltage += arrayToInt(joltage)
	}
	fmt.Println("Part 2:")
	fmt.Println("Total Joltage: " + strconv.Itoa(totalJoltage))
}

// getJoltage finds the highest and second highest digits in a character array
// and returns them concatenated as a single integer
// Example: ["1", "9", "3", "2"] -> 93 (highest=9, secondHighest=3)
func getJoltage(char_array []string) int {
	highest, secondHighest := 0, 0
	for i, char := range char_array {
		charInt, err := strconv.Atoi(char)
		if err != nil {
			fmt.Println("Error converting char to int")
			return 0
		}
		// If we find a new highest (not at the last position), reset secondHighest
		if highest < charInt && i != len(char_array)-1 {
			secondHighest = 0
			highest = charInt
			continue
		}
		if secondHighest < charInt {
			secondHighest = charInt
		}
	}
	result, err := strconv.Atoi(strconv.Itoa(highest) + strconv.Itoa(secondHighest))
	if err != nil {
		fmt.Println("Error converting int to string")
		return 0
	}
	return result
}

// getJoltageN recursively finds the n highest values from a character array
// It works by finding the highest value, adding it to joltages, then recursing
// on the remaining array after the highest value's position
func getJoltageN(char_array []string, n int, joltages []int) []int {
	if n == 0 {
		return joltages
	}
	highest := 0
	index := 0

	// Search through array leaving room for remaining n-1 values
	for i, char := range char_array[:len(char_array)-n+1] {
		charInt, err := strconv.Atoi(char)
		if err != nil {
			fmt.Println("Error converting char to int")
			return []int{}
		}
		if highest < charInt {
			highest = charInt
			index = i
		}
	}
	joltages = append(joltages, highest)
	return getJoltageN(char_array[index+1:], n-1, joltages)
}

// arrayToInt converts an array of integers into a single int by concatenating them
// Example: [9, 8, 7] -> 987
func arrayToInt(arr []int) int {
	var result int
	for _, digit := range arr {
		result = result*10 + digit
	}
	return result
}
