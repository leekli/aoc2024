package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	// Data setup
	rawInput, err := readFileToString("input.txt")

	if err != nil {
		os.Exit(-1)
	}

	// Part 1
	p1Output := Part1(rawInput)

	fmt.Printf("Day 8, Part 1 Output: %d\n", p1Output)

	// Part 2
	//p2Output := Part2(rawInput)

	//fmt.Printf("Day 8, Part 2 Output: %d\n", p2Output)	
}

func Part1(input string) int {
	total := 0

	mainMap := StringTo2DArray(input)
	antiNodeMap := CreateAntiNodeMap(mainMap)

	antennaLocations := BuildAntennaLocationsDict(mainMap)

	// Go through the antenna locations dict
	// on the first key ('a' in example case):
		// do a while here... while current key length is not 0?
	// --> get the [0] location and save in var, use this as the base
	// --> start a loop from current key, from index 1 until last one
	// --> on first index of loop (index 1 in first case) get hold of source co-ordinates (index 0 we have from before loop), and get hold of current index co-ordinates (index 1 in first case)
	// --> find the XY difference between them 
	// --> now, for the first co-ords (index 0 in loop iteration 1) we need pretend go backwards using the XY difference: use .GetLocationIfAntennaGoesBackwards() once we have those co-ords, check they're valid with .IsNewRowAndColWithinBounds(), and if they are... add them to the anti node map .AddAntiNodeToMap()
	// --> now, do the same with the second co-ords (index 1), pretend go forwards using difference xy, CHECK THEY'RE VALID BOUNDS, and add to anti node map if they are
		// finally STILL IN WHILE LOOP: slice the current key from [1:] to get rid of the source index we've just tried on

	for key, _ := range antennaLocations {
		for len(antennaLocations[key]) != 0 {
			firstLoc := antennaLocations[key][0]

			for i := 1; i < len(antennaLocations[key]); i++ {
				secondLoc := antennaLocations[key][i]

				difference := GetXYDifferenceBetweenTwoAntennas(firstLoc, secondLoc)

				locIfGoingBack := GetLocationIfAntennaGoesBackwards(firstLoc, difference)

				isBackWithinBounds := IsNewRowAndColWithinBounds(mainMap, locIfGoingBack[0], locIfGoingBack[1])

				if isBackWithinBounds {
					AddAntiNodeToMap(antiNodeMap, locIfGoingBack[0], locIfGoingBack[1])
				}

				locIfGoingForward := GetLocationIfAntennaGoesForwards(secondLoc, difference)

				isForwardWithinBounds := IsNewRowAndColWithinBounds(mainMap, locIfGoingForward[0], locIfGoingForward[1])

				if isForwardWithinBounds {
					AddAntiNodeToMap(antiNodeMap, locIfGoingForward[0], locIfGoingForward[1])
				}
			}

			antennaLocations[key] = antennaLocations[key][1:]
		}
	}

	// Once all anti-nodes located, count how many we have
	total = CountAllAntiNodeLocations(antiNodeMap)

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

    var fullMap [][]string

    for _, lineRow := range lineRows {
        splitChars := strings.Split(lineRow, "")

        var finalChars []string

        for _, numStr := range splitChars {

            finalChars = append(finalChars, numStr)
        }

        fullMap = append(fullMap, finalChars)
    }

    return fullMap
}

func CreateAntiNodeMap(theMap [][]string) [][]int {
	rows := len(theMap)
	cols := len(theMap[0])
 
	antiNodeMap := make([][]int, rows)
	
	for i := range antiNodeMap {
		antiNodeMap[i] = make([]int, cols)
	}
 
	return antiNodeMap
 }
 
 func AddAntiNodeToMap(antiNodeMap [][]int, row int, col int) {
	if (antiNodeMap[row][col] == 0) {
		antiNodeMap[row][col] = 1
	}
 }
 
 func CountAllAntiNodeLocations(antiNodeMap [][]int) int {
	 totalAntinodeLocations := 0
 
	 for i := 0; i < len(antiNodeMap); i++ {
		 for j := 0; j < len(antiNodeMap[i]); j++ {
			 if antiNodeMap[i][j] == 1 {
				totalAntinodeLocations++
			 }
		 }
	 }
 
	 return totalAntinodeLocations
 }

 func IsNewRowAndColWithinBounds(theMap [][]string, newRow int, newCol int) bool {
	isWithinBounds := false

	if newRow >= 0 && newRow < len(theMap) && newCol >= 0 && newCol < len(theMap[0]) {
		isWithinBounds = true	
	}

	return isWithinBounds
 }

 func BuildAntennaLocationsDict(theMap [][]string) map[string][][2]int {
	antennasList := make(map[string][][2]int)

	for i := 0; i < len(theMap); i++ {
        for j := 0; j < len(theMap[i]); j++ {
            currentElement := theMap[i][j]

			charAsRune := []rune(theMap[i][j])

            if unicode.IsLetter(charAsRune[0]) || unicode.IsDigit(charAsRune[0]) {
				if currentElement != "." {
		                antennasList[currentElement] = append(antennasList[currentElement], [2]int{i, j})			
				}
            }
        }
	}

	// Remove any antennas which only have 1 set of co-ordinates, this means there's only one of them and isn't any more, therefore violating the puzzle rule 'In particular, an antinode occurs at any point that is perfectly in line with two antennas of the same frequency'
	for key, _ := range antennasList {
        if len(antennasList[key]) == 1 {
			delete(antennasList, key)
		}
    }

	return antennasList
 }

 func GetXYDifferenceBetweenTwoAntennas(c1 [2]int, c2 [2]int) [2]int {
	difference := [2]int{}

	num1 := c2[0] - c1[0]
	num2 := c2[1] - c1[1]

	difference = [2]int{num1, num2}

	return difference
 }

 func GetLocationIfAntennaGoesBackwards(coOrds [2]int, differenceCoOrds[2]int) [2]int {
	newLocation := [2]int{}

	num1Diff := coOrds[0] - differenceCoOrds[0]
	num2Diff := coOrds[1] - differenceCoOrds[1]

	newLocation = [2]int{num1Diff, num2Diff}

	return newLocation
 }

 func GetLocationIfAntennaGoesForwards(coOrds [2]int, differenceCoOrds[2]int) [2]int {
	newLocation := [2]int{}

	num1Diff := coOrds[0] + differenceCoOrds[0]
	num2Diff := coOrds[1] + differenceCoOrds[1]

	newLocation = [2]int{num1Diff, num2Diff}

	return newLocation
 }