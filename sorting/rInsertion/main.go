package main

import "fmt"

func main() {
	arr := []int{54, 25, 12, 22, 110}
	insertionSort(arr, 1)
	fmt.Println(arr)
}

func insertionSort(arr []int, n int) {
	if n >= len(arr) {
		return
	}

	insertionSort(arr, n+1)

	for i := range arr {
		if i > n {
			break
		}
		if arr[i] > arr[n] {
			insert(arr, i, n)
		}
	}

}

func insert(arr []int, i int, j int) {
	x := arr[j]
	copy(arr[i+1:], arr[i:j])
	arr[i] = x
}
