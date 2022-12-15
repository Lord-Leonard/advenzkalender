package utils

import (
	"bufio"
	"os"
)

func GetScanner(file string) *bufio.Scanner {
	f, err := os.Open(file)
	check(err)

	var calorieList []int
	calorieList = append(calorieList, 0)

	return bufio.NewScanner(f)
}
