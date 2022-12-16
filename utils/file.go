package utils

import (
	"bufio"
	"os"
)

func GetScanner(file string) *bufio.Scanner {
	f, err := os.Open(file)
	check(err)

	return bufio.NewScanner(f)
}
