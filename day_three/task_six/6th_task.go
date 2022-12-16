package main

import (
	"advenzkalender/day_three"
	"advenzkalender/utils"
	"bytes"
)

func main() {
	scanner := utils.GetScanner("misk/input_3rd_task")
	var prioritySum int

	for scanner.Scan() {
		rucksackOne := []byte(scanner.Text())
		scanner.Scan()
		rucksackTwo := []byte(scanner.Text())
		scanner.Scan()
		rucksackThree := []byte(scanner.Text())

		for _, inhalt := range rucksackOne {
			if bytes.ContainsAny(rucksackTwo, string(inhalt)) && bytes.ContainsAny(rucksackThree, string(inhalt)) {
				println(inhalt)
				priority := day_three.DeterminePriority(inhalt)

				prioritySum += int(priority)
				break
			}
		}
	}

	println(prioritySum)

}
