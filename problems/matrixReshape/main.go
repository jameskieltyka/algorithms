package main

import "fmt"

func main() {
	nums := [][]int{{4, 3}, {2, 1}}
	fmt.Println(matrixReshape(nums, 4, 1))
}

map[string]

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if r*c != len(nums)*len(nums[0]) {
		return nums
	}

	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, 0, c)
	}

	row := 0

	for i := range nums {
		for j := range nums[i] {
			if len(res[row]) >= c {
				row++
			}
			res[row] = append(res[row], nums[i][j])
		}
	}

	return res
}
