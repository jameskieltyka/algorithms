package main

import "fmt"

//Given a string containing only three types of characters: '(', ')' and '*', write a function to check whether this string is valid. We define the validity of a string by these rules:

//Any left parenthesis '(' must have a corresponding right parenthesis ')'.
//Any right parenthesis ')' must have a corresponding left parenthesis '('.
//Left parenthesis '(' must go before the corresponding right parenthesis ')'.
//'*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string.
//An empty string is also valid.

func main() {
	fmt.Println(checkValidString("(())(())(((()*()()()))()((()()(*()())))(((*)()"))
}

// "(((*()())))(((*)"
// "(())(())(((()*()()()))()((()()(*()())))(((*)()"
// "11221112212121*21*121122122122121211121221112221*"
// "1*"

func checkValidString(s string) bool {
	if s == "" {
		return true
	}

	lowestcount := 0
	highestcount := 0
	for _, r := range s {
		if r == '(' {
			lowestcount++
			highestcount++
		} else if r == ')' {
			lowestcount--
			highestcount--
		} else {
			lowestcount--
			highestcount++
		}

		if highestcount < 0 {
			return false
		}
		if lowestcount < 0 {
			lowestcount = 0
		}
	}

	return lowestcount <= 0

}
