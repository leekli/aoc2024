package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Data setup
	rawInput, err := readFileToString("input.txt")

	if err != nil {
		os.Exit(-1)
	}

	//Part 1
	p1Output := Part1(rawInput)

	fmt.Printf("Day 2, Part 1 Output: %d\n", p1Output)

	// Part 2
	//p2Output := Part2(rawInput)

	//fmt.Printf("Day X, Part 2 Output: %d\n", p2Output)	
}

func Part1(input string) int {
	safeLevels := 0

	string2DArray := StringTo2DArray(input)

	safeLevels = CountTotalSafeLevels(string2DArray)

	return safeLevels
}

func Part2() {}

func readFileToString(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)

	if err != nil {
		return "", err
	}
	
	return string(content), nil
}

func StringTo2DArray(input string) [][]int {
    rows := strings.Split(input, "\n")

    var result [][]int

    for _, row := range rows {
        numStrs := strings.Fields(row)

        var nums []int
        for _, numStr := range numStrs {
            num, err := strconv.Atoi(numStr)

            if err != nil {
                continue
            }

            nums = append(nums, num)
        }

        result = append(result, nums)
    }

    return result
}

func CountTotalSafeLevels(levels [][]int) int {
	totalLevelsSafe := 0

	for i := 0; i < len(levels); i++ {
		isLevelSafe := IsLevelSafe(levels[i])

		if isLevelSafe {
			totalLevelsSafe++
		}
	}

	return totalLevelsSafe
}

func IsLevelSafe(level []int) bool {
	levelIsSafe := false

	allNumsIncOrDec := false
	allNumsWithinRange := false

	for i := 0; i < len(level) - 1; i++ {
		numsWithinRange := IsNumChangingWithinRange(level[i], level[i + 1])

		if numsWithinRange {
			allNumsWithinRange = true
		}

		if !numsWithinRange {
			allNumsWithinRange = false
			break
		}
	}

	allNumsIncOrDec = AreAllNumsIncreasingOrDecreasing(level)

	if allNumsWithinRange && allNumsIncOrDec {
		levelIsSafe = true
	}

	return levelIsSafe
}

func AreAllNumsIncreasingOrDecreasing(level []int) bool {
	increasing := true
    decreasing := true

    for i := 1; i < len(level); i++ {
		// If next num is higher, then the level is not decreasing
        if level[i] > level[i-1] {
            decreasing = false
        }

		// If next num is lower, then the level is not increasing
        if level[i] < level[i-1] {
            increasing = false
        }

		// This is neither an increase or a decrease
		if level[i] == level[i-1] {
			increasing = false
			decreasing = false
		}
    }

    return increasing || decreasing
}

func IsNumChangingWithinRange(num1 int, num2 int) bool {
	isNumWithinRange := false

	difference := num1 - num2

	if (difference >= 1 && difference <= 3) || (difference <= -1 && difference >= -3) {
        isNumWithinRange = true
    }

	return isNumWithinRange
}