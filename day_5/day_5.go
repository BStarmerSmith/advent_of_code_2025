package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day5() {
	part1()
	part2()
}

type IDRange struct {
	lowerBound int
	upperBound int
}

func part1() {
	freshIDCount := 0
	fmt.Println("Part 1:")
	freshIDRange, IDs := readInput("input.txt")
	idRangeArray := idRangeToRange(freshIDRange)

	for _, id := range IDs {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			panic(fmt.Sprintf("Could not convert %s to int: %v", id, err))
		}
		if isIDInFreshIDRange(idInt, idRangeArray) {
			freshIDCount++
		}
	}
	fmt.Println("freshIDCount: ", freshIDCount)
}

func part2() {
	fmt.Println("Part 2:")
	freshIDRange, _ := readInput("input.txt")
	idRangeArray := idRangeToRange(freshIDRange)

	// Sort the ID ranges
	sortIDRange := mergeSortIDRange(idRangeArray, 0)
	// Flatten the ID ranges
	flatIDRange := flattenIDRange(sortIDRange)
	// Get the total number of valid IDs
	validIDCount := getTotalValidIDs(flatIDRange)
	fmt.Println("validIDCount: ", validIDCount)
}

// Flattens the ID ranges by removing any overlapping ranges
func flattenIDRange(idRangeArray []IDRange) []IDRange {
	final := []IDRange{}
	var id_prev IDRange = IDRange{}
	for i, idRange := range idRangeArray {
		if i == 0 {
			id_prev = idRange
			continue
		}
		if idRange.lowerBound <= id_prev.upperBound+1 {
			// They overlap or are adjacent
			id_prev.upperBound = max(id_prev.upperBound, idRange.upperBound)
		} else {
			// They don't overlap
			final = append(final, id_prev)
			id_prev = idRange
		}
		if i == len(idRangeArray)-1 {
			final = append(final, id_prev)
		}
	}
	return final
}

// Sorts the ID ranges using merge sort (sorts lowerBound)
func mergeSortIDRange(idRangeArray []IDRange, numJobs int) []IDRange {
	if len(idRangeArray) < 2 {
		return idRangeArray
	}
	middle := len(idRangeArray) / 2
	left := mergeSortIDRange(idRangeArray[:middle], numJobs+1)
	right := mergeSortIDRange(idRangeArray[middle:], numJobs+1)
	return TopDownMergeSort(left, right)
}

// Merges two sorted ID ranges
func TopDownMergeSort(a []IDRange, b []IDRange) []IDRange {
	final := []IDRange{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i].lowerBound < b[j].lowerBound {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final

}

// Gets the total number of valid IDs
func getTotalValidIDs(idRangeArray []IDRange) int {
	total := 0
	for _, idRange := range idRangeArray {
		total += idRange.upperBound - idRange.lowerBound + 1
	}
	return total
}

// Converts a slice of strings to a slice of IDRange
func idRangeToRange(idRange []string) []IDRange {
	idRangeArray := make([]IDRange, 0)
	for _, id := range idRange {
		idSplit := strings.Split(id, "-")
		lowerBound, err := strconv.Atoi(idSplit[0])
		if err != nil {
			panic(fmt.Sprintf("Could not convert %s to int: %v", id, err))
		}
		upperBound, err := strconv.Atoi(idSplit[1])
		if err != nil {
			panic(fmt.Sprintf("Could not convert %s to int: %v", id, err))
		}
		idRangeArray = append(idRangeArray, IDRange{lowerBound, upperBound})
	}
	return idRangeArray
}

// Checks if an ID is in the fresh ID range
func isIDInFreshIDRange(id int, freshIDRange []IDRange) bool {
	for _, idRange := range freshIDRange {
		if id >= idRange.lowerBound && id <= idRange.upperBound {
			return true
		}
	}
	return false
}

// readInput reads the input file and returns two arrays:
// - freshIdRange: lines before the first empty line
// - IDs: lines after the empty line
func readInput(filename string) (freshIdRange []string, IDs []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Could not open file %s: %v", filename, err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read lines until we hit an empty line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		freshIdRange = append(freshIdRange, line)
	}

	// Read remaining lines after the empty line
	for scanner.Scan() {
		IDs = append(IDs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading file %s: %v", filename, err))
	}

	return freshIdRange, IDs
}
