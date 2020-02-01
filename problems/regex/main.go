package main

import "fmt"

//leetcode.com/problems/regular-expression-matching
// Implement regular expression matching with the following special characters:
// •	. (period) which matches any single character
// •	* (asterisk) which matches zero or more of the preceding element
// That is, implement a function that takes in a string and a valid regular expression and returns whether or not the string matches the regular expression.

func isMatch(s string, p string) bool {
	memo := make(map[string]bool)
	return dp(0, 0, s, p, memo)
}

func dp(i, j int, s, p string, memo map[string]bool) bool {
	key := fmt.Sprintf("%d,%d", i, j)
	if _, ok := memo[key]; ok {
		return memo[key]
	}

	if j == len(p) {
		if i == len(s) {
			memo[key] = true
			return true
		}
		memo[key] = false
		return false
	}

	var firstMatch bool
	if i < len(s) && (s[i] == p[j] || p[j] == '.') {
		firstMatch = true
	}

	if j+1 < len(p) && p[j+1] == '*' {
		memo[key] = (dp(i, j+2, s, p, memo) || (firstMatch && dp(i+1, j, s, p, memo)))
		return memo[key]
	}

	memo[key] = firstMatch && dp(i+1, j+1, s, p, memo)
	return memo[key]

}
