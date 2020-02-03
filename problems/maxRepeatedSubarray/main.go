package main

import "fmt"

func main() {
	fmt.Println(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
}

func findLength(A []int, B []int) int {
	res := 0
	if len(A) == 0 || len(B) == 0 {
		return res
	}

	aLen := len(A)
	bLen := len(B)
	dp := make([][]int, aLen+1)
	for i := range dp {
		dp[i] = make([]int, bLen+1)
	}
	for i := aLen - 1; i >= 0; i-- {
		for j := bLen - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			}
			if res < dp[i][j] {
				res = dp[i][j]
			}
		}
	}
	return res
}
