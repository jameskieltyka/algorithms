package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(completeSet(nums))
}

func completeSet(arr []int) [][]int {
	var res [][]int
	res = append(res, []int{})

	for i := range nums {
		for _, item := range res {
			res = append(res, append([]int{nums[i]}, item...))
		}
	}
	return res
}
