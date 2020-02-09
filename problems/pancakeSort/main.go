package main

import "fmt"

func main() {
	fmt.Println(pancakeSort([]int{3, 2, 4, 1}))
}

func pancakeSort(A []int) []int {
	var res []int
	if isSorted(A) {
		return res
	}

	max := len(A)
	for max != 1 {
		if A[max-1] == max {
			max--
			continue
		}
		for i := range A {
			if A[i] == max {
				res = append(res, i+1)
				flip(A, i)
				res = append(res, max)
				flip(A, max-1)
				break
			}
		}
		max--
	}

	return res
}

func flip(a []int, index int) {
	for i := 0; i <= index/2; i++ {
		a[i], a[index-i] = a[index-i], a[i]
	}
}

func isSorted(a []int) bool {
	last := 0
	for i := range a {
		if a[i] < last {
			return false
		}
		last = a[i]
	}
	return true
}
