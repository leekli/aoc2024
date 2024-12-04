package main

import "testing"

func TestStringTo2DArray_ProducesCorrect2DNumArray(test *testing.T) {
	input := `MMMS
MSAM
AMXS
MSAM`

	output := StringTo2DArray(input)

	if output[0][0] != "M" {
		test.Errorf("Expected: 'M', Received: %s", output[0][0])
	}

	if output[1][0] != "M" {
		test.Errorf("Expected: 'M', Received: %s", output[1][0])
	}

	if output[2][0] != "A" {
		test.Errorf("Expected: 'A', Received: %s", output[1][0])
	}

	if output[0][3] != "S" {
		test.Errorf("Expected: 'S', Received: %s", output[0][3])
	}

	if output[2][2] != "X" {
		test.Errorf("Expected: 'X', Received: %s", output[2][2])
	}
}

func TestIsXmasFound_ReturnsFalseForNoXmasFound(test *testing.T) {
	a := "A"
	b := "B"
	c := "C"
	d := "D"

	output := IsXmasFound(a, b, c, d)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	a = "X"
	b = "M"
	c = "A"
	d = "Z"

	output = IsXmasFound(a, b, c, d)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}
} 

func TestIsXmasFound_ReturnsTrueForXmasFound(test *testing.T) {
	a := "X"
	b := "M"
	c := "A"
	d := "S"

	output := IsXmasFound(a, b, c, d)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}
} 

func TestCanLookHorizontally_ReturnsFalseForGridNotBigEnough(test *testing.T) {
	input := [][]string{{}}

	output := CanLookHorizontally(input, 0, 0)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	input = [][]string{{"A", "A"}, {"B", "B"}, {"C", "C"}, {"D", "D"}}

	output = CanLookHorizontally(input, 0, 0)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	input = [][]string{{"A", "A", "A"}, {"B", "B", "B"}, {"C", "C", "C"}, {"D", "D", "C"}}

	output = CanLookHorizontally(input, 0, 0)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	input = [][]string{{"A", "A", "A", "A"}, {"B", "B", "B", "B"}, {"C", "C", "C", "C"}, {"D", "D", "D", "D"}}

	output = CanLookHorizontally(input, 0, 1)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	output = CanLookHorizontally(input, 0, 2)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	output = CanLookHorizontally(input, 0, 3)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}
} 

func TestCanLookHorizontally_ReturnsTrueForValidArrayIndex(test *testing.T) {
	input := [][]string{{"A", "A", "A", "A"}, {"B", "B", "B", "B"}, {"C", "C", "C", "C"}, {"D", "D", "D", "D"}}

	output := CanLookHorizontally(input, 0, 0)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	output = CanLookHorizontally(input, 1, 0)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	output = CanLookHorizontally(input, 2, 0)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	output = CanLookHorizontally(input, 3, 0)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}
} 

func TestCanLookVertically_ReturnsFalseForGridNotBigEnough(test *testing.T) {
	input := [][]string{{}}

	output := CanLookVertically(input, 0, 0)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	input = [][]string{{"A", "A"}, {"B", "B"}, {"C", "C"}, {"D", "D"}}

	output = CanLookVertically(input, 1, 0)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	output = CanLookVertically(input, 2, 0)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}
} 

func TestCanLookVertically_ReturnsTrueForValidArrayIndex(test *testing.T) {
	input := [][]string{{"A", "A"}, {"B", "B"}, {"C", "C"}, {"D", "D"}}

	output := CanLookVertically(input, 3, 0)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	output = CanLookVertically(input, 3, 3)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	output = CanLookVertically(input, 3, 2)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	output = CanLookVertically(input, 3, 1)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	input = [][]string{{"A", "A", "A", "A", "A"}, {"A", "A", "A", "A", "A"}, {"A", "A", "A", "A", "A"}, {"A", "A", "A", "A", "A"}, {"A", "A", "A", "A", "A"}}

	output = CanLookVertically(input, 4, 3)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}
} 

func TestCanLookBackwards_ReturnsFalseForGridNotBigEnough(test *testing.T) {
	input := [][]string{{}}

	output := CanLookBackwards(input, 0, 0)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	input = [][]string{{"A", "A"}, {"B", "B"}, {"C", "C"}, {"D", "D"}}

	output = CanLookBackwards(input, 0, 0)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	output = CanLookBackwards(input, 0, 1)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	input = [][]string{{"A", "A", "A"}, {"B", "B", "B"}, {"C", "C", "C"}, {"D", "D", "C"}}

	output = CanLookBackwards(input, 0, 2)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	input = [][]string{{"A", "A", "A", "A"}, {"B", "B", "B", "B"}, {"C", "C", "C", "C"}, {"D", "D", "D", "D"}}

	output = CanLookBackwards(input, 1, 2)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}
} 

func TestCanLookBackwards_ReturnsTrueForValidArrayIndex(test *testing.T) {
	input := [][]string{{"A", "A", "A", "A"}, {"B", "B", "B", "B"}, {"C", "C", "C", "C"}, {"D", "D", "D", "D"}}

	output := CanLookBackwards(input, 0, 3)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	output = CanLookBackwards(input, 1, 3)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	output = CanLookBackwards(input, 2, 3)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	output = CanLookBackwards(input, 3, 3)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}
} 

func TestCanLookDiagonallyTopLeft_ReturnsTrueFalseForIfInBounds(test *testing.T) {
	input := [][]string{{"A", "A", "A", "A"}, {"B", "B", "B", "B"}, {"C", "C", "C", "C"}, {"D", "D", "D", "D"}}

	output := CanLookDiagonallyTopLeft(input, 3, 3)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	output = CanLookDiagonallyTopLeft(input, 2, 3)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	output = CanLookDiagonallyTopLeft(input, 3, 0)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}
}

func TestCanLookDiagonallyTopRight_ReturnsTrueFalseForIfInBounds(test *testing.T) {
	input := [][]string{{"A", "A", "A", "A"}, {"B", "B", "B", "B"}, {"C", "C", "C", "C"}, {"D", "D", "D", "D"}}

	output := CanLookDiagonallyTopRight(input, 3, 0)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	output = CanLookDiagonallyTopRight(input, 3, 1)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	output = CanLookDiagonallyTopRight(input, 3, 2)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	output = CanLookDiagonallyTopRight(input, 2, 0)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}
}

func TestCanLookDiagonallyBottomLeft_ReturnsTrueFalseForIfInBounds(test *testing.T) {
	input := [][]string{{"A", "A", "A", "A"}, {"B", "B", "B", "B"}, {"C", "C", "C", "C"}, {"D", "D", "D", "D"}}

	output := CanLookDiagonallyBottomLeft(input, 0, 3)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	output = CanLookDiagonallyBottomLeft(input, 0, 1)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	output = CanLookDiagonallyBottomLeft(input, 0, 2)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	output = CanLookDiagonallyBottomLeft(input, 2, 0)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}
}

func TestCanLookDiagonallyBottomRight_ReturnsTrueFalseForIfInBounds(test *testing.T) {
	input := [][]string{{"A", "A", "A", "A"}, {"B", "B", "B", "B"}, {"C", "C", "C", "C"}, {"D", "D", "D", "D"}}

	output := CanLookDiagonallyBottomRight(input, 0, 0)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	output = CanLookDiagonallyBottomRight(input, 0, 1)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	output = CanLookDiagonallyBottomRight(input, 0, 2)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	output = CanLookDiagonallyBottomRight(input, 2, 0)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}
}

func TestDoWordSearch_CountXmasFound_ReturnsTotalForXmasFound(test *testing.T) {
	input := `..X...
.SAMX.
.A..A.
XMAS.S
.X....`

	wordSearchGrid := StringTo2DArray(input)

	output := DoWordSearch_CountXmasFound(wordSearchGrid)

	if output != 4 {
		test.Errorf("Expected: 4, Received: %d", output)
	}

	input = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	wordSearchGrid = StringTo2DArray(input)

	output = DoWordSearch_CountXmasFound(wordSearchGrid)

	if output != 18 {
		test.Errorf("Expected: 18, Received: %d", output)
	}
}