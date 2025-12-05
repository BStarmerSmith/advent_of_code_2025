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
}

func part1() {
	freshIDCount := 0
	fmt.Println("Part 1:")
	freshIDRange, IDs := readInput("input.txt")
	idRangeInts := idRangeToInts(freshIDRange)

	for _, id := range IDs {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			panic(fmt.Sprintf("Could not convert %s to int: %v", id, err))
		}
		if isIDInFreshIDRange(idInt, idRangeInts) {
			freshIDCount++
		}
	}
	fmt.Println("freshIDCount: ", freshIDCount)
}

func idRangeToInts(idRange []string) []int {
	idRangeInts := make([]int, 0)
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
		for x := lowerBound; x <= upperBound; x++ {
			idRangeInts = append(idRangeInts, x)
		}
	}
	return idRangeInts
}

func isIDInFreshIDRange(id int, freshIDRange []int) bool {
	for _, freshID := range freshIDRange {
		if id == freshID {
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
