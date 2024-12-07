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

	// Part 1
	p1Output := Part1(rawInput)

	fmt.Printf("Day 7, Part 1 Output: %d\n", p1Output)

	// Part 2
	//p2Output := Part2(rawInput)

	//fmt.Printf("Day 7, Part 2 Output: %d\n", p2Output)	
}

func Part1(input string) int {
	total := 0

	equationsList := StringTo2DIntArray(input)

	for i := 0; i <len(equationsList); i++ {
		equationsListWithoutTestValue := equationsList[i][1:]

		operatorsToTry := []string{"+", "*"}

		testValue := GetTestValue(equationsList[i])

		combosToTry := GenerateCombinationsToTry(equationsListWithoutTestValue, operatorsToTry)

		isEquationTrue := CanEquationBeMadeTrue(combosToTry, testValue)

		if isEquationTrue {
			total = total + testValue
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

func StringTo2DIntArray(input string) [][]int {
	replacedInput := strings.Replace(input, ":", "", -1)

    rows := strings.Split(replacedInput, "\n")

    var numsList [][]int

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

        numsList = append(numsList, nums)
    }

    return numsList
}

func GetTestValue(equation []int) int {
	return equation[0]
}

func GenerateCombinationsToTry(equation []int, operators []string) []string {
	if len(equation) == 0 {
		return []string{}
	}

	if len(equation) == 1 {
		return []string{strconv.Itoa(equation[0])}
	}

	var combinationsCollection []string

	// Some good ol' recursion ♻️
	currentCollection := GenerateCombinationsToTry(equation[1:], operators)

	for _, current := range currentCollection {

		for _, op := range operators {
			expr := fmt.Sprintf("%s %s %s", strconv.Itoa(equation[0]), op, current)
			combinationsCollection = append(combinationsCollection, expr)
		}
	}

	return combinationsCollection
}

func CanEquationBeMadeTrue(combosToTry []string, testValue int) bool {
	isEquationTrue := false

	for i := 0; i < len(combosToTry); i++ {
		splitCombo := strings.Split(combosToTry[i], " ")

		firstNum, _ := strconv.Atoi(splitCombo[0])
		splitCombo = splitCombo[1:]

		currentTotal := firstNum

		for len(splitCombo) != 0 {
			if splitCombo[0] == "*" {
				num, _ := strconv.Atoi(splitCombo[1])

				currentTotal = currentTotal * num
			}

			if splitCombo[0] == "+" {
				num, _ := strconv.Atoi(splitCombo[1])

				currentTotal = currentTotal + num
			}

			splitCombo = splitCombo[2:]
		}

		if currentTotal == testValue {
			isEquationTrue = true
			break
		}
	}

	return isEquationTrue
}