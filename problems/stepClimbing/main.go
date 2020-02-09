package main

import "fmt"

//There exists a staircase with N steps, and you can climb up either 1 or 2 steps at a time. Given N, write a function that returns the number of unique ways you can climb the staircase. The order of the steps matters.

func main() {
	memo := make(map[int]int)
	memo[0] = 1
	fmt.Println(climb(50, memo))
}

var availableSteps = []int{1, 2}

func climb(target int, memo map[int]int) int {
	if v, ok := memo[target]; ok {
		return v
	}

	res := 0
	for _, s := range availableSteps {
		if target >= s {
			res += climb(target-s, memo)
		}
	}

	memo[target] = res
	return res
}
