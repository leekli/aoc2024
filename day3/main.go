package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Data setup
	rawInput, err := readFileToString("input.txt")

	if err != nil {
		os.Exit(-1)
	}

	//Part 1
	p1Output := Part1(rawInput)

	fmt.Printf("Day 3, Part 1 Output: %d\n", p1Output)

	// Part 2
	//p2Output := Part2(rawInput)

	//fmt.Printf("Day X, Part 2 Output: %d\n", p2Output)	
}

func Part1(input string) int {
	total := 0

	foundMulsSlice := ExtractMulsFromString(input)

	total = SumAllMuls(foundMulsSlice)

	return total
}

func Part2() {}

func readFileToString(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)

	if err != nil {
		return "", err
	}
	
	return string(content), nil
}

func ExtractMulsFromString(input string) []string {
	foundMuls := []string{}

	mulRegEx := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	matches := mulRegEx.FindAllString(input, -1)

    for _, match := range matches {
		foundMuls = append(foundMuls, match)
    }

	return foundMuls
}

func MultiplySingleMul(singleMul string) int {
	sumTotal := 0

	getNumsRegex := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	match := getNumsRegex.FindStringSubmatch(singleMul)

	num1, err1 := strconv.Atoi(match[1])
    num2, err2 := strconv.Atoi(match[2])

    if err1 != nil || err2 != nil {
        fmt.Printf("Could not convert numbers")

		return 0
    }

	sumTotal = num1 * num2

	return sumTotal
}

func SumAllMuls(mulsFound []string) int {
	total := 0

	for i := 0; i < len(mulsFound); i++ {
		mulTotal := MultiplySingleMul(mulsFound[i])

		total += mulTotal
	}


	return total
}