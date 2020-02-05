package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(encode("AAAABBBCCDAA"))
	fmt.Println(decode("4A3B2C1D2A"))
}

func encode(s string) string {
	lastChar := ' '
	counter := 0
	var sb strings.Builder
	for i, c := range s {
		if i > 0 && c != lastChar {
			sb.WriteString(strconv.Itoa(counter))
			sb.WriteRune(lastChar)
			counter = 0
		}
		lastChar = c
		counter++
	}
	sb.WriteString(strconv.Itoa(counter))
	sb.WriteRune(lastChar)
	return sb.String()
}

func decode(s string) string {
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		var sNum strings.Builder
		for unicode.IsDigit(rune(s[i])) {
			sNum.WriteRune(rune(s[i]))
			i++
		}
		num, _ := strconv.Atoi(sNum.String())
		sb.WriteString(strings.Repeat(string(s[i]), num))
	}
	return sb.String()
}
