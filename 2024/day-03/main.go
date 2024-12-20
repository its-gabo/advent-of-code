package main

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strings"
)

var inputData string

func init() {
	currentDirectory := utils.GetCurrentModDirectory()
	inputData = utils.GetInputDataFromFile(currentDirectory + "/day-03/dayThree.txt")
}

func main() {
	fmt.Println("Part 1 answer: ", partOne())
	fmt.Println("Part 2 answer: ", partTwo())
}

func partOne() int {
	matches := [][]int{}

	matchedSequences := parseInput(inputData, `mul\([0-9]*\,[0-9]*\)`)

	for _, match := range matchedSequences {
		nums := strings.Split(strings.TrimRight(strings.TrimLeft(match, "mul("), ")"), ",")

		matches = append(matches, []int{utils.CastToInt(nums[0]) , utils.CastToInt(nums[1])})
	}

	multiplicationsResult := 0

	for _, pair := range matches {
		multiplicationsResult += pair[0] * pair[1]
	}

	return multiplicationsResult
}

func partTwo() int {
	isMultiplicationEnabled := true
	multiplicationsResult := 0

	matchedSequences := parseInput(inputData, `(mul\([0-9]*\,[0-9]*\)|do\(\)|don't\(\))`)

	for _, match := range matchedSequences {
		match = trimParentheses(match)

		if(match == "do") {
			isMultiplicationEnabled = true
		} else if(match == "don't") {
			isMultiplicationEnabled = false
		}

		if(isMultiplicationEnabled && strings.Contains(match, "mul")) {
			nums := strings.Split(strings.ReplaceAll(match, "mul", ""), ",")

			multiplicationsResult += utils.CastToInt(nums[0]) * utils.CastToInt(nums[1])
		}

	}

	return multiplicationsResult
}

func parseInput(inputData string, regex string) (matchedSequences []string) {
	reg := regexp.MustCompile(regex)

	matchedSequences = append(matchedSequences, reg.FindAllString(inputData, -1)...)

	return matchedSequences
}

func trimParentheses(str string) string {
	return strings.ReplaceAll(strings.ReplaceAll(str, "(", ""), ")", "")
}