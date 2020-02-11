package main

import "fmt"

func main() {
	// fmt.Println(shortestPalindrome("aacecaaa"))
	// fmt.Println(shortestPalindrome("aacdedcaa"))
	// fmt.Println(shortestPalindrome("abcd"))
	// fmt.Println(shortestPalindrome("cabcbds"))
	fmt.Println(shortestPalindrome("adcba"))
}

func shortestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	largestEnd := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[largestEnd] == s[i] {
			largestEnd++
		}
	}

	add := []rune(s)[largestEnd:]
	for i := 0; i <= len(add)/2; i++ {
		add[i], add[len(add)-1-i] = add[len(add)-1-i], add[i]
	}

	if 

	return string(add) + shortestPalindrome(s[:largestEnd]) + s[largestEnd:]

}
