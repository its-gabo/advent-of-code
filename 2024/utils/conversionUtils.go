package utils

import (
	"strconv"
)

func CastToInt(input string) int {
    number, err := strconv.Atoi(input)

    if err != nil {
        panic("Error converting string to int\n" + err.Error())
    }

    return number
}
