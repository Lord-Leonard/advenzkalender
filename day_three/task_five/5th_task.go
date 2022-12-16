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
		rucksackInhalt := []byte(scanner.Text())
		rucksackGroesse := len(rucksackInhalt)
		compartementOne := rucksackInhalt[:rucksackGroesse/2]
		compartementTwo := rucksackInhalt[rucksackGroesse/2:]

		for _, value := range compartementOne {

			var2 := bytes.ContainsRune(compartementTwo, rune(value))
			if var2 {
				priority := day_three.DeterminePriority(value)

				prioritySum += int(priority)

			}
		}
	}
}
