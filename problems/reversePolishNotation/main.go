package main

import (
	"fmt"
	"strconv"
)

func main() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens))
}

func evalRPN(tokens []string) int {
	res := 0
	var numStack []int
	for i := range tokens {
		switch tokens[i] {
		case "*":
			a, b := numStack[len(numStack)-1], numStack[len(numStack)-2]
			res = a * b
			numStack = numStack[0 : len(numStack)-1]
			numStack[len(numStack)-1] = res
		case "/":
			a, b := numStack[len(numStack)-1], numStack[len(numStack)-2]
			res = b / a
			numStack = numStack[0 : len(numStack)-1]
			numStack[len(numStack)-1] = res
		case "+":
			a, b := numStack[len(numStack)-1], numStack[len(numStack)-2]
			res = a + b
			numStack = numStack[0 : len(numStack)-1]
			numStack[len(numStack)-1] = res
		case "-":
			a, b := numStack[len(numStack)-1], numStack[len(numStack)-2]
			res = b - a
			numStack = numStack[0 : len(numStack)-1]
			numStack[len(numStack)-1] = res
		default:
			num, _ := strconv.Atoi(tokens[i])
			numStack = append(numStack, num)
		}
	}

	return numStack[len(numStack)-1]
}
