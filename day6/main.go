package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	// Data setup
	rawInput, err := readFileToString("input.txt")

	if err != nil {
		os.Exit(-1)
	}

	// Part 1
	p1Output := Part1(rawInput)

	fmt.Printf("Day 6, Part 1 Output: %d\n", p1Output)

	// Part 2
	//p2Output := Part2(rawInput)

	//fmt.Printf("Day 6, Part 2 Output: %d\n", p2Output)	
}

func Part1(input string) int {
	// Start a counter for distinct positions
	totalPositions := 0

	// Set up main map with input & an empty visited logging map
	mainMap := StringTo2DArray(input)
	visitedLogMap := CreateVisitedLogArray(mainMap)

	// While guard is still on the map
	// --> Get guards current position and which way they are facing
	// --> Get the co ordinates of the next location depending which way they are facing, and check if that next location is an obstruction or not
	// --> if it is an obstruction: turn right 90 degrees
	// --> if it is NOT an obstruction, move forward

	for IsGuardOnMapAnymore(mainMap) {
		currRow, currCol := GetGuardCurrentPosition(mainMap)
		currDirection := GetWhichWayGuardIsFacing(mainMap)

		nextRow, nextCol := GetNextStepFoward_DirectionDepending(mainMap, currRow, currCol, currDirection)

		isNextMoveAnObstruction := IsNextMoveAnObstruction(mainMap, nextRow, nextCol)

		if isNextMoveAnObstruction {
			Turn90DegreesRight(mainMap, currRow, currCol, currDirection)
		}

		if !isNextMoveAnObstruction {
			TakeStepFoward(mainMap, currRow, currCol, currDirection, visitedLogMap)
		}
	}

	// After breaking out of the map when guard has gone, count all visited locations
	totalPositions = CountTotalLocationsVisited(visitedLogMap)

	return totalPositions
}

func Part2() {}

func readFileToString(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)

	if err != nil {
		return "", err
	}
	
	return string(content), nil
}

func StringTo2DArray(input string) [][]string {
    lineRows := strings.Split(input, "\n")

    var fullMap [][]string

    for _, lineRow := range lineRows {
        splitChars := strings.Split(lineRow, "")

        var finalChars []string

        for _, numStr := range splitChars {

            finalChars = append(finalChars, numStr)
        }

        fullMap = append(fullMap, finalChars)
    }

    return fullMap
}

func GetGuardCurrentPosition(theMap [][]string) (int, int) {
	var row int
	var col int

	permittedArrows := []string{"<", ">", "^", "v"}

	for i := 0; i < len(theMap); i++ {
		for j := 0; j < len(theMap[i]); j++ {
			if slices.Contains(permittedArrows, theMap[i][j]) {
				row = i
				col = j
				break
			}
		}
	}

	return row, col
}

func GetWhichWayGuardIsFacing(theMap [][]string) string {
	var direction string

	permittedArrows := []string{"<", ">", "^", "v"}

	for i := 0; i < len(theMap); i++ {
		for j := 0; j < len(theMap[i]); j++ {
			if slices.Contains(permittedArrows, theMap[i][j]) {
				direction = theMap[i][j]
				break
			}
		}
	}

	return direction
}

func GetNextStepFoward_DirectionDepending(theMap [][]string, row int, col int, direction string) (int, int) {
		// Get next location moving up
		if direction == "^" {
			newRow := row - 1
			newCol := col
	
			if newRow >= 0 && newRow < len(theMap) && newCol >= 0 && newCol < len(theMap[0]) {
				return newRow, newCol	
			}
		}

		// Get next location moving down
		if direction == "v" {
			newRow := row + 1
			newCol := col
	
			if newRow >= 0 && newRow < len(theMap) && newCol >= 0 && newCol < len(theMap[0]) {
				return newRow, newCol	
			}
		}

		// Get next location moving right
		if direction == ">" {
			newRow := row
			newCol := col + 1
	
			if newRow >= 0 && newRow < len(theMap) && newCol >= 0 && newCol < len(theMap[0]) {
				return newRow, newCol	
			}
		}

		// Get next location moving left
		if direction == "<" {
			newRow := row
			newCol := col - 1
	
			if newRow >= 0 && newRow < len(theMap) && newCol >= 0 && newCol < len(theMap[0]) {
				return newRow, newCol	
			}
		}

	return 0, 0
}

func IsNextMoveAnObstruction(theMap [][]string, nextStepRow int, nextStepCol int) bool {
	isObstruction := false

	if nextStepRow >= 0 && nextStepRow <= len(theMap) && nextStepCol >= 0 && nextStepCol <= len(theMap[0]) {
		if (theMap[nextStepRow][nextStepCol] == "#") {
			isObstruction = true
		}
	}	

	return isObstruction
}

func Turn90DegreesRight(theMap [][]string, row int, col int, direction string) {
	// If facing up, turn right 90 degrees to face right
	if theMap[row][col] == "^" {
		theMap[row][col] = ">"
		return
	}

	// If facing right, turn right 90 degrees to face down
	if theMap[row][col] == ">" {
		theMap[row][col] = "v"
		return
	}

	// If facing down, turn right 90 degrees to face left
	if theMap[row][col] == "v" {
		theMap[row][col] = "<"
		return
	}

	// If facing left, turn right 90 degrees to face up
	if theMap[row][col] == "<" {
		theMap[row][col] = "^"
		return
	}
}

func TakeStepFoward(theMap [][]string, row int, col int, direction string, visitedMap [][]int) {
	// Move up 
	if direction == "^" {
		// Mark this location as visited before moving on
		MarkLocationAsVisited(visitedMap, row, col)

		// Change current to '.'
		theMap[row][col] = "."

		newRow := row - 1
		newCol := col

		if newRow >= 0 && newRow < len(theMap) && newCol >= 0 && newCol < len(theMap[0]) {
			// Change next position to '^'
			theMap[newRow][newCol] = "^"	
		}
	}

	// Move down
	if direction == "v" {
		// Mark this location as visited before moving on
		MarkLocationAsVisited(visitedMap, row, col)

		// Change current to '.'
		theMap[row][col] = "."

		newRow := row + 1
		newCol := col

		if newRow >= 0 && newRow < len(theMap) && newCol >= 0 && newCol < len(theMap[0]) {
			// Change next position to '^'
			theMap[newRow][newCol] = "v"	
		}
	}

	// Move right
	if direction == ">" {
		// Mark this location as visited before moving on
		MarkLocationAsVisited(visitedMap, row, col)

		// Change current to '.'
		theMap[row][col] = "."

		newRow := row
		newCol := col + 1

		if newRow >= 0 && newRow < len(theMap) && newCol >= 0 && newCol < len(theMap[0]) {
			// Change next position to '^'
			theMap[newRow][newCol] = ">"	
		}
	}

	if direction == "<" {
		// Mark this location as visited before moving on
		MarkLocationAsVisited(visitedMap, row, col)

		// Change current to '.'
		theMap[row][col] = "."

		newRow := row
		newCol := col - 1

		if newRow >= 0 && newRow < len(theMap) && newCol >= 0 && newCol < len(theMap[0]) {
			// Change next position to '^'
			theMap[newRow][newCol] = "<"	
		}
	}
}

func IsGuardOnMapAnymore(theMap [][]string) bool {
	isGuardOnMap := false

	permittedArrows := []string{"<", ">", "^", "v"}

	for i := 0; i < len(theMap); i++ {
		for j := 0; j < len(theMap[i]); j++ {
			if slices.Contains(permittedArrows, theMap[i][j]) {
				isGuardOnMap = true
				break
			}
		}
	}

	return isGuardOnMap
}

func CreateVisitedLogArray(theMap [][]string) [][]int {
   rows := len(theMap)
   cols := len(theMap[0])

   visitedLogArray := make([][]int, rows)
   
   for i := range visitedLogArray {
	visitedLogArray[i] = make([]int, cols)
   }

   return visitedLogArray
}

func MarkLocationAsVisited(visitedLogMap [][]int, row int, col int) {
	if (visitedLogMap[row][col] == 0) {
		visitedLogMap[row][col] = 1
	}
}

func CountTotalLocationsVisited(visitedLogMap [][]int) int {
	totalVisited := 0

	for i := 0; i < len(visitedLogMap); i++ {
		for j := 0; j < len(visitedLogMap[i]); j++ {
			if visitedLogMap[i][j] == 1 {
				totalVisited++
			}
		}
	}

	return totalVisited
}