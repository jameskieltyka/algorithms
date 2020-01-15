package main

import "fmt"

import "sort"

func main() {
	candidate := []int{2, 3, 6, 7}
	fmt.Println(combinationSum(candidate, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}

	sort.Ints(candidates)

	combinationHelper(&res, candidates, target, []int{})
	return res
}

func combinationHelper(res *[][]int, candidates []int, target int, temp []int) {
	if target == 0 {
		t := make([]int, len(temp))
		copy(t, temp)
		*res = append(*res, t)
		return
	}

	for _, i := range candidates {
		if i <= target {
			if len(temp) > 0 && i < temp[len(temp)-1] {
				continue
			}
			temp = append(temp, i)
			combinationHelper(res, candidates, target-i, temp)
			temp = temp[0 : len(temp)-1]
		}
	}
}
