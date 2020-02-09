package main

import "fmt"

import "math/rand"

func main() {
	arr := []int{12, 54, 25, 12, 22, 110, 110, 110, 1}
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quicksort(arr []int, low, high int) {
	if low >= high {
		return
	}

	index := partition(arr, low, high)
	quicksort(arr, low, index-1)
	quicksort(arr, index+1, high)
}

func partition(arr []int, low, high int) int {

	index := rand.Intn(high-low) + low
	arr[index], arr[high] = arr[high], arr[index]

	i, j := low, high-1
	for i <= j {
		switch {
		case arr[i] <= arr[high]:
			i++
		case arr[j] >= arr[high]:
			j--
		default:
			arr[j], arr[i] = arr[i], arr[j]
			j--
			i++
		}
	}
	arr[j+1], arr[high] = arr[high], arr[j+1]
	return j + 1
}
