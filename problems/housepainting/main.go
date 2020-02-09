package main

import (
	"fmt"
	"math"
)

//this was done with an inccorect assumption that the cost array index for each colour increments after it has been used

func main() {
	cost := [][]int{{2, 2, 3}, {3, 0, 0}, {3, 0, 0}}
	fmt.Println(minCost(cost))
}

func minCost(cost [][]int) int {
	k := len(cost)

	indexTrack := make([]int, k)
	return dfs(cost, indexTrack, -1, 0, 0)

}

func dfs(cost [][]int, index []int, lastSelected int, pos int, c int) int {
	if pos == len(index) {
		return c
	}
	min := math.MaxInt32
	for i := range index {
		if i == lastSelected {
			continue
		}
		index[i]++
		costPath := dfs(cost, index, i, pos+1, c+cost[i][index[i]-1])
		if costPath < min {
			min = costPath
		}
		index[i]--
	}
	return min
}
