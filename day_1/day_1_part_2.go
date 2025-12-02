package main

import (
	"advent_of_code_2025/utils"
	"fmt"
	"strconv"
)

var p2dialStart = 50
var p2timesAtZero = 0

func Day1P2() {
	lines := utils.ReadFileLines("input.txt")
	for _, line := range lines {
		direction, distance := splitString(line)
		dialNext := rotateDialP2(p2dialStart, direction, distance)

		p2dialStart = dialNext
	}
	fmt.Println("P2 Times at Zero: " + strconv.Itoa(p2timesAtZero))
}

// rotateDial recursively rotates the dial until it reaches the distance
func rotateDialP2(dialStart int, direction string, distance int) int {
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
			p2timesAtZero++
		}
		return rotateDialP2(nextPos, direction, distance-1)

	case "R":
		nextPos := dialStart + 1
		if nextPos > upperBound {
			nextPos = lowerBound
		}
		if nextPos == 0 {
			p2timesAtZero++
		}
		return rotateDialP2(nextPos, direction, distance-1)

	default:
		panic("Invalid direction")
	}
}
