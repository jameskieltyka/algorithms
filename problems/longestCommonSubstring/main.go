package main

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1))
	for i := range dp {
		dp[i] = make([]int, len(text2))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return lcs(text1, text2, len(text1)-1, len(text2)-1, dp)
}

func lcs(text1, text2 string, i, j int, dp [][]int) int {
	if i < 0 || j < 0 {
		return 0
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}

	if text1[i] == text2[j] {
		dp[i][j] = lcs(text1, text2, i-1, j-1, dp) + 1
	} else {
		dp[i][j] = max(lcs(text1, text2, i-1, j, dp), lcs(text1, text2, i, j-1, dp))
	}

	return dp[i][j]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
