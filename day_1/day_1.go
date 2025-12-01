package main

import (
	"advent_of_code_2025/utils"
	"fmt"
)

var dialStart = 50
var lowerBound = 0
var upperBound = 99
var timesAtZero = 0

func splitString(s string) (string, int) {
	return s[:1], utils.MustAtoi(s[1:])
}

func Day1() {
	lines := utils.ReadFileLines("input.txt")
	for _, line := range lines {
		direction, distance := splitString(line)
		dialNext := rotateDial(dialStart, direction, distance)

		dialStart = dialNext
	}
	fmt.Println(timesAtZero)
}

// rotateDial recursively rotates the dial until it reaches the distance
func rotateDial(dialStart int, direction string, distance int) int {
	// Base case: no more distance to travel
	if distance == 0 {
		return dialStart
	}

	switch direction {
	case "L":
		nextPos := dialStart - 1
		if nextPos < lowerBound {
			nextPos = upperBound
		}
		if nextPos == 0 {
			timesAtZero++
		}
		return rotateDial(nextPos, direction, distance-1)

	case "R":
		nextPos := dialStart + 1
		if nextPos > upperBound {
			nextPos = lowerBound
		}
		if nextPos == 0 {
			timesAtZero++
		}
		return rotateDial(nextPos, direction, distance-1)

	default:
		panic("Invalid direction")
	}
}
