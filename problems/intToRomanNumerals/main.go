package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}

var charRoman map[int][]string = map[int][]string{0: []string{"I", "V"},
	1: []string{"X", "L"},
	2: []string{"C", "D"},
	3: []string{"M"}}

func intToRoman(num int) string {
	if num == 0 {
		return ""
	}

	return roman(num, 0)

}

func roman(num int, index int) string {
	if num == 0 {
		return ""
	}
	calc := num % 10
	num /= 10

	var res string
	if calc <= 3 {
		res += strings.Repeat(charRoman[index][0], calc)
	} else if calc <= 8 {
		for j := calc; j < 5; j++ {
			res += charRoman[index][0]
		}
		res += charRoman[index][1]
		for j := calc; j > 5; j-- {
			res += charRoman[index][0]
		}
	} else {
		for j := calc; j > 8; j-- {
			res += charRoman[index][0]
		}
		res += charRoman[index+1][0]
	}

	return roman(num, index+1) + res

}
