package main

import (
	"advenzkalender/utils"
	"fmt"
	"strings"
)

func main() {
	m := make(map[string]int)
	m["A"] = 1 //Rock
	m["B"] = 2 //Paper
	m["C"] = 3 //Scissor

	m["X"] = 0 //Lose
	m["Y"] = 3 //Draw
	m["Z"] = 6 //Win

	m["lose"] = 0
	m["draw"] = 3
	m["win"] = 6

	scanner := utils.GetScanner("misk/input_2nd_task")
	var points int
	for scanner.Scan() {
		roundSlice := strings.Split(scanner.Text(), " ")

		//gameResult := determineWinner(m, roundSlice)
		myChoice := determineChoice(m, roundSlice[0], roundSlice[1])
		points += determinePoints(m, roundSlice[1], myChoice)

		fmt.Printf("\n%v", myChoice)

	}
	fmt.Printf("\n%v", points)
}

func determinePoints(m map[string]int, result string, myChoice string) int {
	return m[myChoice] + m[result]
}

func determinePointsNew(m map[string]int, result string, myChoice string) int {
	return m[myChoice]
}

func determineChoice(m map[string]int, opponentChoice string, forecast string) string {
	switch m[forecast] {
	case 0: //lose
		switch m[opponentChoice] {
		case 1: // Rock
			return "C"
		case 2: // Paper
			return "A"
		case 3: // Scissor
			return "B"

		}
	case 3: //draw
		return opponentChoice
	case 6: //win
		switch m[opponentChoice] {
		case 1: // Rock
			return "B"
		case 2: // Paper
			return "C"
		case 3: // Scissor
			return "A"
		}
	}
	panic("unable to determine Choice")
}

func determineWinner(m map[string]int, slice []string) string {
	opponentChoice := slice[0]
	myChoice := slice[1]

	// draw
	if m[opponentChoice] == m[myChoice] {
		return "draw"
	}

	if m[opponentChoice] == 1 && m[myChoice] == 2 {
		return "win"
	}
	if m[opponentChoice] == 1 && m[myChoice] == 3 {
		return "lose"
	}
	if m[opponentChoice] == 2 && m[myChoice] == 1 {
		return "lose"
	}
	if m[opponentChoice] == 2 && m[myChoice] == 3 {
		return "win"
	}
	if m[opponentChoice] == 3 && m[myChoice] == 1 {
		return "win"
	}
	if m[opponentChoice] == 3 && m[myChoice] == 2 {
		return "lose"
	}

	panic("no winner determined")
}
