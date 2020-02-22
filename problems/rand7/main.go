package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand7())
}

func rand7() int {
	sum := 21
	for sum >= 20 {
		sum = rand5() + rand5()*5
	}
	return sum % 7
}

func rand5() int {
	return rand.Intn(5)
}
