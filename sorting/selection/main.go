package main

import "fmt"

func main() {
	arr := []int{54, 25, 12, 22, 110}
	selectionSort(arr)
	fmt.Println(arr)
}

func selectionSort(arr []int) {
	for i := range arr {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		swap(arr, min, i)
	}
}

func swap(arr []int, i, j int) {
	arr[j], arr[i] = arr[i], arr[j]
}
