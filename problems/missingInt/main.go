package main

import "fmt"

//Given an array of integers, find the first missing positive integer
//in linear time and constant space. In other words, find the lowest positive integer
//that does not exist in the array.
//The array can contain duplicates and negative numbers as well.

func main() {
	numbers := []int{3, 4, -1, 1, 2}

	fmt.Println(missing(numbers))

	numbers = []int{6, 1, 3, 4, 7, 2}
	fmt.Println(missing(numbers))

	numbers = []int{1, 2, 0}
	fmt.Println(missing(numbers))

	numbers = []int{1, 2}
	fmt.Println(missing(numbers))
}

func missing(arr []int) int {
	for i := range arr {
		for arr[i] > 0 && arr[i] < len(arr) && arr[i] != arr[arr[i]-1] {
			arr[i], arr[arr[i]-1] = arr[arr[i]-1], arr[i]
		}
	}

	fmt.Println(arr)
	for i := range arr {
		if arr[i] != i+1 {
			return i + 1
		}
	}

	return len(arr) + 1
}
