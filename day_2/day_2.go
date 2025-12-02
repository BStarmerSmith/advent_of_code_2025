package main

import (
	"advent_of_code_2025/utils"
	"fmt"
	"strconv"
	"strings"
)

var P1invalidIdsCount int = 0
var P1totalIdsCount int = 0
var P2invalidIdsCount int = 0
var P2totalIdsCount int = 0

func Day2() {
	lines := utils.ReadFileLines("input.txt")
	ids := utils.SplitByDelimiter(lines[0], ",")
	for _, id := range ids {
		splitId := utils.SplitByDelimiter(id, "-")
		if splitId[0][0] == '0' || splitId[1][0] == '0' {
			continue
		}
		firstID, err := strconv.Atoi(splitId[0])
		if err != nil {
			panic(err)
		}
		lastID, err := strconv.Atoi(splitId[1])
		if err != nil {
			panic(err)
		}
		invalidIds(firstID, lastID)
	}
	fmt.Println("P1 Invalid IDs Count: ", P1invalidIdsCount)
	fmt.Println("P1 Total IDs Count: ", P1totalIdsCount)
	fmt.Println("P2 Invalid IDs Count: ", P2invalidIdsCount)
	fmt.Println("P2 Total IDs Count: ", P2totalIdsCount)

}

// invalid IDs by looking for any ID which is made only of some sequence of digits repeated twice
// None of the numbers have leading zeroes
func invalidIds(firstID int, lastID int) {

	for i := firstID; i <= lastID; i++ {
		valString := strconv.Itoa(i)
		if IsRepeatingP1(valString) {
			P1invalidIdsCount++
			P1totalIdsCount += i
		}
		if isRepeatingP2(valString) {
			P2invalidIdsCount++
			P2totalIdsCount += i
		}
	}
}

// P1 invalid IDs by looking for any ID which is made only of some sequence of digits repeated twice
func IsRepeatingP1(s string) bool {
	// Length must be even to be a sequence repeated exactly twice
	if len(s)%2 != 0 {
		return false
	}
	// Check if first half equals second half
	mid := len(s) / 2
	return s[:mid] == s[mid:]
}

// P2 invalid IDs by looking for any ID which is made only of some sequence of digits repeated twice
func isRepeatingP2(s string) bool {
	if len(s) <= 1 {
		return false
	}
	// Double the string and trim the first and last characters
	doubled := s + s
	trimmed := doubled[1 : len(doubled)-1]
	return strings.Contains(trimmed, s)
}
