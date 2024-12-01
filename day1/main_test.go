package main

import (
	"testing"
)

func TestConvertInputToTwoLists_ReturnsEmptySlicesForEmptyInput(test *testing.T) {
	input := ""

	numListOne, numListTwo := ConvertInputToTwoLists(input)

	if len(numListOne) != 0 {
		test.Errorf("Expected: 0, Received: %d", len(numListOne))
	}

	if len(numListTwo) != 0 {
		test.Errorf("Expected: 0, Received: %d", len(numListTwo))
	}
}

func TestConvertInputToTwoLists_ReturnsTwoSlicesForSingleLineOfNums(test *testing.T) {
	input := "3   4"

	numListOne, numListTwo := ConvertInputToTwoLists(input)

	if numListOne[0] != 3 {
		test.Errorf("Expected: 3, Received: %d", numListOne[0])
	}

	if numListTwo[0] != 4 {
		test.Errorf("Expected: 4, Received: %d", numListTwo[0])
	}
}

func TestConvertInputToTwoLists_ReturnsTwoSlicesForMultipleLinesOfNums(test *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	numListOne, numListTwo := ConvertInputToTwoLists(input)

	if numListOne[0] != 3 {
		test.Errorf("Expected: 3, Received: %d", numListOne[0])
	}

	if numListTwo[0] != 4 {
		test.Errorf("Expected: 4, Received: %d", numListTwo[0])
	}

	if numListOne[1] != 4 {
		test.Errorf("Expected: 3, Received: %d", numListOne[1])
	}

	if numListTwo[1] != 3 {
		test.Errorf("Expected: 4, Received: %d", numListTwo[1])
	}

	if numListOne[2] != 2 {
		test.Errorf("Expected: 2, Received: %d", numListOne[2])
	}

	if numListTwo[2] != 5 {
		test.Errorf("Expected: 5, Received: %d", numListTwo[2])
	}

	if numListOne[5] != 3 {
		test.Errorf("Expected: 3, Received: %d", numListOne[5])
	}

	if numListTwo[5] != 3 {
		test.Errorf("Expected: 3, Received: %d", numListTwo[5])
	}
}

func TestSortBothNumLists_ReturnsTwoSortedSlicesForSingleNumbers(test *testing.T) {
	numListOne := []int{3}
	numListTwo := []int{4}

	sortedListOne, sortedListTwo := SortBothNumLists(numListOne, numListTwo)

	if sortedListOne[0] != 3 {
		test.Errorf("Expected: 3, Received: %d", sortedListOne[0])
	}

	if sortedListTwo[0] != 4 {
		test.Errorf("Expected: 4, Received: %d", sortedListTwo[0])
	}
}

func TestSortBothNumLists_ReturnsTwoSortedSlicesForMultipleNums(test *testing.T) {
	numListOne := []int{3, 4, 2, 1, 3, 3}
	numListTwo := []int{4, 3, 5, 3, 9, 3}

	sortedListOne, sortedListTwo := SortBothNumLists(numListOne, numListTwo)

	if sortedListOne[0] != 1 {
		test.Errorf("Expected: 1, Received: %d", sortedListOne[0])
	}

	if sortedListTwo[0] != 3 {
		test.Errorf("Expected: 3, Received: %d", sortedListTwo[0])
	}

	if sortedListOne[1] != 2 {
		test.Errorf("Expected: 2, Received: %d", sortedListOne[1])
	}

	if sortedListTwo[1] != 3 {
		test.Errorf("Expected: 3, Received: %d", sortedListTwo[1])
	}

	if sortedListOne[2] != 3 {
		test.Errorf("Expected: 3, Received: %d", sortedListOne[2])
	}

	if sortedListTwo[2] != 3 {
		test.Errorf("Expected: 3, Received: %d", sortedListTwo[2])
	}

	if sortedListOne[3] != 3 {
		test.Errorf("Expected: 3, Received: %d", sortedListOne[3])
	}

	if sortedListTwo[3] != 4 {
		test.Errorf("Expected: 4, Received: %d", sortedListTwo[3])
	}

	if sortedListOne[4] != 3 {
		test.Errorf("Expected: 3, Received: %d", sortedListOne[4])
	}

	if sortedListTwo[4] != 5 {
		test.Errorf("Expected: 5, Received: %d", sortedListTwo[4])
	}

	if sortedListOne[5] != 4 {
		test.Errorf("Expected: 4, Received: %d", sortedListOne[5])
	}

	if sortedListTwo[5] != 9 {
		test.Errorf("Expected: 9, Received: %d", sortedListTwo[5])
	}
}

func TestCountDistancesBetweenListNums_ReturnsTotalForSingleList(test *testing.T) {
	numListOne := []int{3}
	numListTwo := []int{4}

	output := CountDistancesBetweenListNums(numListOne, numListTwo)

	if output != 1 {
		test.Errorf("Expected: 1, Received: %d", output)
	}
}

func TestCountDistancesBetweenListNums_ReturnsTotalForMultipleList(test *testing.T) {
	numListOne := []int{1, 2, 3, 3, 3, 4}
	numListTwo := []int{3, 3, 3, 4, 5, 9}

	output := CountDistancesBetweenListNums(numListOne, numListTwo)

	if output != 11 {
		test.Errorf("Expected: 11, Received: %d", output)
	}
}

func TestPart1_ReturnsZeroForEmptyInput(test *testing.T) {
	input := ""

	output := Part1(input)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)
	}
}

func TestPart1_ReturnsZeroForSingleNumsList(test *testing.T) {
	input := "3   4"

	output := Part1(input)

	if output != 1 {
		test.Errorf("Expected: 1, Received: %d", output)
	}
}

func TestPart1_ReturnsZeroForMultipleNumsLists(test *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	output := Part1(input)

	if output != 11 {
		test.Errorf("Expected: 11, Received: %d", output)
	}
}

func TestCountRightHandListNumFrequency_ReturnsMapOfNumFrequencies_SingleNumList(test *testing.T) {
	input := []int{3}

	output := CountRightHandListNumFrequency(input)

	valueToCheck, exists := output[3]

	if !exists && valueToCheck != 1 {
		test.Errorf("Expected: 1, Received: %d", valueToCheck)
	}
}

func TestCountRightHandListNumFrequency_ReturnsMapOfNumFrequencies_MultipleNums(test *testing.T) {
	input := []int{3, 3, 3}

	output := CountRightHandListNumFrequency(input)

	valueToCheck, exists := output[3]

	if !exists && valueToCheck != 3 {
		test.Errorf("Expected: 3, Received: %d", valueToCheck)
	}

	input = []int{3, 3, 3, 4, 5, 9}

	output = CountRightHandListNumFrequency(input)

	valueToCheck, exists = output[3]

	if !exists && valueToCheck != 3 {
		test.Errorf("Expected: 3, Received: %d", valueToCheck)
	}

	valueToCheck, exists = output[4]

	if !exists && valueToCheck != 1 {
		test.Errorf("Expected: 4, Received: %d", valueToCheck)
	}

	valueToCheck, exists = output[5]

	if !exists && valueToCheck != 1 {
		test.Errorf("Expected: 1, Received: %d", valueToCheck)
	}

	valueToCheck, exists = output[9]

	if !exists && valueToCheck != 1 {
		test.Errorf("Expected: 1, Received: %d", valueToCheck)
	}
}

func TestCountSimiliarityScore_ReturnsTotalForSingleNum_ZeroFrequency(test *testing.T) {
	numList := []int{3}

	numFrequencyMap := make(map[int]int)

	output := CountSimiliarityScore(numList, numFrequencyMap)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)
	}
}

func TestCountSimiliarityScore_ReturnsTotalForSingleNum_SingleFrequency(test *testing.T) {
	numList := []int{3}

	numFrequencyMap := make(map[int]int)
	numFrequencyMap[3] = 1

	output := CountSimiliarityScore(numList, numFrequencyMap)

	if output != 3 {
		test.Errorf("Expected: 3, Received: %d", output)
	}

	numList = []int{3}

	numFrequencyMap = make(map[int]int)
	numFrequencyMap[3] = 3

	output = CountSimiliarityScore(numList, numFrequencyMap)

	if output != 9 {
		test.Errorf("Expected: 9, Received: %d", output)
	}

	numList = []int{4}

	numFrequencyMap = make(map[int]int)
	numFrequencyMap[4] = 1

	output = CountSimiliarityScore(numList, numFrequencyMap)

	if output != 4 {
		test.Errorf("Expected: 4, Received: %d", output)
	}

	numList = []int{2}

	numFrequencyMap = make(map[int]int)

	output = CountSimiliarityScore(numList, numFrequencyMap)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)
	}

	numList = []int{1}

	numFrequencyMap = make(map[int]int)

	output = CountSimiliarityScore(numList, numFrequencyMap)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)
	}
}

func TestCountSimiliarityScore_ReturnsTotalForMultiNumsAndFrequencies(test *testing.T) {
	numList := []int{3, 4, 2, 1, 3, 3}

	numFrequencyMap := make(map[int]int)
	numFrequencyMap[3] = 3
	numFrequencyMap[4] = 1

	output := CountSimiliarityScore(numList, numFrequencyMap)

	if output != 31 {
		test.Errorf("Expected: 31, Received: %d", output)
	}
}

func TestPart2_CorrectTotalReturned_SingleNumList(test *testing.T) {
	input := "3   4"

	output := Part2(input)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)
	}

	input = "3   3"

	output = Part2(input)

	if output != 3 {
		test.Errorf("Expected: 3, Received: %d", output)
	}
}

func TestPart2_CorrectTotalReturned_MultipleNumsList(test *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	output := Part2(input)

	if output != 31 {
		test.Errorf("Expected: 31, Received: %d", output)
	}
}