package main

import (
	"testing"
)

func TestConvertRawInputToBlocksByID_ReturnsTheNewStringForSingleDigitIDs(test *testing.T) {
	input := "123"

	output := ConvertRawInputToBlocksByID(input)

	if len(output) != 5 {
		test.Errorf("Expected: 5, Received: %d", len(output))	
	}

	if output[0] != "0" {
		test.Errorf("Expected: '0', Received: %s", output[0])
	}

	if output[1] != ".." {
		test.Errorf("Expected: '..', Received: %s", output[1])
	}

	if output[2] != "1" {
		test.Errorf("Expected: '1', Received: %s", output[2])
	}

	if output[3] != "1" {
		test.Errorf("Expected: '1', Received: %s", output[3])
	}

	if output[4] != "1" {
		test.Errorf("Expected: '1', Received: %s", output[4])
	}

	input = "12345"

	output = ConvertRawInputToBlocksByID(input)

	if len(output) != 11 {
		test.Errorf("Expected: 11, Received: %d", len(output))	
	}

	if output[0] != "0" {
		test.Errorf("Expected: '0', Received: %s", output[0])
	}

	if output[1] != ".." {
		test.Errorf("Expected: '..', Received: %s", output[1])
	}

	if output[2] != "1" {
		test.Errorf("Expected: '1', Received: %s", output[2])
	}

	if output[3] != "1" {
		test.Errorf("Expected: '1', Received: %s", output[3])
	}

	if output[4] != "1" {
		test.Errorf("Expected: '1', Received: %s", output[4])
	}

	if output[5] != "...." {
		test.Errorf("Expected: '....', Received: %s", output[5])
	}

	if output[6] != "2" {
		test.Errorf("Expected: '2', Received: %s", output[6])
	}

	if output[7] != "2" {
		test.Errorf("Expected: '2', Received: %s", output[7])
	}
	
	if output[8] != "2" {
		test.Errorf("Expected: '2', Received: %s", output[8])
	}

	if output[9] != "2" {
		test.Errorf("Expected: '2', Received: %s", output[9])
	}

	if output[10] != "2" {
		test.Errorf("Expected: '2', Received: %s", output[10])
	}

	input = "2333133121414131402"

	output = ConvertRawInputToBlocksByID(input)

	if len(output) != 36 {
		test.Errorf("Expected: 36, Received: %d", len(output))	
	}

	if output[0] != "0" {
		test.Errorf("Expected: '0', Received: %s", output[0])
	}

	if output[1] != "0" {
		test.Errorf("Expected: '0'' Received: %s", output[1])
	}

	if output[2] != "..." {
		test.Errorf("Expected: '...' Received: %s", output[2])
	}

	if output[35] != "9" {
		test.Errorf("Expected: '9' Received: %s", output[35])
	}
}

func TestConvertRawInputToBlocksByID_ReturnsTheNewStringForMultiDigitIDs(test *testing.T) {
	input := "233313312141413140215"

	output := ConvertRawInputToBlocksByID(input)

	if len(output) != 42 {
		test.Errorf("Expected: 42, Received: %d", len(output))	
	}

	if output[0] != "0" {
		test.Errorf("Expected: '0', Received: %s", output[0])
	}

	if output[1] != "0" {
		test.Errorf("Expected: '0'' Received: %s", output[1])
	}

	if output[6] != "..." {
		test.Errorf("Expected: '...' Received: %s", output[6])
	}

	if output[41] != "10" {
		test.Errorf("Expected: '10' Received: %s", output[41])
	}

	if output[41] != "10" {
		test.Errorf("Expected: '10' Received: %s", output[41])
	}
}

func TestFindTotalNumberOfFileBlocks_ReturnsTotalNumOfBlocks(test *testing.T) {
	rawInput := "12345"
	input := ConvertRawInputToBlocksByID(rawInput)

	output := FindTotalNumberOfFileBlocks(input)

	if output != 9 {
		test.Errorf("Expected: 9, Received: %d", output)
	}

	rawInput = "2333133121414131402"
	input = ConvertRawInputToBlocksByID(rawInput)

	output = FindTotalNumberOfFileBlocks(input)

	if output != 28 {
		test.Errorf("Expected: 28, Received: %d", output)
	}

	// Test below includes a double digit ID
	rawInput = "233313312141413140215"
	input = ConvertRawInputToBlocksByID(rawInput)

	output = FindTotalNumberOfFileBlocks(input)

	if output != 33 {
		test.Errorf("Expected: 33, Received: %d", output)
	}
} 

func TestFindNextLastFileBlockToMoveIndex_ReturnsIndexOfNextFileBlock(test *testing.T) {
	rawInput := "12345"
	input := ConvertRawInputToBlocksByID(rawInput)

	output := FindNextLastFileBlockToMoveIndex(input)

	if output != 10 {
		test.Errorf("Expected: 10, Received: %d", output)
	}

	input = []string{"0", "2", ".", "1", "1", "1", "....", "2", "2", "2", "2", "."}

	output = FindNextLastFileBlockToMoveIndex(input)

	if output != 10 {
		test.Errorf("Expected: 10, Received: %d", output)
	}

	input = []string{"0", "2", "2", "1", "1", "1", "....", "2", "2", "2", ".."}

	output = FindNextLastFileBlockToMoveIndex(input)

	if output != 9 {
		test.Errorf("Expected: 9, Received: %d", output)
	}

	input = []string{"0", "2", "2", "1", "1", "1", "2", "2", "..", "2", "...."}

	output = FindNextLastFileBlockToMoveIndex(input)

	if output != 9 {
		test.Errorf("Expected: 9, Received: %d", output)
	}

	input = []string{"0", "2", "2", "1", "1", "1", "2", "2", "..", "2", "....", "10"}

	output = FindNextLastFileBlockToMoveIndex(input)

	if output != 11 {
		test.Errorf("Expected: 11, Received: %d", output)
	}
} 

func TestFindNextFirstDotSetAvailable_ReturnsIndexOfNextDot(test *testing.T) {
	input := []string{"0", "..", "1", "1", "1", "....", "2", "2", "2", "2", "2"}

	fileBlockIndex := FindNextLastFileBlockToMoveIndex(input)
	lengthOfIndex := GetLengthOfIndexItem(input, fileBlockIndex)

	output := FindNextFirstDotSetAvailable(input, lengthOfIndex)

	if output != 1 {
		test.Errorf("Expected: 1, Received: %d", output)
	}

	input = []string{"0", "2", ".", "1", "1", "1", "....", "2", "2", "2", "2", "2"}

	fileBlockIndex = FindNextLastFileBlockToMoveIndex(input)
	lengthOfIndex = GetLengthOfIndexItem(input, fileBlockIndex)

	output = FindNextFirstDotSetAvailable(input, lengthOfIndex)

	if output != 2 {
		test.Errorf("Expected: 2, Received: %d", output)
	}

	input = []string{"0", "2", ".", "1", "1", "1", "....", "2", "2", "2", "2", "2", "10"}

	fileBlockIndex = FindNextLastFileBlockToMoveIndex(input)
	lengthOfIndex = GetLengthOfIndexItem(input, fileBlockIndex)

	output = FindNextFirstDotSetAvailable(input, lengthOfIndex)

	if output != 6 {
		test.Errorf("Expected: 6, Received: %d", output)
	}
} 

func TestIsFileBlockStringFullyMoved_ReturnsFalseForNotFullyMoved(test *testing.T) {
	input := []string{"0", "..", "1", "1", "1", "1", "....", "2", "2", "2", "2", "2"}

	fileBlocksTotal := FindTotalNumberOfFileBlocks(input)

	output := IsFileBlockStringFullyMoved(input, fileBlocksTotal)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	input = []string{"0", "2", ".", "1", "1", "1", "....", "2", "2", "2", "2", "."}
	fileBlocksTotal = FindTotalNumberOfFileBlocks(input)

	output = IsFileBlockStringFullyMoved(input, fileBlocksTotal)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	input = []string{"0", "2", "2", "1", "1", "1", "....", "2", "2", "2", ".", "."}
	fileBlocksTotal = FindTotalNumberOfFileBlocks(input)

	output = IsFileBlockStringFullyMoved(input, fileBlocksTotal)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	input = []string{"0", "2", "2", "1", "1", "1", "2", "...", "2", "2", ".", "."}
	fileBlocksTotal = FindTotalNumberOfFileBlocks(input)

	output = IsFileBlockStringFullyMoved(input, fileBlocksTotal)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	input = []string{"0", "2", "2", "1", "1", "1", "2", "2", "..", "2", ".", ".", ".", "."}
	fileBlocksTotal = FindTotalNumberOfFileBlocks(input)

	output = IsFileBlockStringFullyMoved(input, fileBlocksTotal)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}
} 

func TestIsFileBlockStringFullyMoved_ReturnsTrueForFullyMoved(test *testing.T) {
	input := []string{"0", "2", "2", "1", "1", "1", "2", "2", "2", "......"}

	fileBlocksTotal := FindTotalNumberOfFileBlocks(input)

	output := IsFileBlockStringFullyMoved(input, fileBlocksTotal)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	input = []string{"0", "0", "9", "9", "8", "1", "1", "1", "8", "8", "8", "2", "7", "7", ".", "."}
	fileBlocksTotal = FindTotalNumberOfFileBlocks(input)

	output = IsFileBlockStringFullyMoved(input, fileBlocksTotal)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}
}

func TestMoveFileBlocks_ReturnsTheNewString(test *testing.T) {
	input := []string{"0", "..", "1", "1", "1", "....", "2", "2", "2", "2", "2"}

	output := MoveFileBlocks(input)

	if output[0] != "0" {
		test.Errorf("Expected: '0', Received: %s", output[0])
	}	

	if output[1] != "2" {
		test.Errorf("Expected: '2', Received: %s", output[1])
	}	

	if output[2] != "2" {
		test.Errorf("Expected: '2', Received: %s", output[2])
	}	

	if output[3] != "1" {
		test.Errorf("Expected: '1', Received: %s", output[3])
	}	

	if output[4] != "1" {
		test.Errorf("Expected: '1', Received: %s", output[4])
	}	

	if output[5] != "1" {
		test.Errorf("Expected: '1', Received: %s", output[5])
	}	

	if output[6] != "2" {
		test.Errorf("Expected: '2', Received: %s", output[6])
	}	

	if output[7] != "2" {
		test.Errorf("Expected: '2' Received: %s", output[7])
	}	

	if output[8] != "2" {
		test.Errorf("Expected: '2' Received: %s", output[8])
	}	

	if output[9] != "." {
		test.Errorf("Expected: '.' Received: %s", output[9])
	}	

	rawInput := "2333133121414131402"
	input = ConvertRawInputToBlocksByID(rawInput)

	output = MoveFileBlocks(input)

	if output[0] != "0" {
		test.Errorf("Expected: '0' Received: %s", output[0])
	}	
	
	if output[2] != "9" {
		test.Errorf("Expected: '9' Received: %s", output[2])
	}	

	if output[27] != "6" {
		test.Errorf("Expected: '6' Received: %s", output[27])
	}	

	if output[28] != "." {
		test.Errorf("Expected: '.' Received: %s", output[28])
	}	
}

func TestCalculateCheckSum_ReturnsCorrectSum(test *testing.T) {
	input := []string{"0", "0", "9", "9", "8", "1", "1", "1", "8", "8", "8", "2", "7", "7", "7", "3", "3", "3", "6", "4", "4", "6", "5", "5", "5", "5", "6", "6"}

	output := CalculateCheckSum(input)

	if output != 1928 {
		test.Errorf("Expected: 1928, Received: %d", output)
	}	
}

func TestPart1_ReturnsCorrectAnswerForTestInput(test *testing.T) {
	//input := "2333133121414131402"

	input := "2847796712879033433718603712718817354736281243964888223682694953479950792417879182187383805728208021969669819725574153728544608865468619788443786228879041777294508133256267708489825197157576995557114914777260466750629195935891235828221856509474654017546089994652235862211759249849175492377084661971734055948923317412805147549454262167624990989160957780697726661657713798189894156137528180563284441320838047593273"
	//y := ConvertRawInputToBlocksByID(x)

	//fmt.Println(y)

	output := Part1(input)

	if output != 1928 {
		test.Errorf("Expected: 1928, Received: %d", output)
	}	
}