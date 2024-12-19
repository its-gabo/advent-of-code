package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

var inputData string

func init() {
	currentDirectory := utils.GetCurrentModDirectory()
	inputData = utils.GetInputDataFromFile(currentDirectory + "/day-02/dayTwo.txt")
}

func main() {
	fmt.Println("Part 1 answer: ", partOne())
	fmt.Println("Part 2 answer: ", partTwo())
}

func partOne() int {
	rows := parseInput(inputData)

	safeRowsCounter := 0

	for _, row := range rows {
		if isRowSafe(row) {
			safeRowsCounter++
		}
	}
		
	return safeRowsCounter;
}

func partTwo() int {
	rows := parseInput(inputData)

	safeRowsCounter := 0

	for _, row := range rows {
		if isRowSafeWithErrorMargin(row) {
			safeRowsCounter++
		}
	}
		
	return safeRowsCounter;
}

func isRowSafe(row []int) bool {
	isIncreasing, isDecreasing := false, false

	for i := 1; i < len(row); i++ {
		diff := row[i] - row[i - 1]

		if diff > 0 {
			isIncreasing = true
		} else {
			isDecreasing = true
		}

		if utils.AbsInt(diff) > 3 || diff == 0 || (isDecreasing && isIncreasing) {
			return false
		}
	}

	return true
}


func isRowSafeWithErrorMargin(row []int) bool {
	for i := 0; i < len(row); i++ {
		rowCopy := make([]int, len(row))
		copy(rowCopy, row)
		
		if i == len(row) - 1 {
			rowCopy = rowCopy[:i]
		} else {
			rowCopy = append(rowCopy[:i], rowCopy[i + 1:]...)
		}
		
		if isRowSafe(rowCopy) {
			return true
		}
	}

	return false
}

func parseInput(inputData string) (rows [][]int) {
	rowsString := strings.Split(inputData, "\n")

	rows = make([][]int, len(rowsString))

	for i, row := range rowsString {
		for _, number := range strings.Split(row, " ") {
			rows[i] = append(rows[i], utils.CastToInt(number))
		}
	}

	return rows
}