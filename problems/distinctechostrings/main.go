package main

import "fmt"

func main() {
	fmt.Println(distinctEchoSubstrings("abcabcabc"))
	fmt.Println(distinctEchoSubstrings("leetcodeleetcode"))
}

func distinctEchoSubstrings(text string) int {
	res := make(map[string]bool)
	for i := 2; i <= len(text); i += 2 {
		for j := 0; j <= len(text)-i; j++ {
			if checkEcho(text[j : j+i]) {
				res[text[j:j+i]] = true
			}
		}
	}

	return len(res)
}

func checkEcho(str string) bool {
	length := len(str)
	if str[0:length/2] == str[length/2:length] {
		return true
	}
	return false
}
