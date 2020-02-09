package main

import (
	"strings"
)

//https://leetcode.com/problems/ones-and-zeroes

func findMaxForm(strs []string, m int, n int) int {
	zeros := make([]int, len(strs))
	ones := make([]int, len(strs))
	for i, s := range strs {
		zeros[i] = strings.Count(s, "0")
		ones[i] = strings.Count(s, "1")
	}

	var dp = make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for s := range strs {
		for i := m; i >= zeros[s]; i-- {
			for j := n; j >= ones[s]; j-- {
				dp[i][j] = max(1+dp[i-zeros[s]][j-ones[s]], dp[i][j])
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
