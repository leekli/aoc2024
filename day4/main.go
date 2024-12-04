package main

import (
	"fmt"
	"os"
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

	fmt.Printf("Day 4, Part 1 Output: %d\n", p1Output)

	// Part 2
	//p2Output := Part2(rawInput)

	//fmt.Printf("Day X, Part 2 Output: %d\n", p2Output)	
}

func Part1(input string) int {
	total := 0

	wordSearchGrid := StringTo2DArray(input)

	total = DoWordSearch_CountXmasFound(wordSearchGrid)

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

func StringTo2DArray(input string) [][]string {
    lineRows := strings.Split(input, "\n")

    var wordSearchArray [][]string

    for _, lineRow := range lineRows {
        splitChars := strings.Split(lineRow, "")

        var finalChars []string

        for _, numStr := range splitChars {

            finalChars = append(finalChars, numStr)
        }

        wordSearchArray = append(wordSearchArray, finalChars)
    }

    return wordSearchArray
}

func IsXmasFound(strA string, strB string, strC string, strD string) bool {
	xmasFound := false

	if strA == "X" && strB == "M" && strC == "A" && strD == "S" {
		xmasFound = true
	}

	return xmasFound
}

func CanLookHorizontally(grid [][]string, startingRow int, startingCol int) bool {
	canLookHorizontal := false

	if len(grid) == 0 {
		return canLookHorizontal
	}

	indexToCheck := startingCol + 3

	if indexToCheck >= 0 && indexToCheck < len(grid[startingRow]) {
		canLookHorizontal = true
	}

	return canLookHorizontal
}

func CanLookVertically(grid [][]string, startingRow int, startingCol int) bool {
	// canLookVertical := false

	// if len(grid) == 0 {
	// 	return canLookVertical
	// }

	// indexToCheck := startingRow - 3

	// if indexToCheck >= 0 && indexToCheck < len(grid) {
	// 	canLookVertical = true
	// }

	// return canLookVertical

	canLookVertical := false

	if len(grid) == 0 {
		return canLookVertical
	}

	indexToCheckUp := startingRow - 3
	indexToCheckDown := startingRow + 3

	if (indexToCheckUp >= 0 && indexToCheckUp < len(grid)) || (indexToCheckDown >= 0 && indexToCheckDown < len(grid)) {
		canLookVertical = true
	}

	return canLookVertical
}

func CanLookBackwards(grid [][]string, startingRow int, startingCol int) bool {
	canLookBackward := false

	if len(grid) == 0 {
		return canLookBackward
	}

	indexToCheck := startingCol - 3

	if indexToCheck >= 0 && indexToCheck < len(grid[startingRow]) {
		canLookBackward = true
	}

	return canLookBackward
}

func CanLookDiagonallyTopLeft(grid [][]string, startingRow int, startingCol int) bool {
	canLookDiagonal := false

	if len(grid) == 0 {
		return canLookDiagonal
	}

	rowToCheck := startingRow - 3
	colToCheck := startingCol - 3

	if rowToCheck >= 0 && rowToCheck < len(grid) && colToCheck >= 0 && colToCheck < len(grid[startingRow]) {
		canLookDiagonal = true
	}

	return canLookDiagonal
}

func CanLookDiagonallyTopRight(grid [][]string, startingRow int, startingCol int) bool {
	canLookDiagonal := false

	if len(grid) == 0 {
		return canLookDiagonal
	}

	rowToCheck := startingRow - 3
	colToCheck := startingCol + 3

	if rowToCheck >= 0 && rowToCheck < len(grid) && colToCheck >= 0 && colToCheck < len(grid[startingRow]) {
		canLookDiagonal = true
	}

	return canLookDiagonal
}

func CanLookDiagonallyBottomLeft(grid [][]string, startingRow int, startingCol int) bool {
	canLookDiagonal := false

	if len(grid) == 0 {
		return canLookDiagonal
	}

	rowToCheck := startingRow + 3
	colToCheck := startingCol - 3

	if rowToCheck >= 0 && rowToCheck < len(grid) && colToCheck >= 0 && colToCheck < len(grid[startingRow]) {
		canLookDiagonal = true
	}

	return canLookDiagonal
}

func CanLookDiagonallyBottomRight(grid [][]string, startingRow int, startingCol int) bool {
	canLookDiagonal := false

	if len(grid) == 0 {
		return canLookDiagonal
	}

	rowToCheck := startingRow + 3
	colToCheck := startingCol + 3

	if rowToCheck >= 0 && rowToCheck < len(grid) && colToCheck >= 0 && colToCheck < len(grid[startingRow]) {
		canLookDiagonal = true
	}

	return canLookDiagonal
}

func DoWordSearch_CountXmasFound(wordSearchGrid [][]string) int {
	total := 0

	for i := 0; i < len(wordSearchGrid); i++ {
		for j := 0; j < len(wordSearchGrid[i]); j++ {
			
			// Check horizontally
			if CanLookHorizontally(wordSearchGrid, i, j) {
				strA := wordSearchGrid[i][j]
				strB := wordSearchGrid[i][j + 1]
				strC := wordSearchGrid[i][j + 2]
				strD := wordSearchGrid[i][j + 3]

				if IsXmasFound(strA, strB, strC, strD) {
					total++
				}
			}

			// Check vertically upwards
			if CanLookVertically(wordSearchGrid, i, j) && i - 3 >= 0 {
				strA := wordSearchGrid[i][j]
				strB := wordSearchGrid[i - 1][j]
				strC := wordSearchGrid[i - 2][j]
				strD := wordSearchGrid[i - 3][j]

				if IsXmasFound(strA, strB, strC, strD) {
					total++
				}
			}

			// Check vertically downwards
			if CanLookVertically(wordSearchGrid, i, j) && i + 3 < len(wordSearchGrid) {
				strA := wordSearchGrid[i][j]
				strB := wordSearchGrid[i + 1][j]
				strC := wordSearchGrid[i + 2][j]
				strD := wordSearchGrid[i + 3][j]

				if IsXmasFound(strA, strB, strC, strD) {
					total++
				}
			}

			// Check backwards
			if CanLookBackwards(wordSearchGrid, i, j) {
				strA := wordSearchGrid[i][j]
				strB := wordSearchGrid[i][j - 1]
				strC := wordSearchGrid[i][j - 2]
				strD := wordSearchGrid[i][j - 3]

				if IsXmasFound(strA, strB, strC, strD) {
					total++
				}
			}

			// Check diagonally top-left
			if CanLookDiagonallyTopLeft(wordSearchGrid, i, j) {
				strA := wordSearchGrid[i][j]
				strB := wordSearchGrid[i - 1][j - 1]
				strC := wordSearchGrid[i - 2][j - 2]
				strD := wordSearchGrid[i - 3][j - 3]

				if IsXmasFound(strA, strB, strC, strD) {
					total++
				}
			}

			// Check diagonally top-right
			if CanLookDiagonallyTopRight(wordSearchGrid, i, j) {
				strA := wordSearchGrid[i][j]
				strB := wordSearchGrid[i - 1][j + 1]
				strC := wordSearchGrid[i - 2][j + 2]
				strD := wordSearchGrid[i - 3][j + 3]

				if IsXmasFound(strA, strB, strC, strD) {
					total++
				}
			}

			// Check diagonally bottom-left
			if CanLookDiagonallyBottomLeft(wordSearchGrid, i, j) {
				strA := wordSearchGrid[i][j]
				strB := wordSearchGrid[i + 1][j - 1]
				strC := wordSearchGrid[i + 2][j - 2]
				strD := wordSearchGrid[i + 3][j - 3]

				if IsXmasFound(strA, strB, strC, strD) {
					total++
				}
			}

			// Check diagonally bottom-right
			if CanLookDiagonallyBottomRight(wordSearchGrid, i, j) {
				strA := wordSearchGrid[i][j]
				strB := wordSearchGrid[i + 1][j + 1]
				strC := wordSearchGrid[i + 2][j + 2]
				strD := wordSearchGrid[i + 3][j + 3]

				if IsXmasFound(strA, strB, strC, strD) {
					total++
				}
			}
		}
	}

	return total
}