package main

import "fmt"

func main() {
	arr := []int{54, 25, 12, 22, 110}
	insertionSort(arr)
	fmt.Println(arr)
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			j := 0
			for arr[i] > arr[j] {
				j++
			}
			insert(arr, i, j)
		}
	}
}

func insert(arr []int, i, j int) {
	val := arr[i]
	copy(arr[j+1:], arr[j:i])
	arr[j] = val
}
