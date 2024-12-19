package main

import (
	"advent-of-code/utils"
	"fmt"
	"sort"
	"strings"
)

var inputData string

func init() {
	currentDirectory := utils.GetCurrentModDirectory()
	inputData = utils.GetInputDataFromFile(currentDirectory + "/day-01/dayOne.txt")
}

func main() {
	fmt.Println("Part 1 answer: ", partOne())
	fmt.Println("Part 2 answer: ", partTwo())
}

func partOne() int {
	listLeft, listRight := parseInput(inputData)

	sort.Ints(listLeft)
	sort.Ints(listRight)

	distance := 0

	for  i := range listLeft {
		distance += utils.AbsInt(listLeft[i] - listRight[i])
	}
	
	return distance
}

func partTwo() int {
	listLeft, listRight := parseInput(inputData)

	sort.Ints(listLeft)
	sort.Ints(listRight)

	repeatCount := map[int]int{}

	for _, num := range listRight {
		repeatCount[num]++
	}

	similarityScore := 0

	for _, num := range listLeft {
		similarityScore += num * repeatCount[num]
	}

	return similarityScore;
}

func parseInput(inputData string) (listLeft, listRight []int) {
	for _, line := range strings.Split(inputData, "\n") {
		numbers := strings.Split(line, "   ")

		listLeft = append(listLeft, utils.CastToInt(numbers[0]))
		listRight = append(listRight, utils.CastToInt(numbers[1]))
	}

	return listLeft, listRight
}