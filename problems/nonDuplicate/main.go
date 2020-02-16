package main

import "fmt"

// Given an array of integers where every integer occurs three times except for one integer, which only occurs once, find and return the non-duplicated integer.
// For example, given [6, 1, 3, 3, 3, 6, 6], return 1. Given [13, 19, 13, 13], return 19.
// Do this in O(N) time and O(1) space.

func main() {
	fmt.Println(nonDuplicate([]int{6, 1, 3, 3, 3, 6, 6}))
	fmt.Println(nonDuplicate([]int{13, 19, 13, 13}))
	fmt.Println(nonDuplicate([]int{1, 2, 2, 3, 1, 2, 1}))

}

func nonDuplicate(arr []int) int {
	ones := 0
	twos := 0
	for i := range arr {
		twos = twos | (ones & arr[i])
		ones = ones ^ arr[i]
		clearMask := ^(ones & twos)
		ones &= clearMask
		twos &= clearMask
	}

	return ones
}
