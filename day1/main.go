package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
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

	fmt.Printf("Day 1, Part 1 Output: %d\n", p1Output)
}

func Part1(input string) int {
	total := 0

	numListOne, numListTwo := ConvertInputToTwoLists(input)

	sortedListOne, sortedListTwo := SortBothNumLists(numListOne, numListTwo)

	total = CountDistancesBetweenListNums(sortedListOne, sortedListTwo)

	return total
}

func readFileToString(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)

	if err != nil {
		return "", err
	}
	
	return string(content), nil
}

func ConvertInputToTwoLists(input string) ([]int, []int) {
	var numListOne []int
    var numListTwo []int

	if input == "" {
		return numListOne, numListTwo
	}

	splitLines := strings.Split(input, "\n")

	for _, line := range splitLines {

        numbers := strings.Fields(line)

		num1, err1 := strconv.Atoi(numbers[0])
		num2, err2 := strconv.Atoi(numbers[1])

		if err1 == nil && err2 == nil {
			numListOne = append(numListOne, num1)
			numListTwo = append(numListTwo, num2)
		} else {
			fmt.Println("Error converting string to int:", err1, err2)
		}
    }

    return numListOne, numListTwo
}

func SortBothNumLists(numListOne []int, numListTwo []int) ([]int, []int) {
	sort.Ints(numListOne)
	sort.Ints(numListTwo)

	return numListOne, numListTwo
}

func CountDistancesBetweenListNums(numListOne []int, numListTwo []int) int {
	totalDistance := 0

	for i := 0; i < len(numListOne); i++ {
		currentDistanceDiff := numListTwo[i] - numListOne[i]

		if currentDistanceDiff < 0 {
			currentDistanceDiff = -currentDistanceDiff
		}

		totalDistance += currentDistanceDiff
	}

	return totalDistance
}