package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}

			if absDiff(sum, target) < absDiff(target, res) {
				res = sum
			}

			if sum < target {
				j++
			} else if sum > target {
				k--
			}
		}

	}

	return res
}

func absDiff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
