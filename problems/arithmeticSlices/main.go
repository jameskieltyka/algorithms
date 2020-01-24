package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, 4, 5, 0, 1, 2}
	fmt.Println(numberOfArithmeticSlices(ints))
}

func numberOfArithmeticSlices(A []int) int {
	res := 0
	if len(A) < 3 {
		return res
	}

	current := 2
	diff := A[1] - A[0]

	for i := 2; i < len(A); i++ {
		temp := A[i] - A[i-1]

		if diff == temp {
			current++
			if current >= 3 {
				res += (current - 2)
			}
		} else {
			current = 2
			diff = temp
		}
	}

	return res
}
