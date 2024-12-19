package utils

import (
	"os"
	"strings"
)

func GetInputDataFromFile(filePath string) string {
    inputData, err := os.ReadFile(filePath)

    if err != nil {
        panic("Error while reading file\n" + err.Error())
    }

    return string(strings.ReplaceAll(string(inputData), "\r", ""))
}
