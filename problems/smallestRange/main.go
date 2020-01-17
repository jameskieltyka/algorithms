package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{1, 3, 6}
	fmt.Println(smallestRangeII(ints, 3))
}

func smallestRangeII(A []int, K int) int {
	max := len(A)
	sort.Ints(A)

	diff := A[max-1] - A[0]

	for i := 0; i < max-1; i++ {
		a, b := A[i], A[i+1]
		high := maxInt(A[max-1]-K, a+K)
		low := minInt(A[0]+K, b-K)
		diff = minInt(diff, high-low)
	}

	return diff
}

func maxInt(i int, j int) int {
	if i < j {
		return j
	}
	return i
}

func minInt(i int, j int) int {
	if i < j {
		return i
	}
	return j
}
