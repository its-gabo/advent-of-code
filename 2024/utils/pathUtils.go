package utils

import (
	gmp "github.com/zenpk/go-mod-path"
)

func GetCurrentModDirectory() string {
    currentDirectory, err := gmp.GetNearestPath()

    if err != nil {
        panic("Error while getting current directory\n" + err.Error())
    }

    return currentDirectory
}
