package main

import "fmt"

func main() {
	fmt.Println(trappedVolume([]int{2, 1, 2}))
	fmt.Println(trappedVolume([]int{3, 0, 1, 3, 0, 5}))
	fmt.Println(trappedVolume([]int{0, 0, 4, 0, 0, 6, 0, 0, 3, 0, 5, 0, 1, 0, 0, 0}))
}

func trappedVolume(height []int) int {
	leftMax := -1
	rightMax := -1
	res := 0
	left, right := 0, len(height)-1
	for left <= right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				res += rightMax - height[right]
			}
			right--
		}
	}

	return res
}
