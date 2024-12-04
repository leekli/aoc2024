package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
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
	p2Output := Part2(rawInput)

	fmt.Printf("Day 3, Part 2 Output: %d\n", p2Output)	
}

func Part1(input string) int {
	total := 0

	foundMulsSlice := ExtractMulsFromString(input)

	total = SumAllMuls(foundMulsSlice)

	return total
}

func Part2(input string) int {
	total := 0

	total = CorruptedMemory_Part2(input)

	return total
}

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

func CorruptedMemory_Part2(input string) int {
	sumTotal := 0
		
	// 'At the beginning of the program, mul instructions are enabled' - from AOC instructions
	mulEnabled := true

	// Regex patterns required
	mulPattern := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
    doPattern := regexp.MustCompile(`do\(\)`)
    dontPattern := regexp.MustCompile(`don't\(\)`)

    // Find all matches for mul(XXX,XXX) instructions and do() and don't() events, use string indexes
    mulMatches := mulPattern.FindAllStringSubmatchIndex(input, -1)
    doMatches := doPattern.FindAllStringIndex(input, -1)
    dontMatches := dontPattern.FindAllStringIndex(input, -1)

    // Combine all matches into a single slice of events to work through
    events := make([][3]int, 0, len(mulMatches)+len(doMatches)+len(dontMatches))

    for _, match := range mulMatches {
        events = append(events, [3]int{match[0], match[len(match)-1], 0}) // 0 is a mul() event
    }
    for _, match := range doMatches {
        events = append(events, [3]int{match[0], match[len(match)-1], 1}) // 1 is a do() event
    }
    for _, match := range dontMatches {
        events = append(events, [3]int{match[0], match[len(match)-1], 2}) // 2 is a don't() event
    }

    // Sort all events by their starting index
    sort.Slice(events, func(i, j int) bool {
        return events[i][0] < events[j][0]
    })

    // Process all the events in order
    for _, event := range events {
        switch event[2] {
			case 0: // A mul() event
				if mulEnabled {
					singleMul := input[event[0]:event[1] + 1]

					singleSum := MultiplySingleMul(singleMul)

					sumTotal += singleSum
				}
			case 1: // A do() event allows future mul's to happen
				mulEnabled = true
			case 2: // A don't() event stops future mul's from happening
				mulEnabled = false
        }
    }
	
	return sumTotal
}