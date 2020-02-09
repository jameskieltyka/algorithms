package main

import "fmt"

//https://leetcode.com/problems/triangle/

//TOP DOWN
func minimumTotalTD(triangle [][]int) int {
	memo := make(map[string]int)
	return minSubTotal(triangle, 0, 0, memo)
}

func minSubTotal(triangle [][]int, height int, index int, memo map[string]int) int {
	if height == len(triangle)-1 {
		return triangle[height][index]
	}
	key := fmt.Sprintf("%d,%d", height, index)
	if n, ok := memo[key]; ok {
		return n
	}

	res := triangle[height][index] + min(minSubTotal(triangle, height+1, index, memo), minSubTotal(triangle, height+1, index+1, memo))
	memo[key] = res
	return res
}

//BOTTOM UP
func minimumTotalBU(triangle [][]int) int {
	n := len(triangle) - 1
	cache := triangle[n]
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			cache[j] = triangle[i][j] + min(cache[j], +cache[j+1])
		}
	}
	return cache[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
