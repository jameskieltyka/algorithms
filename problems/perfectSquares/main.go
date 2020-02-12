package main

import "math"

//https://leetcode.com/problems/perfect-squares

func numSquares(n int) int {

	memo := make([]int, n+1)

	return minSquares(n, memo)
}

func minSquares(n int, memo []int) int {
	if n == 0 {
		return 0
	}

	if memo[n] != 0 {
		return memo[n]
	}

	min := math.MaxInt32
	for j := 1; j*j <= n; j++ {
		res := minSquares(n-(j*j), memo) + 1
		if res < min {
			min = res
		}
	}

	memo[n] = min
	return min
}

// //bottom up
// func numSquares(n int) int {

//     dp := make([]int, n+1)
//     for i := range dp{
//         dp[i] = math.MaxInt32
//     }

//     dp[0] = 0
//     for i := 1; i <=n; i++ {
//         for j:= 1; i - j*j >= 0; j++{
//             dp[i] = min(dp[i], dp[i-j*j]+1)
//         }
//     }

//     return dp[n]
// }

// func min(a, b int) int {
//     if a > b {
//         return b
//     }
//     return a
// }
