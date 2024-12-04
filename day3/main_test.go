package main

import (
	"testing"
)

func TestExtractMulsFromString_ReturnsEmptySliceForNoInput(test *testing.T) {
	input := ""

	output := ExtractMulsFromString(input)

	if len(output) != 0 {
		test.Errorf("Expected: 0, Received: %d", len(output))
	}
}

func TestExtractMulsFromString_ReturnsOneValidMulForSingleValidMulGiven(test *testing.T) {
	input := "mul(1,2)"

	output := ExtractMulsFromString(input)

	if len(output) != 1 {
		test.Errorf("Expected: 1, Received: %d", len(output))
	}

	input = "mul(14,32)"

	output = ExtractMulsFromString(input)

	if len(output) != 1 {
		test.Errorf("Expected: 1, Received: %d", len(output))
	}

	input = "mul(191,262)"

	output = ExtractMulsFromString(input)

	if len(output) != 1 {
		test.Errorf("Expected: 1, Received: %d", len(output))
	}

	input = "mul(602,165)"

	output = ExtractMulsFromString(input)

	if len(output) != 1 {
		test.Errorf("Expected: 1, Received: %d", len(output))
	}

	input = "mul(8,71)"

	output = ExtractMulsFromString(input)

	if len(output) != 1 {
		test.Errorf("Expected: 1, Received: %d", len(output))
	}

	input = "mul(12,1)"

	output = ExtractMulsFromString(input)

	if len(output) != 1 {
		test.Errorf("Expected: 1, Received: %d", len(output))
	}

	input = "mul(13,999)"

	output = ExtractMulsFromString(input)

	if len(output) != 1 {
		test.Errorf("Expected: 1, Received: %d", len(output))
	}
}

func TestExtractMulsFromString_ReturnsNoValidMulForSingleInvalidMulGiven(test *testing.T) {
	input := "mil(1,2)"

	output := ExtractMulsFromString(input)

	if len(output) != 0 {
		test.Errorf("Expected: 0, Received: %d", len(output))
	}

	input = "mul(1,9999)"

	output = ExtractMulsFromString(input)

	if len(output) != 0 {
		test.Errorf("Expected: 0, Received: %d", len(output))
	}

	input = "mul[14,32]"

	output = ExtractMulsFromString(input)

	if len(output) != 0 {
		test.Errorf("Expected: 0, Received: %d", len(output))
	}

	input = "mul ( 2 , 4 )"

	output = ExtractMulsFromString(input)

	if len(output) != 0 {
		test.Errorf("Expected: 0, Received: %d", len(output))
	}

	input = "mul(4*"

	output = ExtractMulsFromString(input)

	if len(output) != 0 {
		test.Errorf("Expected: 0, Received: %d", len(output))
	}

	input = "mul(6,9!"

	output = ExtractMulsFromString(input)

	if len(output) != 0 {
		test.Errorf("Expected: 0, Received: %d", len(output))
	}

	input = "?(12,34)"

	output = ExtractMulsFromString(input)

	if len(output) != 0 {
		test.Errorf("Expected: 0, Received: %d", len(output))
	}
}

func TestExtractMulsFromString_ReturnsValidMulsForStringWithNoCorruptedChars(test *testing.T) {
	input := "mul(1,2)mul(11,22)mul(371,997)"

	output := ExtractMulsFromString(input)

	if len(output) != 3 {
		test.Errorf("Expected: 3, Received: %d", len(output))
	}
}

func TestExtractMulsFromString_ReturnsNoMulsForStringWithCorruptedCharsAndInValidMuls(test *testing.T) {
	input := "mul[3,7]!$mul(32,64]^mul(4*"

	output := ExtractMulsFromString(input)

	if len(output) != 0 {
		test.Errorf("Expected: 0, Received: %d", len(output))
	}
}

func TestExtractMulsFromString_ReturnsValidMulsForStringWithCorruptedChars(test *testing.T) {
	input := "mul(1,2)^mul(11,22)@mul(371,997)"

	output := ExtractMulsFromString(input)

	if len(output) != 3 {
		test.Errorf("Expected: 3, Received: %d", len(output))
	}

	input = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	output = ExtractMulsFromString(input)

	if len(output) != 4 {
		test.Errorf("Expected: 4, Received: %d", len(output))
	}
}

func TestMultiplySingleMul_ReturnsSumForGivenValidMul(test *testing.T) {
	input := "mul(2,4)"

	output := MultiplySingleMul(input)

	if output != 8 {
		test.Errorf("Expected: 8, Received: %d", output)
	}

	input = "mul(5,5)"

	output = MultiplySingleMul(input)

	if output != 25 {
		test.Errorf("Expected: 25, Received: %d", output)
	}

	input = "mul(11,8)"

	output = MultiplySingleMul(input)

	if output != 88 {
		test.Errorf("Expected: 88, Received: %d", output)
	}

	input = "mul(8,5)"

	output = MultiplySingleMul(input)

	if output != 40 {
		test.Errorf("Expected: 40, Received: %d", output)
	}

	input = "mul(999,999)"

	output = MultiplySingleMul(input)

	if output != 998001 {
		test.Errorf("Expected: 998001, Received: %d", output)
	}
}

func TestSumAllMuls_ReturnsZeroForNoMuls(test *testing.T) {
	input := []string{}

	output := SumAllMuls(input)

	if output != 0 {
		test.Errorf("Expected: 0, Received: %d", output)
	}
}

func TestSumAllMuls_ReturnsTotalForSingleMul(test *testing.T) {
	input := []string{"mul(2,4)"}

	output := SumAllMuls(input)

	if output != 8 {
		test.Errorf("Expected: 8, Received: %d", output)
	}
}

func TestSumAllMuls_ReturnsTotalForManyMuls(test *testing.T) {
	input := []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}

	output := SumAllMuls(input)

	if output != 161 {
		test.Errorf("Expected: 161, Received: %d", output)
	}
}

func TestCorruptedMemory_Part2_CorrectTestOutput(test *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	output := CorruptedMemory_Part2(input)

	if output != 48 {
		test.Errorf("Expected: 48, Received: %d", output)
	}
}