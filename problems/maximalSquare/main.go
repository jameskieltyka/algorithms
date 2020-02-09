package main

//https://leetcode.com/problems/maximal-square

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	dp := make([][]int, len(matrix)+1)
	for i := range dp {
		dp[i] = make([]int, len(matrix[0])+1)
	}

	res := 0
	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[0]); j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(min(dp[i-1][j-1], dp[i][j-1]), dp[i-1][j]) + 1
				if res < dp[i][j] {
					res = dp[i][j]
				}
			}
		}
	}

	math.Pow

	return res * res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
