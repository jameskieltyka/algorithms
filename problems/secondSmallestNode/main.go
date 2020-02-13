package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	res := &smallest{one: -1, two: -1}
	smallestTwo(root, res)
	return res.two
}

type smallest struct {
	one int
	two int
}

func smallestTwo(node *TreeNode, res *smallest) {
	if node == nil {
		return
	}

	smallestTwo(node.Left, res)
	if res.one == -1 || node.Val < res.one {
		res.two = res.one
		res.one = node.Val
	} else if res.one != node.Val && (node.Val < res.two || res.two == -1) {
		res.two = node.Val
	}

	smallestTwo(node.Right, res)
}

func tree() *node {
	root := &node{
		val: 1,
	}

	left := &node{
		val: 1,
	}

	right := &node{
		val: 2,
	}

	rightl := &node{
		val: 1,
	}

	rightr := &node{
		val: 3,
	}
	right.left = rightl
	right.right = rightr
	root.left = left
	root.right = right

	return root
}
