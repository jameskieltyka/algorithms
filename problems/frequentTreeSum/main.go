package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	sums := make(map[int]int)
	maxCount := 0
	findSum(root, &maxCount, sums)

	maxKeys := []int{}
	for k, v := range sums {
		if v == maxCount {
			maxKeys = append(maxKeys, k)
		}
	}
	return maxKeys
}

func findSum(node *TreeNode, maxCount *int, sums map[int]int) int {
	if node == nil {
		return 0
	}

	total := findSum(node.Left, maxCount, sums) + findSum(node.Right, maxCount, sums) + node.Val
	sums[total]++
	if sums[total] > *maxCount {
		*maxCount = sums[total]
	}

	return total
}
