package main

import "fmt"

// Given a list of integers S and a target number k, write a function that returns a subset of S that adds up to k.
// If such a subset cannot be made, then return null.

func main() {
	var res []int
	arr := []int{12, 1, 61, 5, 9, 2}
	fmt.Println(subset(arr, 24, res))
}

func subset(arr []int, target int, res []int) []int {
	if target == 0 {
		return res
	}

	for i := range arr {
		if arr[i] == -1 {
			continue
		}
		temp := arr[i]
		res = append(res, arr[i])
		arr[i] = -1

		option := subset(arr, target-res[len(res)-1], res)
		if option != nil {
			return option
		}

		res = res[:len(res)-1]
		arr[i] = temp
	}

	return nil
}
