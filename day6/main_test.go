package main

import (
	"testing"
)

// generate 2d map tests ✅
// get current position of guard tests ✅
// get which current way guard is facing tests ✅
// is next step an obstruction tests ✅
// taking step forward if not obstruction (updates guard position) tests ✅
// turn right90degrees-update map with new arrow (1 of 4 arrows ^ > < v) depending where currently are tests ✅
// moving forward - change position on map depending on which arrow (left/right/up/down) current tests ✅
// updating guard visited points to X tests ✅
// is guard on the map anymore tests (^ < > v no longer on the map) ✅
// count total Xs on the map tests ✅

func generateTestMap() [][]string {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	testMap := StringTo2DArray(input)

	return testMap
}

func TestStringTo2DArray_Produces2DArrayFromInput(test *testing.T) {
	input := `....#.....
.........#`

	output := StringTo2DArray(input)

	if len(output) != 2 {
		test.Errorf("Expected: 2, Received: %d", len(output))
	}

	if output[0][0] != "." {
		test.Errorf("Expected: '.', Received: %s", output[0][0])
	}

	if output[1][9] != "#" {
		test.Errorf("Expected: '#', Received: %s", output[1][9])
	}

	input = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
	
		output = StringTo2DArray(input)
	
		if len(output) != 10 {
			test.Errorf("Expected: 10, Received: %d", len(output))
		}
}

func TestGetGuardCurrentPosition_ReturnsGuardPositionRowAndColumn(test *testing.T) {
	testMap := generateTestMap()

	row, col := GetGuardCurrentPosition(testMap)

	if row != 6 {
		test.Errorf("Expected: 6, Received: %d", row)
	}

	if col != 4 {
		test.Errorf("Expected: 4, Received: %d", col)
	}

	testMapRaw := `....#.....
........>#
..........
..#.......
.......#..
..........
.#........
........#.
#.........
......#...`

	testMap = StringTo2DArray(testMapRaw)

	row, col = GetGuardCurrentPosition(testMap)

	if row != 1 {
		test.Errorf("Expected: 1, Received: %d", row)
	}

	if col != 8 {
		test.Errorf("Expected: 8, Received: %d", col)
	}

	testMapRaw = `....#.....
.........#
..........
..#.......
.......#..
..........
.#........
........#.
#.........
......#v..`

	testMap = StringTo2DArray(testMapRaw)

	row, col = GetGuardCurrentPosition(testMap)

	if row != 9 {
		test.Errorf("Expected: 9, Received: %d", row)
	}

	if col != 7 {
		test.Errorf("Expected: 7, Received: %d", col)
	}

	testMapRaw = `....#.....
.........#
..........
..#.......
.......#..
..........
.#........
....<...#.
#.........
......#...`

	testMap = StringTo2DArray(testMapRaw)

	row, col = GetGuardCurrentPosition(testMap)

	if row != 7 {
		test.Errorf("Expected: 7, Received: %d", row)
	}

	if col != 4 {
		test.Errorf("Expected: 4, Received: %d", col)
	}
}

func TestGetWhichWayGuardIsFacing_ReturnsWhichWayGuardIsFacing(test *testing.T) {
	testMap := generateTestMap()

	output := GetWhichWayGuardIsFacing(testMap)

	if output != "^" {
		test.Errorf("Expected: '^', Received: %s", output)		
	}

	testMapRaw := `....#.....
........>#
..........
..#.......
.......#..
..........
.#........
........#.
#.........
......#...`

	testMap = StringTo2DArray(testMapRaw)

	output = GetWhichWayGuardIsFacing(testMap)

	if output != ">" {
		test.Errorf("Expected: '>', Received: %s", output)		
	}

	testMapRaw = `....#.....
.........#
..........
..#.......
.......#..
..........
.#........
........#.
#.........
......#v..`

	testMap = StringTo2DArray(testMapRaw)

	output = GetWhichWayGuardIsFacing(testMap)

	if output != "v" {
		test.Errorf("Expected: 'v', Received: %s", output)		
	}

	testMapRaw = `....#.....
.........#
..........
..#.......
.......#..
..........
.#........
....<...#.
#.........
......#...`

	testMap = StringTo2DArray(testMapRaw)

	output = GetWhichWayGuardIsFacing(testMap)

	if output != "<" {
		test.Errorf("Expected: '<', Received: %s", output)		
	}

	testMapRaw = `....#.....
.........#
..........
..#.......
.......#..
..........
.#........
........#.
#.........
......#...`

	testMap = StringTo2DArray(testMapRaw)

	output = GetWhichWayGuardIsFacing(testMap)

	if output != "" {
		test.Errorf("Expected: '', Received: %s", output)		
	}
}

func TestIsNextMoveAnObstruction_ReturnsFalseForNoObstruction(test *testing.T) {
	testMap := generateTestMap()

	row, col := GetGuardCurrentPosition(testMap)

	output := IsNextMoveAnObstruction(testMap, row - 1, col)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)	
	}
} 

func TestIsNextMoveAnObstruction_ReturnsTrueForObstruction(test *testing.T) {
	testMapRaw := `....#.....
	........>#
	..........
	..#.......
	.......#..
	..........
	.#........
	........#.
	#.........
	......#...`

	testMap := StringTo2DArray(testMapRaw)

	row, col := GetGuardCurrentPosition(testMap)

	output := IsNextMoveAnObstruction(testMap, row, col + 1)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)	
	}
} 

func TestTurn90DegreesRight_ChangesFromUpToRight(test *testing.T) {
	testMap := generateTestMap()

	curRowPos, curColPos := GetGuardCurrentPosition(testMap)
	direction := GetWhichWayGuardIsFacing(testMap)

	if testMap[curRowPos][curColPos] != "^" {
		test.Errorf("Expected: '^', Received: %s", testMap[curRowPos][curColPos])	
	}

	Turn90DegreesRight(testMap, curRowPos, curColPos, direction)

	if testMap[curRowPos][curColPos] != ">" {
		test.Errorf("Expected: '>', Received: %s", testMap[curRowPos][curColPos])	
	}
}

func TestTurn90DegreesRight_ChangesFromRightToDown(test *testing.T) {
	testMap := generateTestMap()

	curRowPos, curColPos := GetGuardCurrentPosition(testMap)
	direction := GetWhichWayGuardIsFacing(testMap)

	testMap[curRowPos][curColPos] = ">"

	if testMap[curRowPos][curColPos] != ">" {
		test.Errorf("Expected: '>', Received: %s", testMap[curRowPos][curColPos])	
	}

	Turn90DegreesRight(testMap, curRowPos, curColPos, direction)

	if testMap[curRowPos][curColPos] != "v" {
		test.Errorf("Expected: 'v', Received: %s", testMap[curRowPos][curColPos])	
	}
}

func TestTurn90DegreesRight_ChangesFromDownToLeft(test *testing.T) {
	testMap := generateTestMap()

	curRowPos, curColPos := GetGuardCurrentPosition(testMap)
	direction := GetWhichWayGuardIsFacing(testMap)

	testMap[curRowPos][curColPos] = "v"

	if testMap[curRowPos][curColPos] != "v" {
		test.Errorf("Expected: 'v', Received: %s", testMap[curRowPos][curColPos])	
	}

	Turn90DegreesRight(testMap, curRowPos, curColPos, direction)

	if testMap[curRowPos][curColPos] != "<" {
		test.Errorf("Expected: '<', Received: %s", testMap[curRowPos][curColPos])	
	}
}

func TestTurn90DegreesRight_ChangesFromLeftToUp(test *testing.T) {
	testMap := generateTestMap()

	curRowPos, curColPos := GetGuardCurrentPosition(testMap)
	direction := GetWhichWayGuardIsFacing(testMap)

	testMap[curRowPos][curColPos] = "<"

	if testMap[curRowPos][curColPos] != "<" {
		test.Errorf("Expected: '<', Received: %s", testMap[curRowPos][curColPos])	
	}

	Turn90DegreesRight(testMap, curRowPos, curColPos, direction)

	if testMap[curRowPos][curColPos] != "^" {
		test.Errorf("Expected: '^', Received: %s", testMap[curRowPos][curColPos])	
	}
}

func TestTakeStepFoward_MovesForwardIfFacingUp(test *testing.T) {
	testMap := generateTestMap()

	curRowPos, curColPos := GetGuardCurrentPosition(testMap)
	direction := GetWhichWayGuardIsFacing(testMap)
	visitedMap := CreateVisitedLogArray(testMap)

	TakeStepFoward(testMap, curRowPos, curColPos, direction, visitedMap)

	newRowPos, newColPos := GetGuardCurrentPosition(testMap) 

	if newRowPos != 5 {
		test.Errorf("Expected: 5, Received: %d", newRowPos)	
	}

	if newColPos != 4 {
		test.Errorf("Expected: 4, Received: %d", newRowPos)	
	}

	if visitedMap[6][4] != 1 {
		test.Errorf("Expected: 1, Received: %d", visitedMap[6][4])	
	}
} 

func TestTakeStepFoward_MovesForwardIfFacingDown(test *testing.T) {
	testMapRaw := `....#.....
.........#
..........
..#.......
.......#..
..........
.#......v.
..........
#.........
......#...`

	testMap := StringTo2DArray(testMapRaw)

	curRowPos, curColPos := GetGuardCurrentPosition(testMap)
	direction := GetWhichWayGuardIsFacing(testMap)
	visitedMap := CreateVisitedLogArray(testMap)

	TakeStepFoward(testMap, curRowPos, curColPos, direction, visitedMap)

	newRowPos, newColPos := GetGuardCurrentPosition(testMap) 

	if newRowPos != 7 {
		test.Errorf("Expected: 7, Received: %d", newRowPos)	
	}

	if newColPos != 8 {
		test.Errorf("Expected: 8, Received: %d", newRowPos)	
	}

	if visitedMap[6][8] != 1 {
		test.Errorf("Expected: 1, Received: %d", visitedMap[6][8])	
	}
} 

func TestTakeStepFoward_MovesForwardIfFacingRight(test *testing.T) {
	testMapRaw := `....#.....
.........#
..........
..#.......
.......#..
..........
.#........
..........
#..>......
......#...`

	testMap := StringTo2DArray(testMapRaw)

	curRowPos, curColPos := GetGuardCurrentPosition(testMap)
	direction := GetWhichWayGuardIsFacing(testMap)
	visitedMap := CreateVisitedLogArray(testMap)

	TakeStepFoward(testMap, curRowPos, curColPos, direction, visitedMap)

	newRowPos, newColPos := GetGuardCurrentPosition(testMap) 

	if newRowPos != 8 {
		test.Errorf("Expected: 8, Received: %d", newRowPos)	
	}

	if newColPos != 4 {
		test.Errorf("Expected: 4, Received: %d", newRowPos)	
	}

	if visitedMap[8][3] != 1 {
		test.Errorf("Expected: 1, Received: %d", visitedMap[8][3])	
	}
} 

func TestTakeStepFoward_MovesForwardIfFacingLeft(test *testing.T) {
	testMapRaw := `....#.....
.........#
..........
..#...<...
.......#..
..........
.#........
..........
#.........
......#...`

	testMap := StringTo2DArray(testMapRaw)

	curRowPos, curColPos := GetGuardCurrentPosition(testMap)
	direction := GetWhichWayGuardIsFacing(testMap)
	visitedMap := CreateVisitedLogArray(testMap)

	TakeStepFoward(testMap, curRowPos, curColPos, direction, visitedMap)

	newRowPos, newColPos := GetGuardCurrentPosition(testMap) 

	if newRowPos != 3 {
		test.Errorf("Expected: 3, Received: %d", newRowPos)	
	}

	if newColPos != 5 {
		test.Errorf("Expected: 5, Received: %d", newRowPos)	
	}

	if visitedMap[3][6] != 1 {
		test.Errorf("Expected: 1, Received: %d", visitedMap[3][6])	
	}
} 

func TestTakeStepFoward_DoesNotCauseErrorIfGuardGoesOffMapBounds_InAnyDirection(test *testing.T) {
	testMapRaw := `....#....>
.........#
..........
..#.......
.......#..
..........
.#........
..........
#.........
......#...`
	
		testMap := StringTo2DArray(testMapRaw)
	
		curRowPos, curColPos := GetGuardCurrentPosition(testMap)
		direction := GetWhichWayGuardIsFacing(testMap)
		visitedMap := CreateVisitedLogArray(testMap)
	
		TakeStepFoward(testMap, curRowPos, curColPos, direction, visitedMap)

	testMapRaw = `....#.....
.........#
..........
..#.......
.......#..
..........
.#........
..........
#.........
......#..v`
	
		testMap = StringTo2DArray(testMapRaw)

		curRowPos, curColPos = GetGuardCurrentPosition(testMap)
		direction = GetWhichWayGuardIsFacing(testMap)
		visitedMap = CreateVisitedLogArray(testMap)
		
		TakeStepFoward(testMap, curRowPos, curColPos, direction, visitedMap)


	testMapRaw = `....#.....
.........#
..........
..#.......
.......#..
<.........
.#........
..........
#.........
......#...`
	
		testMap = StringTo2DArray(testMapRaw)

		curRowPos, curColPos = GetGuardCurrentPosition(testMap)
		direction = GetWhichWayGuardIsFacing(testMap)
		visitedMap = CreateVisitedLogArray(testMap)
		
		TakeStepFoward(testMap, curRowPos, curColPos, direction, visitedMap)


	testMapRaw = `....#.....
....^....#
..........
..#.......
.......#..
..........
.#........
..........
#.........
......#...`
	
		testMap = StringTo2DArray(testMapRaw)

		curRowPos, curColPos = GetGuardCurrentPosition(testMap)
		direction = GetWhichWayGuardIsFacing(testMap)
		visitedMap = CreateVisitedLogArray(testMap)
		
		TakeStepFoward(testMap, curRowPos, curColPos, direction, visitedMap)
}

func TestIsGuardOnMapAnymore_ReturnsTrueForGuardOnMap(test *testing.T) {
	testMap := generateTestMap()

	output := IsGuardOnMapAnymore(testMap)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)	
	}

	testMap[6][4] = "."
	testMap[6][5] = ">"

	output = IsGuardOnMapAnymore(testMap)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)	
	}

	testMap[6][5] = "."
	testMap[3][2] = "v"

	output = IsGuardOnMapAnymore(testMap)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)	
	}

	testMap[3][2] = "."
	testMap[9][4] = "<"

	output = IsGuardOnMapAnymore(testMap)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)	
	}
}

func TestIsGuardOnMapAnymore_ReturnsFalseForGuardNotOnMap(test *testing.T) {
	testMap := generateTestMap()

	testMap[6][4] = "."

	output := IsGuardOnMapAnymore(testMap)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)	
	}
}

func TestCreateVisitedLogArray_ReturnsArraySameSizeAsMapAllSetToZero(test *testing.T) {
	testMap := generateTestMap()

	output := CreateVisitedLogArray(testMap)

	if len(output) != 10 {
		test.Errorf("Expected: 10, Received: %d", len(output))	
	}

	if len(output[0]) != 10 {
		test.Errorf("Expected: 10, Received: %d", len(output))	
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

func TestCreateMarkLocationAsVisited_UpdatesArrayWithVisitedLocation(test *testing.T) {
	testMap := generateTestMap()

	visitedMap := CreateVisitedLogArray(testMap)

	MarkLocationAsVisited(visitedMap, 2, 3)

	if visitedMap[2][3] != 1 {
		test.Errorf("Expected: 1, Received: %d", visitedMap[2][3])	
	}

	MarkLocationAsVisited(visitedMap, 4, 6)

	if visitedMap[4][6] != 1 {
		test.Errorf("Expected: 1, Received: %d", visitedMap[4][6])	
	}

	MarkLocationAsVisited(visitedMap, 2, 3)

	if visitedMap[0][9] != 0 {
		test.Errorf("Expected: 0, Received: %d", visitedMap[0][9])	
	}
}

func TestCountTotalLocationsVisited_ReturnsTotalForZero(test *testing.T) {
	testMap := generateTestMap()

	visitedMap := CreateVisitedLogArray(testMap)

	output := CountTotalLocationsVisited(visitedMap)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)	
	}	
}

func TestCountTotalLocationsVisited_ReturnsTotalForOneOrMore(test *testing.T) {
	testMap := generateTestMap()

	visitedMap := CreateVisitedLogArray(testMap)

	MarkLocationAsVisited(visitedMap, 0, 1)
	MarkLocationAsVisited(visitedMap, 1, 2)
	MarkLocationAsVisited(visitedMap, 1, 2)
	MarkLocationAsVisited(visitedMap, 2, 3)
	MarkLocationAsVisited(visitedMap, 4, 5)

	output := CountTotalLocationsVisited(visitedMap)

	if output != 4 {
		test.Errorf("Expected: 4, Received: %d", output)	
	}	
}

func TestPart1_ReturnsTotalForTestInput(test *testing.T) {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	output := Part1(input)

	if output != 41 {
		test.Errorf("Expected: 41, Received: %d", output)	
	}	
}