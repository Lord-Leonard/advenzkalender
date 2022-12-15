package main

import (
	"advenzkalender/utils"
	"fmt"
)

func main() {
	scanner := utils.GetScanner("misk/input_3rd_task")

	for scanner.Scan() {
		fmt.Printf("\n%v", scanner.Text())
	}

	
}
