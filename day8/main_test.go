package main

import (
	"testing"
)

func generateTestMap() [][]string {
	input := `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

	testMap := StringTo2DArray(input)

	return testMap
}

func TestFuncName_TestBehaviourDesc(test *testing.T) {
	// input := ""

	output := 0 // func call here

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)
	}
}

func TestStringTo2DArray_Produces2DArrayFromInput(test *testing.T) {
	input := `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

	output := StringTo2DArray(input)

	if len(output) != 12 {
		test.Errorf("Expected: 12, Received: %d", len(output))
	}

	if output[0][0] != "." {
		test.Errorf("Expected: '.', Received: %s", output[0][0])
	}

	if output[1][8] != "0" {
		test.Errorf("Expected: '0', Received: %s", output[1][8])
	}

	if output[8][8] != "A" {
		test.Errorf("Expected: 'A', Received: %s", output[8][8])
	}
}

func TestCreateAntiNodeMap_ReturnsArraySameSizeAsMapAllSetToZero(test *testing.T) {
	testMap := generateTestMap()

	output := CreateAntiNodeMap(testMap)

	if len(output) != 12 {
		test.Errorf("Expected: 12, Received: %d", len(output))	
	}

	if len(output[0]) != 12 {
		test.Errorf("Expected: 12, Received: %d", len(output))	
	}

	if output[0][0] != 0 {
		test.Errorf("Expected: 0, Received: %d", output[0][0])	
	}

	if output[2][3] != 0 {
		test.Errorf("Expected: 0, Received: %d", output[2][3])	
	}

	if output[4][8] != 0 {
		test.Errorf("Expected: 0, Received: %d", output[4][8])	
	}

	if output[9][1] != 0 {
		test.Errorf("Expected: 0, Received: %d", output[9][1])	
	}

	if output[9][9] != 0 {
		test.Errorf("Expected: 0, Received: %d", output[9][9])	
	}
}

func TestAddAntiNodeToMap_UpdatesArrayWithAntiNode(test *testing.T) {
	testMap := generateTestMap()

	antiNodeMap := CreateAntiNodeMap(testMap)

	AddAntiNodeToMap(antiNodeMap, 2, 3)

	if antiNodeMap[2][3] != 1 {
		test.Errorf("Expected: 1, Received: %d", antiNodeMap[2][3])	
	}

	AddAntiNodeToMap(antiNodeMap, 4, 6)

	if antiNodeMap[4][6] != 1 {
		test.Errorf("Expected: 1, Received: %d", antiNodeMap[4][6])	
	}

	AddAntiNodeToMap(antiNodeMap, 2, 3)

	if antiNodeMap[0][9] != 0 {
		test.Errorf("Expected: 0, Received: %d", antiNodeMap[0][9])	
	}
}

func TestCountAllAntiNodeLocations_ReturnsTotalForZero(test *testing.T) {
	testMap := generateTestMap()

	antiNodeMap := CreateAntiNodeMap(testMap)

	output := CountAllAntiNodeLocations(antiNodeMap)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)	
	}	
}

func TestCountAllAntiNodeLocations_ReturnsTotalForOneOrMore(test *testing.T) {
	testMap := generateTestMap()

	antiNodeMap := CreateAntiNodeMap(testMap)

	AddAntiNodeToMap(antiNodeMap, 0, 1)
	AddAntiNodeToMap(antiNodeMap, 1, 2)
	AddAntiNodeToMap(antiNodeMap, 1, 2)
	AddAntiNodeToMap(antiNodeMap, 2, 3)
	AddAntiNodeToMap(antiNodeMap, 4, 5)

	output := CountAllAntiNodeLocations(antiNodeMap)

	if output != 4 {
		test.Errorf("Expected: 4, Received: %d", output)	
	}	
}

func TestIsNewRowAndColWithinBounds_ReturnsTrueForWithinBounds(test *testing.T) {
	testMap := generateTestMap()

	testRow := 0
	testCol := 0

	output := IsNewRowAndColWithinBounds(testMap, testRow, testCol)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)			
	}

	testRow = 2
	testCol = 3

	output = IsNewRowAndColWithinBounds(testMap, testRow, testCol)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)			
	}

	testRow = 11
	testCol = 11

	output = IsNewRowAndColWithinBounds(testMap, testRow, testCol)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)			
	}
}

func TestIsNewRowAndColWithinBounds_ReturnsFalseForOutOfBounds(test *testing.T) {
	testMap := generateTestMap()

	testRow := 12
	testCol := 12

	output := IsNewRowAndColWithinBounds(testMap, testRow, testCol)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)			
	}

	testRow = 0
	testCol = 12

	output = IsNewRowAndColWithinBounds(testMap, testRow, testCol)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)			
	}

	testRow = 5
	testCol = 13

	output = IsNewRowAndColWithinBounds(testMap, testRow, testCol)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)			
	}

	testRow = 6
	testCol = -1

	output = IsNewRowAndColWithinBounds(testMap, testRow, testCol)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)			
	}

	testRow = 13
	testCol = 0

	output = IsNewRowAndColWithinBounds(testMap, testRow, testCol)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)			
	}
}

func TestBuildAntennaLocationsDict_ReturnsDictOfAntennaKeysAndLocationValues(test *testing.T) {
	input := `..........
..........
..........
....a.....
........a.
.....a....
..........
..........
..........
..........`

	testMap := StringTo2DArray(input)

	output := BuildAntennaLocationsDict(testMap)

	if _, exists := output["a"]; exists {
		if len(output["a"]) != 3 {
			test.Errorf("Expected: 2, Received: %d", len(output["3"]))			
		}

		if output["a"][0][0] != 3 {
			test.Errorf("Expected: 3, Received: %d", output["a"][0][0])				
		}

		if output["a"][0][1] != 4 {
			test.Errorf("Expected: 4, Received: %d", output["a"][0][1])				
		}

		if output["a"][1][0] != 4 {
			test.Errorf("Expected: 4, Received: %d", output["a"][1][0])				
		}

		if output["a"][1][1] != 8 {
			test.Errorf("Expected: 8, Received: %d", output["a"][1][1])				
		}

		if output["a"][2][0] != 5 {
			test.Errorf("Expected: 3, Received: %d", output["a"][2][0])				
		}

		if output["a"][2][1] != 5 {
			test.Errorf("Expected: 5, Received: %d", output["a"][2][1])				
		}
	}
}

func TestBuildAntennaLocationsDict_RemovesAnySingleLocationAntennas(test *testing.T) {
	input := `..........
..........
..........
....a.....
........a.
.....a....
..........
.....A....
..........
..........`

	testMap := StringTo2DArray(input)

	output := BuildAntennaLocationsDict(testMap)

	_, exists := output["A"]

	if exists {
		test.Errorf("Expected: 'A' not to exists, Received: 'A' exists")	
	}
}

func TestGetXYDifferenceBetweenTwoAntennas_ReturnsDifference(test *testing.T) {
	testC1 := [2]int{3, 4}
	testC2 := [2]int{4, 8}

	output := GetXYDifferenceBetweenTwoAntennas(testC1, testC2)

	if output[0] != 1 {
		test.Errorf("Expected: 1, Received: %d", output[0])		
	}

	if output[1] != 4 {
		test.Errorf("Expected: 4, Received: %d", output[1])		
	}	
	
	testC1 = [2]int{3, 4}
	testC2 = [2]int{5, 5}

	output = GetXYDifferenceBetweenTwoAntennas(testC1, testC2)

	if output[0] != 2 {
		test.Errorf("Expected: 2, Received: %d", output[0])		
	}

	if output[1] != 1 {
		test.Errorf("Expected: 1, Received: %d", output[1])		
	}

	testC1 = [2]int{4, 8}
	testC2 = [2]int{5, 5}

	output = GetXYDifferenceBetweenTwoAntennas(testC1, testC2)

	if output[0] != 1 {
		test.Errorf("Expected: 1, Received: %d", output[0])		
	}

	if output[1] != -3 {
		test.Errorf("Expected: -3, Received: %d", output[1])		
	}
}

func TestGetLocationIfAntennaGoesBackwards_ReturnsNewCoords(test *testing.T) {
	originalCoOrds := [2]int{3, 4}
	difference := [2]int{1, 4}

	output := GetLocationIfAntennaGoesBackwards(originalCoOrds, difference)

	if output[0] != 2 {
		test.Errorf("Expected: 2, Received: %d", output[0])	
	}

	if output[1] != 0 {
		test.Errorf("Expected: 0, Received: %d", output[1])	
	}

	originalCoOrds = [2]int{3, 4}
	difference = [2]int{2, 1}

	output = GetLocationIfAntennaGoesBackwards(originalCoOrds, difference)

	if output[0] != 1 {
		test.Errorf("Expected: 1, Received: %d", output[0])	
	}

	if output[1] != 3 {
		test.Errorf("Expected: 3, Received: %d", output[1])	
	}

	originalCoOrds = [2]int{4, 8}
	difference = [2]int{1, -3}

	output = GetLocationIfAntennaGoesBackwards(originalCoOrds, difference)

	if output[0] != 3 {
		test.Errorf("Expected: 3, Received: %d", output[0])	
	}

	if output[1] != 11 {
		test.Errorf("Expected: 11, Received: %d", output[1])	
	}
}

func TestGetLocationIfAntennaGoesForwards_ReturnsNewCoords(test *testing.T) {
	originalCoOrds := [2]int{4, 8}
	difference := [2]int{1, 4}

	output := GetLocationIfAntennaGoesForwards(originalCoOrds, difference)

	if output[0] != 5 {
		test.Errorf("Expected: 5, Received: %d", output[0])	
	}

	if output[1] != 12 {
		test.Errorf("Expected: 12, Received: %d", output[1])	
	}

	originalCoOrds = [2]int{5, 5}
	difference = [2]int{2, 1}

	output = GetLocationIfAntennaGoesForwards(originalCoOrds, difference)

	if output[0] != 7 {
		test.Errorf("Expected: 7, Received: %d", output[0])	
	}

	if output[1] != 6 {
		test.Errorf("Expected: 6, Received: %d", output[1])	
	}

	originalCoOrds = [2]int{5, 5}
	difference = [2]int{1, -3}

	output = GetLocationIfAntennaGoesForwards(originalCoOrds, difference)

	if output[0] != 6 {
		test.Errorf("Expected: 6, Received: %d", output[0])	
	}

	if output[1] != 2 {
		test.Errorf("Expected: 2, Received: %d", output[1])	
	}
}

func TestPart1_ReturnsAnswerForTestSmallerInput(test *testing.T) {
	input := `..........
..........
..........
....a.....
........a.
.....a....
..........
..........
..........
..........`

	output := Part1(input)	

	if output != 4 {
		test.Errorf("Expected: 4, Received: %d", output)	
	}
}

func TestPart1_ReturnsAnswerForTestBiggerInput(test *testing.T) {
	input := `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

	output := Part1(input)	

	if output != 14 {
		test.Errorf("Expected: 14, Received: %d", output)	
	}
}