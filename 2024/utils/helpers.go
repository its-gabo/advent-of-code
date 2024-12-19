package utils

import (
	"os"
	"strconv"
	"strings"

	gmp "github.com/zenpk/go-mod-path"
)

func GetCurrentModDirectory() string {
	currentDirectory, err := gmp.GetNearestPath()

	if err != nil {
		panic("Error while getting current directory\n" + err.Error())
	}

	return currentDirectory
}

func GetInputDataFromFile(filePath string) string {
	inputData, err := os.ReadFile(filePath)

	if err != nil {
		panic("Error while reading file\n" + err.Error())
	}

	return string(strings.ReplaceAll(string(inputData), "\r", ""))
}

func CastToInt(input string) int {
	number, err := strconv.Atoi(input)

	if err != nil {
		panic("Error converting string to int\n" + err.Error())
	}

	return number
}

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
