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

	fmt.Printf("Day 5, Part 1 Output: %d\n", p1Output)

	// Part 2
	//p2Output := Part2(rawInput)

	//fmt.Printf("Day 5, Part 2 Output: %d\n", p2Output)	
}

func Part1(input string) int {
	total := 0

	orderRulesList, pagesToProduceList := GenerateOrderRulesAndPagesToProduceLists(input)

	for i := 0; i < len(pagesToProduceList); i++ {
		if IsUpdateInRightOrder(orderRulesList, pagesToProduceList[i]) {
			middleNum := FindMiddlePageNumber(pagesToProduceList[i])

			total += middleNum
		}
	}

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

func GenerateOrderRulesAndPagesToProduceLists(input string) ([][]int, [][]int) {
	splitInput := strings.Split(input, "\n\n")

	rawOrderRulesList := splitInput[0]
	rawPagesToProduceList := splitInput[1]

	var finalOrderRulesList [][]int
	lines1 := strings.Split(rawOrderRulesList, "\n")

	for _, line := range lines1 {
		pair := strings.Split(line, "|")
		num1, _ := strconv.Atoi(pair[0])
		num2, _ := strconv.Atoi(pair[1])

		finalOrderRulesList = append(finalOrderRulesList, []int{num1, num2})
	}

	var finalPagesToProduceList [][]int
	lines2 := strings.Split(rawPagesToProduceList, "\n")

	for _, line := range lines2 {
		strNums := strings.Split(line, ",")
		var nums []int
		for _, strNum := range strNums {
			num, _ := strconv.Atoi(strNum)
			nums = append(nums, num)
		}
		finalPagesToProduceList = append(finalPagesToProduceList, nums)
	}

	return finalOrderRulesList, finalPagesToProduceList
}

func IsUpdateInRightOrder(ruleList [][]int, pagesList []int) bool {
	// Directed graph problem!
	
	isUpdateInRightOrder := false

	// Start a map with pages set to 0
	inDegreeMap := GenerateInDegreeMap(pagesList)

	// Work through each page number in the pagesList given, then the rules list and if the current page has a dependency on another page (the 2nd number in each rulesList pair), increase the value in the map (in degree value) if it exists
	for i := 0; i < len(pagesList); i++ {
		for j := 0; j < len(ruleList); j++ {
			if ruleList[j][0] == pagesList[i] {
				_, exists := inDegreeMap[ruleList[j][1]]

				if exists {
					inDegreeMap[ruleList[j][1]]++
				}	
			}
		}
	}

	// Now that the in-degree map is populated, we need take the values and order them
	type mapKeyValues struct {
		Key   int
		Value int
	}

	// Get the key-value pairs from the in-degree map produced
	var sortedPairs []mapKeyValues
	for key, value := range inDegreeMap {
		sortedPairs = append(sortedPairs, mapKeyValues{key, value})
	}

	// Sort the values into an order from 0 (no dependency) incrementally
	sort.Slice(sortedPairs, func(i, j int) bool {
		return sortedPairs[i].Value < sortedPairs[j].Value
	})

	// Take the sorted pairs and set up a final sortedKeys list by taking just the key which should now be in sorted order
	var sortedKeys []int
	for _, pair := range sortedPairs {
		sortedKeys = append(sortedKeys, pair.Key)
	}

	// Now compare the 2 lists, if in same order as the originally given pageList, then the update is in the right order, false if not
	for i := 0; i < len(sortedKeys); i++ {
		if len(sortedKeys) == len(pagesList) {
			if sortedKeys[i] == pagesList[i] {
				isUpdateInRightOrder = true
			} else {
				isUpdateInRightOrder = false
				break
			}
		}
	}

	return isUpdateInRightOrder
}

func GenerateInDegreeMap(pagesList []int) map[int]int {
	var inDegreeMap = make(map[int]int)

	for i := 0; i < len(pagesList); i++ {
		_, exists := inDegreeMap[pagesList[i]]

		if !exists {
			inDegreeMap[pagesList[i]] = 0
		}
	}

	return inDegreeMap
}

func FindMiddlePageNumber(pageList []int) int {
	if len(pageList) <= 2 {
		return pageList[0]
	}

	if len(pageList) == 3 {
		return pageList[1]
	}

	return pageList[len(pageList)/2]
}