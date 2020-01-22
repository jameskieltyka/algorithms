package main

import "fmt"

import "sort"

func main() {
	root := tree()
	fmt.Println(verticalTraversal(root))
}

//For each node at position (X, Y), its left and right children respectively will be at positions (X-1, Y-1) and (X+1, Y-1).
//Running a vertical line from X = -infinity to X = +infinity, whenever the vertical line touches some nodes, we report the values of the nodes in order from top to bottom (decreasing Y coordinates).

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalTraversal(root *TreeNode) [][]int {
	res := make(map[int][]int)
	min := addEntry(root, res, 0)
	lenRes := len(res)
	vTrav := make([][]int, lenRes)
	for i, v := range res {
		v = sort.IntSlice(v)
		vTrav[i-min] = v
	}
	return vTrav
}

func addEntry(t *TreeNode, res map[int][]int, level int) int {
	if t == nil {
		return level + 1
	}
	if _, ok := res[level]; ok {
		res[level] = append(res[level], t.Val)
	} else {
		res[level] = []int{t.Val}
	}

	minR := addEntry(t.Right, res, level+1)
	minL := addEntry(t.Left, res, level-1)

	if minR < minL {
		return minR
	}
	return minL
}

func tree() *TreeNode {
	root := &TreeNode{
		Val: 3,
	}

	left := &TreeNode{
		Val: 9,
	}

	right := &TreeNode{
		Val: 20,
	}
	root.Left = left
	root.Right = right

	rightLeft := &TreeNode{
		Val: 15,
	}

	rightRight := &TreeNode{
		Val: 15,
	}

	root.Right.Left = rightLeft
	root.Right.Right = rightRight
	return root
}
