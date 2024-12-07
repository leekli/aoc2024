package main

import (
	"slices"
	"testing"
)

func TestStringTo2DIntArray_ReturnsNumsList2DArray(test *testing.T) {
	input := "190: 10 19"

	output := StringTo2DIntArray(input)

	if len(output) != 1 {
		test.Errorf("Expected: 1, Received: %d", len(output))
	}

	if output[0][0] != 190 {
		test.Errorf("Expected: 190, Received: %d", output[0][0])	
	}

	if output[0][1] != 10 {
		test.Errorf("Expected: 10, Received: %d", output[0][1])	
	}

	if output[0][2] != 19 {
		test.Errorf("Expected: 19, Received: %d", output[0][1])	
	}

	input = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

	output = StringTo2DIntArray(input)

	if len(output) != 9 {
		test.Errorf("Expected: 9, Received: %d", len(output))
	}
}

func TestGetTestValue_ReturnsCorrectFirstValue(test *testing.T) {
	input := []int{190, 10, 19}

	output := GetTestValue(input)

	if output != 190 {
		test.Errorf("Expected: 190, Received: %d", output)	
	}

	input = []int{3267, 81, 40, 27}

	output = GetTestValue(input)

	if output != 3267 {
		test.Errorf("Expected: 3267, Received: %d", output)	
	}

	input = []int{161011, 16, 10, 13}

	output = GetTestValue(input)

	if output != 161011 {
		test.Errorf("Expected: 161011, Received: %d", output)	
	}
}

func TestGenerateCombinationsToTry_ReturnsTrueForEquation(test *testing.T) {
	input := []int{10, 19}
	operators := []string{"+", "*"}

	output := GenerateCombinationsToTry(input, operators)

	if len(output) != 2 {
		test.Errorf("Expected: 2, Received: %v", len(output))	
	}

	if !slices.Contains(output, "10 + 19") {
		test.Errorf("Expected: '10 + 19'")			
	}

	if !slices.Contains(output, "10 * 19") {
		test.Errorf("Expected: '10 * 19'")			
	}

	input = []int{81, 40, 27}
	operators = []string{"+", "*"}

	output = GenerateCombinationsToTry(input, operators)

	if len(output) != 4 {
		test.Errorf("Expected: 4, Received: %v", len(output))	
	}
}

func TestCanEquationBeMadeTrue_ReturnsTrueForValidEquation(test *testing.T) {
	fullNumList := []int{190, 10, 19}
	numInput := fullNumList[1:]
	operators := []string{"+", "*"}
	testValue := GetTestValue(fullNumList)
	combos := GenerateCombinationsToTry(numInput, operators)

	output := CanEquationBeMadeTrue(combos, testValue)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)	
	}

	fullNumList = []int{3267, 81, 40, 27}
	numInput = fullNumList[1:]
	operators = []string{"+", "*"}
	testValue = GetTestValue(fullNumList)
	combos = GenerateCombinationsToTry(numInput, operators)

	output = CanEquationBeMadeTrue(combos, testValue)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)	
	}

	fullNumList = []int{292, 11, 6, 16, 20}
	numInput = fullNumList[1:]
	operators = []string{"+", "*"}
	testValue = GetTestValue(fullNumList)
	combos = GenerateCombinationsToTry(numInput, operators)

	output = CanEquationBeMadeTrue(combos, testValue)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)	
	}
}

func TestCanEquationBeMadeTrue_ReturnsFalseForNonValidEquation(test *testing.T) {
	fullNumList := []int{83, 17, 5}
	numInput := fullNumList[1:]
	operators := []string{"+", "*"}
	testValue := GetTestValue(fullNumList)
	combos := GenerateCombinationsToTry(numInput, operators)

	output := CanEquationBeMadeTrue(combos, testValue)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)	
	}

	fullNumList = []int{156, 15, 6}
	numInput = fullNumList[1:]
	operators = []string{"+", "*"}
	testValue = GetTestValue(fullNumList)
	combos = GenerateCombinationsToTry(numInput, operators)

	output = CanEquationBeMadeTrue(combos, testValue)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)	
	}

	fullNumList = []int{7290, 6, 8, 6, 15}
	numInput = fullNumList[1:]
	operators = []string{"+", "*"}
	testValue = GetTestValue(fullNumList)
	combos = GenerateCombinationsToTry(numInput, operators)

	output = CanEquationBeMadeTrue(combos, testValue)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)	
	}

	fullNumList = []int{161011, 16, 10, 13}
	numInput = fullNumList[1:]
	operators = []string{"+", "*"}
	testValue = GetTestValue(fullNumList)
	combos = GenerateCombinationsToTry(numInput, operators)

	output = CanEquationBeMadeTrue(combos, testValue)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)	
	}

	fullNumList = []int{192, 17, 8, 14}
	numInput = fullNumList[1:]
	operators = []string{"+", "*"}
	testValue = GetTestValue(fullNumList)
	combos = GenerateCombinationsToTry(numInput, operators)

	output = CanEquationBeMadeTrue(combos, testValue)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)	
	}

	fullNumList = []int{21037, 9, 7, 18, 13}
	numInput = fullNumList[1:]
	operators = []string{"+", "*"}
	testValue = GetTestValue(fullNumList)
	combos = GenerateCombinationsToTry(numInput, operators)

	output = CanEquationBeMadeTrue(combos, testValue)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)	
	}
}

func TestPart1_ReturnsCorrectTotalForTestInput(test *testing.T) {
	input := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

	output := Part1(input)

	if output != 3749 {
		test.Errorf("Expected: 3749, Received: %d", output)		
	}
}

func TestPart2_ReturnsCorrectTotalForTestInput(test *testing.T) {
	input := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

	output := Part2(input)

	if output != 11387 {
		test.Errorf("Expected: 11387, Received: %d", output)		
	}
}