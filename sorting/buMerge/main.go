package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{54, 25, 12, 22, 110}
	mergeSort(arr)
	fmt.Println(arr)
}

func mergeSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i = i + i {
		for lo := 0; lo < n-i; lo += i + i {
			merge(arr, lo, lo+i, int(math.Min(float64(lo+i+i), float64(n))))
		}
	}
}

func merge(arr []int, low, mid, high int) {
	temp := make([]int, high-low)
	lowT, midT := low, mid
	for count := 0; count < high-low; count++ {
		switch {
		case lowT >= mid:
			temp[count] = arr[midT]
			midT++
		case midT >= high:
			temp[count] = arr[lowT]
			lowT++
		case arr[lowT] < arr[midT]:
			temp[count] = arr[lowT]
			lowT++
		default:
			temp[count] = arr[midT]
			midT++
		}
	}

	copy(arr[low:high], temp)
}
