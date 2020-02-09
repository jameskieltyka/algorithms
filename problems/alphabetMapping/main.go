package main

import "fmt"

import "strconv"

//Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the number of ways it can be decoded.
//For example, the message '111' would give 3, since it could be decoded as 'aaa', 'ka', and 'ak'.
//You can assume that the messages are decodable. For example, '001' is not allowed.

func main() {
	fmt.Println(uniqueMessages("111"))
}

func uniqueMessages(s string) int {
	runes := []rune(s)
	return uniqueMessage(runes, 0)
}

func uniqueMessage(runes []rune, pos int) int {
	if pos >= len(runes) {
		return 1
	}

	var count = 0
	if runes[pos] != 0 {
		count += uniqueMessage(runes, pos+1)
	}

	if pos+1 < len(runes) {
		//consider changing this to a map for faster operation
		num, err := strconv.Atoi(string(runes[pos : pos+2]))
		if err != nil {
			return 1
		}

		if num > 0 && num < 27 {
			count += uniqueMessage(runes, pos+2)
		}

	}
	return count
}
