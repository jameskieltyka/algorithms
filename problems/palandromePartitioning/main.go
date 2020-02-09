package main

import "fmt"

//https://leetcode.com/problems/palindrome-partitioning/

func main() {
	fmt.Println(partition("aaa"))
}

func partition(s string) [][]string {
	res := [][]string{}
	hasPalandrome(s, 0, &res, []string{})
	return res
}

func hasPalandrome(s string, l int, res *[][]string, temp []string) {
	if l >= len(s) {
		t := make([]string, len(temp))
		copy(t, temp)
		*res = append(*res, t)
	}

	for i := l + 1; i <= len(s); i++ {
		if isPalandrome(s[l:i]) {
			temp = append(temp, s[l:i])
			hasPalandrome(s, i, res, temp)
			temp = temp[0 : len(temp)-1]
		}
	}
}

func isPalandrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
