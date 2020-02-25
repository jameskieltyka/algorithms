package main

import "fmt"

//Given an array of numbers, find the maximum sum of any contiguous subarray of the array.
func main() {
	arr := []int{34, -50, 42, 14, -5, 86}
	fmt.Println(maxSubarray(arr))

	arr = []int{-5, -1, -8, -9}
	fmt.Println(maxSubarray(arr))
}

func maxSubarray(arr []int) int {
	max := 0
	current := 0
	for i := range arr {
		current += arr[i]
		if current < arr[i] {
			current = arr[i]
		}

		if current > max {
			max = current
		}
	}

	return max
}