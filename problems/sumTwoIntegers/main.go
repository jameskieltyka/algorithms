package main

import "fmt"

func main() {
	fmt.Println(getSum(-12, -8))
}

func getSum(a int, b int) int {
	res, co := a^b, a&b
	for co != 0 {
		co = co << 1
		res, co = res^co, res&co
	}
	return res
}
