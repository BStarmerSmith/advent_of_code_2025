package main

import (
	"advent_of_code_2025/utils"
	"fmt"
)

var emptyChar string = "."
var filledChar string = "@"
var canAccessChar string = "x"

func Day4() {
	lines1 := utils.ReadFile2DChar("input.txt")
	part1(lines1)
	lines2 := utils.ReadFile2DChar("input.txt")
	part2(lines2)
}

func part1(lines [][]string) {
	fmt.Println("Part 1:")
	freeSquares := 0
	for x, line := range lines {
		for y, char := range line {
			if char == filledChar {
				if isRoleFree(lines, x, y, 4) {
					lines[x][y] = canAccessChar
					freeSquares++
				}
			}
		}
	}

	fmt.Println("freeSquares: ", freeSquares)
}

func part2(lines [][]string) {
	fmt.Println("Part 2:")
	finished := false
	totalFreeSquares := 0
	for !finished {
		freeSquares := 0
		freeSquareLocations := make([][]int, 0)

		for x, line := range lines {
			for y, char := range line {
				if char == filledChar {
					if isRoleFree(lines, x, y, 4) {
						// lines[x][y] = canAccessChar
						freeSquares++
						freeSquareLocations = append(freeSquareLocations, []int{x, y})
					}
				}
			}
		}

		for _, location := range freeSquareLocations {
			lines[location[0]][location[1]] = canAccessChar
		}

		totalFreeSquares += freeSquares
		replaceCanAccessChar(lines)

		if freeSquares == 0 {
			finished = true
		}
	}

	fmt.Println("totalFreeSquares: ", totalFreeSquares)
}

func printLines(lines [][]string) {
	for _, line := range lines {
		for y, _ := range line {
			fmt.Print(line[y])
		}
		fmt.Println()
	}
}

func replaceCanAccessChar(lines [][]string) {
	for x, line := range lines {
		for y, char := range line {
			if char == canAccessChar {
				lines[x][y] = emptyChar
			}
		}
	}
}

func isOutOfBounds(lines [][]string, x int, y int) bool {
	return x < 0 || x >= len(lines) || y < 0 || y >= len(lines[x])
}

func isRoleFree(lines [][]string, x int, y int, maxSpaces int) bool {
	takenSquares := 0
	// Top Row
	if !checkRollsCanAccess(lines, x-1, y-1) {
		takenSquares++
	}
	if !checkRollsCanAccess(lines, x-1, y) {
		takenSquares++
	}

	if !checkRollsCanAccess(lines, x-1, y+1) {
		takenSquares++
	}
	// Middle Row
	if !checkRollsCanAccess(lines, x, y-1) {
		takenSquares++
	}
	if !checkRollsCanAccess(lines, x, y+1) {
		takenSquares++
	}

	// Bottom Row
	if !checkRollsCanAccess(lines, x+1, y-1) {
		takenSquares++
	}
	if !checkRollsCanAccess(lines, x+1, y) {
		takenSquares++
	}
	if !checkRollsCanAccess(lines, x+1, y+1) {
		takenSquares++
	}

	if takenSquares >= maxSpaces {
		return false
	}
	return true
}

func checkRollsCanAccess(lines [][]string, x int, y int) bool {
	if isOutOfBounds(lines, x, y) {
		return true
	}
	if lines[x][y] == filledChar || lines[x][y] == canAccessChar {
		return false
	} else {
		return true
	}
}
