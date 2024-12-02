package main

import (
	"testing"
)

func TestStringTo2DArray_ProducesCorrect2DNumArray(test *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	output := StringTo2DArray(input)

	if output[0][0] != 7 {
		test.Errorf("Expected: 7, Received: %d", output)
	}

	if output[1][2] != 7 {
		test.Errorf("Expected: 7, Received: %d", output)
	}

	if output[2][3] != 2 {
		test.Errorf("Expected: 2, Received: %d", output)
	}

	if output[5][1] != 3 {
		test.Errorf("Expected: 3, Received: %d", output)
	}
}

func TestIsNumChangingWithinRange_ReturnsFalseForZero(test *testing.T) {
	num1 := 0
	num2 := 0

	output := IsNumChangingWithinRange(num1, num2)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}
}

func TestIsNumChangingWithinRange_ReturnsFalseForNotWithinRange(test *testing.T) {
	num1 := 2
	num2 := 7

	output := IsNumChangingWithinRange(num1, num2)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	num1 = 6
	num2 = 2

	output = IsNumChangingWithinRange(num1, num2)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	num1 = 4
	num2 = 4

	output = IsNumChangingWithinRange(num1, num2)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}
}

func TestIsNumChangingWithinRange_ReturnsTrueForWithinRange(test *testing.T) {
	num1 := 7
	num2 := 6

	output := IsNumChangingWithinRange(num1, num2)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	num1 = 4
	num2 = 2

	output = IsNumChangingWithinRange(num1, num2)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	num1 = 3
	num2 = 6

	output = IsNumChangingWithinRange(num1, num2)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}
}

func TestAreAllNumsIncreasingOrDecreasing_ReturnsFalseForNotAllIncreasing(test *testing.T) {
	input := []int{1, 3, 2, 4, 5}

	output := AreAllNumsIncreasingOrDecreasing(input)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}
}

func TestAreAllNumsIncreasingOrDecreasing_ReturnsFalseForOneSameNumber(test *testing.T) {
	input := []int{8, 6, 4, 4, 1}

	output := AreAllNumsIncreasingOrDecreasing(input)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}
}

func TestAreAllNumsIncreasingOrDecreasing_ReturnsTrueForAllDecreasing(test *testing.T) {
	input := []int{7, 6, 4, 2, 1}

	output := AreAllNumsIncreasingOrDecreasing(input)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	input = []int{9, 7, 6, 2, 1}

	output = AreAllNumsIncreasingOrDecreasing(input)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}
}

func TestAreAllNumsIncreasingOrDecreasing_ReturnsTrueForAllIncreasing(test *testing.T) {
	input := []int{1, 2, 7, 8, 9}

	output := AreAllNumsIncreasingOrDecreasing(input)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	input = []int{1, 3, 6, 7, 9}

	output = AreAllNumsIncreasingOrDecreasing(input)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}
}

func TestIsLevelSafe_ReturnsSafeForLevel(test *testing.T) {
	input := []int{7, 6, 4, 2, 1}

	output := IsLevelSafe(input)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	input = []int{1, 3, 6, 7, 9}

	output = IsLevelSafe(input)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}
}

func TestIsLevelSafe_ReturnsNotSafeForLevel(test *testing.T) {
	input := []int{1, 2, 7, 8, 9}

	output := IsLevelSafe(input)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	input = []int{9, 7, 6, 2, 1}

	output = IsLevelSafe(input)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	input = []int{1, 3, 2, 4, 5}

	output = IsLevelSafe(input)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	input = []int{8, 6, 4, 4, 1}

	output = IsLevelSafe(input)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}
}

func TestCountTotalSafeLevels_ReturnsZeroForNoSafeLevels(test *testing.T) {
	input := [][]int{{1, 2, 1, 8, 9}, {9, 8, 9, 2, 3}}

	output := CountTotalSafeLevels(input)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)
	}
}

func TestCountTotalSafeLevels_ReturnsTotalSafeLevels(test *testing.T) {
	input := [][]int{{7, 6, 4, 2, 1}}

	output := CountTotalSafeLevels(input)

	if output != 1 {
		test.Errorf("Expected: 1, Received: %d", output)
	}

	input = [][]int{{7, 6, 4, 2, 1}, {1, 2, 7, 8, 9}}

	output = CountTotalSafeLevels(input)

	if output != 1 {
		test.Errorf("Expected: 1, Received: %d", output)
	}

	input = [][]int{{7, 6, 4, 2, 1}, {1, 2, 7, 8, 9}, {9, 7, 6, 2, 1}, {1, 3, 2, 4, 5}, {8, 6, 4, 4, 1}, {1, 3, 6, 7, 9}}

	output = CountTotalSafeLevels(input)

	if output != 2 {
		test.Errorf("Expected: 2, Received: %d", output)
	}
}

func TestCountTotalSafeLevels_Part2_ReturnsZeroForNoLevelsAbleToRemove(test *testing.T) {
	input := [][]int{{1, 2, 7, 8, 9}}

	output := CountTotalSafeLevels_Part2(input)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)
	}

	input = [][]int{{9, 7, 6, 2, 1}}

	output = CountTotalSafeLevels_Part2(input)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)
	}
}

func TestCountTotalSafeLevels_Part2_ReturnsSafeTotalForLevelsAbleToRemove_Single(test *testing.T) {
	input := [][]int{{7, 6, 4, 2, 1}}

	output := CountTotalSafeLevels_Part2(input)

	if output != 1 {
		test.Errorf("Expected: 1, Received: %d", output)
	}

	input = [][]int{{1, 3, 6, 7, 9}}

	output = CountTotalSafeLevels_Part2(input)

	if output != 1 {
		test.Errorf("Expected: 1, Received: %d", output)
	}

	input = [][]int{{1, 3, 2, 4, 5}}

	output = CountTotalSafeLevels_Part2(input)

	if output != 1 {
		test.Errorf("Expected: 1, Received: %d", output)
	}

	input = [][]int{{8, 6, 4, 4, 1}}

	output = CountTotalSafeLevels_Part2(input)

	if output != 1 {
		test.Errorf("Expected: 1, Received: %d", output)
	}
}

func TestCountTotalSafeLevels_Part2_ReturnsSafeTotalForLevelsAbleToRemove_Multiple(test *testing.T) {
	input := [][]int{{7, 6, 4, 2, 1}, {1, 2, 7, 8, 9}, {9, 7, 6, 2, 1}, {1, 3, 2, 4, 5}, {8, 6, 4, 4, 1}, {1, 3, 6, 7, 9}}

	output := CountTotalSafeLevels_Part2(input)

	if output != 4 {
		test.Errorf("Expected: 4, Received: %d", output)
	}
}