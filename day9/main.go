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

	// PROBLEM FOR PART 1 IS...... ID'S CAN BE HIGHER THAN 9, D'OH!!! Need to look at this, put everything into an array rather than rely on single string chars?

	// Part 1
	p1Output := Part1(rawInput)

	fmt.Printf("Day 9, Part 1 Output: %d\n", p1Output)

	// Part 2
	//p2Output := Part2(rawInput)

	//fmt.Printf("Day 9, Part 2 Output: %d\n", p2Output)	
}

func Part1(input string) int {
	checksum := 0

	convertedBlocksById := ConvertRawInputToBlocksByID(input)

	movedFileBlocks := MoveFileBlocks(convertedBlocksById)

	checksum = CalculateCheckSum(movedFileBlocks)

	return checksum
}

func Part2() {}

func readFileToString(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)

	if err != nil {
		return "", err
	}
	
	return string(content), nil
}

func ConvertRawInputToBlocksByID(input string) []string {
	splitInput := strings.Split(input, "");
	
	finalString := []string{}

	currentID := 0

	for i := 0; i < len(splitInput); i++ {
		if splitInput[i] != "0" {
			if i % 2 == 0 {
			idAsString := strconv.Itoa(currentID)
			currentElAsNumber, _  := strconv.Atoi(splitInput[i])

			for j := 0; j < currentElAsNumber; j++ {
				stringToAppend := idAsString

				finalString = append(finalString, stringToAppend)	
			}

			currentID++
		}

		if i % 2 == 1 {
			currentElAsNumber, _  := strconv.Atoi(splitInput[i])

			stringToAppend := strings.Repeat(".", currentElAsNumber)

			finalString = append(finalString, stringToAppend)
		}		
		}
	}

	return finalString
}

func MoveFileBlocks(blocks []string) []string {
	fileBlocks := make([]string, len(blocks))
	_ = copy(fileBlocks, blocks)

	fmt.Println(fileBlocks, "before")

	// Need know when everything sorted... first work out how many chars there are which are not a '.'. Save this into a var, we know when the final string is finished being moved when the first X indexes are all not '.'s (e.g. if there is 5 non . chars, and the final string has indexes 0-4 as not '.'s then it's done)

	// Begin the moving (a while loop here checking if everything is sorted or not?)
	// --> Find the next last char to get ready to move (.FindNextLastFileBlockToMoveIndex()), return the index so we know which char to move
	// --> Find the next first set of '.'s available to put a char into (.FindNextFirstDotAvailable()) return the index so we know which dot to move a char into (the set of dots needs to be enough to fit the string going into it)
	// --> Put the char into the dot
	// --> change where the char WAS into a . at the end if a single dot, or move dots along to next index
	// --> needs a bunch of manipulation and moving stuff around!

	totalFileBlocks := FindTotalNumberOfFileBlocks(fileBlocks)
	areAllFileBlocksSorted := false

	for !areAllFileBlocksSorted {
		nextFileToMoveIndex := FindNextLastFileBlockToMoveIndex(fileBlocks)
		fileNameToMove := fileBlocks[nextFileToMoveIndex]
		lengthOfIndex := GetLengthOfIndexItem(fileBlocks, nextFileToMoveIndex)

		nextDotToMoveFileToIndex := FindNextFirstDotSetAvailable(fileBlocks, lengthOfIndex)

		numOfDots := len(fileBlocks[nextDotToMoveFileToIndex])
		
		if numOfDots > 1 {
			numOfDotsToMoveToNextIndex := numOfDots - lengthOfIndex

			fmt.Println("dots left:", lengthOfIndex, numOfDots, numOfDotsToMoveToNextIndex)
			
			if numOfDotsToMoveToNextIndex > 0 {
				dotsStringToSort := strings.Repeat(".", numOfDotsToMoveToNextIndex)

				indexToPutDotsIn := nextDotToMoveFileToIndex + 1

				tempSlice := make([]string, len(fileBlocks) + 1)
				copy(tempSlice[:nextDotToMoveFileToIndex], fileBlocks[:nextDotToMoveFileToIndex])

				copy(tempSlice[indexToPutDotsIn:], fileBlocks[nextDotToMoveFileToIndex:])

				tempSlice[indexToPutDotsIn] = dotsStringToSort
				tempSlice[nextDotToMoveFileToIndex] = fileNameToMove
				tempSlice[nextFileToMoveIndex + 1] = strings.Repeat(".", numOfDotsToMoveToNextIndex)

				fileBlocks = tempSlice
			} 

			if numOfDotsToMoveToNextIndex == 0 {
				fmt.Println(nextDotToMoveFileToIndex, fileNameToMove, nextFileToMoveIndex)
				fileBlocks[nextDotToMoveFileToIndex] = fileNameToMove
				fileBlocks[nextFileToMoveIndex] = strings.Repeat(".", numOfDotsToMoveToNextIndex)
			}


		} 
		
		if numOfDots == 1 {
			fileBlocks[nextDotToMoveFileToIndex] = fileNameToMove

			fileBlocks[nextFileToMoveIndex] = "."
		}

		totalFileBlocks = FindTotalNumberOfFileBlocks(fileBlocks)

		areAllFileBlocksSorted = IsFileBlockStringFullyMoved(fileBlocks, totalFileBlocks)

	}

	

	fmt.Println(fileBlocks, "after")

	return fileBlocks
}

func FindTotalNumberOfFileBlocks(input []string) int {
	totalFileBlocks := 0

	for i := 0; i < len(input); i++ {
		if !strings.Contains(input[i], ".") && !strings.Contains(input[i], " ") {
			totalFileBlocks++
		}
	}

	return totalFileBlocks
}

func FindNextLastFileBlockToMoveIndex(input []string) int {
	fileBlockIndex := 0

	for i := len(input) - 1; i >= 0; i-- {
		if !strings.Contains(input[i], ".") && !strings.Contains(input[i], " ") {
			fileBlockIndex = i
			break
		}
	}

	return fileBlockIndex
}

func GetLengthOfIndexItem(input []string, index int) int {
	return len(input[index])
}

func FindNextFirstDotSetAvailable(input []string, lengthNeeded int) int {
	dotIndex := 0

	for i := 0; i < len(input); i++ {
		if strings.Contains(input[i], ".") && lengthNeeded <= len(input[i]) {
			dotIndex = i
			break
		}
	}

	return dotIndex
}

func IsFileBlockStringFullyMoved(input []string, totalFileBlocks int) bool {
	isFileBlockMoved := true

	for i := 0; i < totalFileBlocks; i++ {
		if strings.Contains(input[i], ".") {
			isFileBlockMoved = false
			break
		}
	}

	return isFileBlockMoved
}

func CalculateCheckSum(input []string) int {
	total := 0

	for i := 0; i < len(input); i++ {
		if strings.Contains(input[i], ".") {
			break
		}

		num, _ := strconv.Atoi(input[i])

		currentSum := i * num

		total += currentSum
	}

	return total
}
