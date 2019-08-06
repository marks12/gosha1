package project

import (
	"os"
	"strings"
)

func GetName() (title string) {

	path, _ := os.Getwd()

	pathArr := strings.Split(path, "/")

	l := len(pathArr)

	if l > 0 {
		l--
	}

	title = pathArr[l]

	return strings.Title(title)
}
