package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	sum := 0
	for i := range nums {
		if sum+nums[i] < nums[i] {
			sum = 0
		}
		sum += nums[i]
		if sum > maxSum {
			maxSum = sum
		}

	}
	return maxSum
}

//DP version
// func maxSubArray(nums []int) int {
// 	dp := make([]int, len(nums))
// 	dp[0] = nums[0]
//     max := nums[0]
// 	for i := 1; i<len(nums); i++ {
// 		if nums[i] + dp[i-1] > nums[i]{
// 			dp[i] = nums[i] + dp[i-1]
// 		} else {
// 			dp[i] = nums[i]
// 		}

//         if max < dp[i]{
//             max = dp[i]
//         }
// 	}
// 	return max
// }
