package main

import "fmt"

func main() {
	arr := []int{54, 25, 12, 22, 110}
	bubbleSort(arr)
	fmt.Println(arr)
}

func bubbleSort(arr []int) {
	for range arr {
		flipped := false
		for j := 1; j < len(arr); j++ {
			if arr[j-1] > arr[j] {
				flipped = true
				swap(arr, j-1, j)
			}
		}
		if !flipped {
			return
		}
	}
}

func swap(arr []int, i, j int) {
	arr[j], arr[i] = arr[i], arr[j]
}
