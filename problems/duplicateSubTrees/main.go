package main

import "strconv"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	subTrees := make(map[string]int)
	var res []*TreeNode
	transitTree(root, subTrees, &res)
	return res
}

func transitTree(node *TreeNode, trees map[string]int, res *[]*TreeNode) string {
	if node == nil {
		return ""
	}

	key := transitTree(node.Left, trees, res) + "." + transitTree(node.Right, trees, res) + "." + strconv.Itoa(node.Val)

	if n, ok := trees[key]; ok && n == 1 {
		*res = append(*res, node)
	}
	trees[key]++

	return key
}
