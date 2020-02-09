package main

//https://leetcode.com/problems/word-break/

func wordBreak(s string, wordDict []string) bool {
	memo := make(map[string]bool)
	return breakS(s, wordDict, memo)
}

func breakS(s string, wordDict []string, memo map[string]bool) bool {
	if s == "" {
		return true
	}

	if res, ok := memo[s]; ok {
		return res
	}

	for _, word := range wordDict {
		if len(word) <= len(s) && s[0:len(word)] == word {
			if breakS(s[len(word):len(s)], wordDict, memo) {
				memo[s] = true
				return true
			}
		}
	}
	memo[s] = false
	return false
}
