package main

import (
	"advenzkalender/utils"
	"fmt"
	"strconv"
)

func main() {
	scanner := utils.GetScanner("misk/input_1st_task")
	var calorieList []int
	calorieList = append(calorieList, 0)
	elvIndex := 0
	for scanner.Scan() {

		if scanner.Text() == "" {
			calorieList = append(calorieList, 0)
			elvIndex++
			continue
		}
		calorie, _ := strconv.Atoi(scanner.Text())
		calorieList[elvIndex] = calorieList[elvIndex] + calorie
	}

	fmt.Printf("%v", calorieList)

	var m int // 1st
	var n int // 2nd
	var o int // 3rd
	var t int // temp
	for _, i := range calorieList {
		if i > m {
			t = m
			m = i
			i = t
		}
		if i > n {
			t = n
			n = i
			i = t
		}
		if i > o {
			o = i
		}
	}

	fmt.Printf("\nBiggest Calories: %v \n", m)

	fmt.Printf("1. %v \n2. %v \n3. %v \n", m, n, o)
	fmt.Printf("total: %v \n", m+n+o)
}
