package main

import (
	"testing"
)

func generateTestLists() ([][]int, [][]int) {
	input := `47|53
	97|13
	97|61
	97|47
	75|29
	61|13
	75|53
	29|13
	97|29
	53|29
	61|53
	97|53
	61|29
	47|13
	75|47
	97|75
	47|61
	75|61
	47|29
	75|13
	53|13
	
	75,47,61,53,29
	97,61,53,29,13
	75,29,13
	75,97,47,61,53
	61,13,29
	97,13,75,29,47`
	
	orderRulesList, pagesToProduceList := GenerateOrderRulesAndPagesToProduceLists(input)

	return orderRulesList, pagesToProduceList
}

func TestGenerateOrderRulesAndPagesToProduceLists_ProducesBothLists(test *testing.T) {
	input := `53|13

75,47,61,53,29`

	orderRulesList, pagesToProduceList := GenerateOrderRulesAndPagesToProduceLists(input)

	if len(orderRulesList) != 1 {
		test.Errorf("Expected: 1, Received: %d", len(orderRulesList))
	}

	if orderRulesList[0][0] != 53 {
		test.Errorf("Expected: 53, Received: %d", orderRulesList[0][0])
	}

	if orderRulesList[0][1] != 13 {
		test.Errorf("Expected: 13, Received: %d", orderRulesList[0][1])
	}

	if len(pagesToProduceList) != 1 {
		test.Errorf("Expected: 1, Received: %d", len(pagesToProduceList))
	}

	if pagesToProduceList[0][0] != 75 {
		test.Errorf("Expected: 75, Received: %d", pagesToProduceList[0][0])
	}

	if pagesToProduceList[0][1] != 47 {
		test.Errorf("Expected: 47, Received: %d", pagesToProduceList[0][1])
	}

	if pagesToProduceList[0][4] != 29 {
		test.Errorf("Expected: 29, Received: %d", pagesToProduceList[0][4])
	}

	input = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	orderRulesList, pagesToProduceList = GenerateOrderRulesAndPagesToProduceLists(input)

	if len(orderRulesList) != 21 {
		test.Errorf("Expected: 21, Received: %d", len(orderRulesList))
	}

	if len(pagesToProduceList) != 6 {
		test.Errorf("Expected: 6, Received: %d", len(pagesToProduceList))
	}
}

func TestGenerateInDegreeMap_ReturnsMapOfPageKeysAndZeroedValues(test *testing.T) {
	input := []int{75, 47, 61, 53, 29}

	output := GenerateInDegreeMap(input)

	if len(output) != 5 {
		test.Errorf("Expected: 5, Received: %d", len(output))
	}
} 

func TestIsUpdateInRightOrder_ReturnsTrueForListInRightOrder(test *testing.T) {
	rulesList := [][]int{{47, 53}, {97, 13}, {97, 61}, {97, 47}, {75, 29}, {61, 13}, {75, 53}, {29, 13}, {97, 29}, {53, 29}, {61, 53}, {97, 53}, {61, 29}, {47, 13}, {75, 47}, {97, 75}, {47, 61}, {75, 61}, {47, 29}, {75, 13}, {53, 13}}

	pagesList := []int{75,47,61,53,29}

	output := IsUpdateInRightOrder(rulesList, pagesList)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	pagesList = []int{97,61,53,29,13}

	output = IsUpdateInRightOrder(rulesList, pagesList)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}

	pagesList = []int{75,29,13}

	output = IsUpdateInRightOrder(rulesList, pagesList)

	if output != true {
		test.Errorf("Expected: true, Received: %v", output)
	}
} 

func TestIsUpdateInRightOrder_ReturnsFalseForListNotInRightOrder(test *testing.T) {
	rulesList := [][]int{{47, 53}, {97, 13}, {97, 61}, {97, 47}, {75, 29}, {61, 13}, {75, 53}, {29, 13}, {97, 29}, {53, 29}, {61, 53}, {97, 53}, {61, 29}, {47, 13}, {75, 47}, {97, 75}, {47, 61}, {75, 61}, {47, 29}, {75, 13}, {53, 13}}

	pagesList := []int{75,97,47,61,53}

	output := IsUpdateInRightOrder(rulesList, pagesList)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	pagesList = []int{61,13,29}

	output = IsUpdateInRightOrder(rulesList, pagesList)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}

	pagesList = []int{97,13,75,29,47}

	output = IsUpdateInRightOrder(rulesList, pagesList)

	if output != false {
		test.Errorf("Expected: false, Received: %v", output)
	}
} 

func TestFindMiddlePageNumber_ReturnsMiddlePageNumber(test *testing.T) {
	pagesList := []int{75}

	output := FindMiddlePageNumber(pagesList)

	if output != 75 {
		test.Errorf("Expected: 75, Received: %d", output)
	}

	pagesList = []int{75, 47}

	output = FindMiddlePageNumber(pagesList)

	if output != 75 {
		test.Errorf("Expected: 75, Received: %d", output)
	}

	pagesList = []int{75,29,13}

	output = FindMiddlePageNumber(pagesList)

	if output != 29 {
		test.Errorf("Expected: 29, Received: %d", output)
	}

	pagesList = []int{75,47,61,53,29}

	output = FindMiddlePageNumber(pagesList)

	if output != 61 {
		test.Errorf("Expected: 61, Received: %d", output)
	}

	pagesList = []int{97,61,53,29,13}

	output = FindMiddlePageNumber(pagesList)

	if output != 53 {
		test.Errorf("Expected: 53, Received: %d", output)
	}
}

func TestPart1_FullTestReturnsAnswerWithTestData(test *testing.T) {
	input := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	output := Part1(input)

	if output != 143 {
		test.Errorf("Expected: 143, Received: %d", output)
	}
}