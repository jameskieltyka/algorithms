package main

func minDistance(word1 string, word2 string) int {
    dp := make([][]int, len(word1))
	for i := range dp {
		row := make([]int, len(word2))
		for j := range row {
			row[j] = -1
		}
		dp[i] = row
	}
    return editDistance(word1, word2, 0, 0, dp)
}

func editDistance(s string, target string, i, j int, dp [][]int) int {
	if len(s) == i {
		return len(target) - j
	}
	if len(target) == j {
		return len(s) - i
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}
	if s[i] == target[j] {
		dp[i][j] = editDistance(s, target, i+1, j+1, dp)
	} else {
		sW := editDistance(s, target, i+1, j+1, dp)
		sA := editDistance(s, target, i, j+1, dp)
		sD := editDistance(s, target, i+1, j, dp)

		dp[i][j] = min(sA, min(sD, sW)) + 1
	}

	return dp[i][j]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
