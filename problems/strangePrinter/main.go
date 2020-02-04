package main

func main() {

}

//https://leetcode.com/problems/strange-printer

func strangePrinter(s string) int {
	dp := make([][]int, len(s))

	for i := range s {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1

		if i+1 < len(s) {
			if s[i] == s[i+1] {
				dp[i][i+1] = 1
			} else {
				dp[i][i+1] = 2
			}
		}
	}

	for i := 2; i < len(s); i++ {
		for j := 0; j+i < len(s); j++ {
			dp[j][j+i] = i + 1
			for k := 0; k < i; k++ {
				temp := dp[j][j+k] + dp[j+k+1][j+i]
				if s[j+k] == s[j+i] {
					temp--
				}
				dp[j][j+i] = min(dp[j][j+i], temp)
			}
		}
	}

	if len(s) == 0 {
		return 0
	}
	return dp[0][len(s)-1]

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
