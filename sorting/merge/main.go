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
