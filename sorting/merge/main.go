package main

import "fmt"

func main() {
	arr := []int{54, 25, 12, 22, 110}
	mergeSort(arr, 0, len(arr))
	fmt.Println(arr)
}

func mergeSort(arr []int, i, n int) {
	if n <= i {
		return
	}
	mid := (i + n) / 2
	mergeSort(arr, i, mid)
	mergeSort(arr, mid+1, n)
	merge(arr, i, mid, n)
}

func merge(arr []int, low, mid, high int) {
	temp := make([]int, high-low)
	count := 0
	lowT, midT := low, mid
	for lowT < mid && midT < high {
		if arr[lowT] < arr[midT] {
			temp[count] = arr[lowT]
			lowT++
		} else {
			temp[count] = arr[midT]
			midT++
		}
		count++
	}

	for lowT < mid {
		temp[count] = arr[lowT]
		lowT++
		count++
	}

	for midT < high {
		temp[count] = arr[midT]
		midT++
		count++
	}

	copy(arr[low:high], temp)
}
