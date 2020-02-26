package main

import (
	"fmt"
	"strconv"
)

func main() {
	root := Tree()
	fmt.Println(TreeValue(root))
}

type TreeNode struct {
	Val   string
	Left  *TreeNode
	Right *TreeNode
}

func TreeValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	switch root.Val {
	case "/":
		return TreeValue(root.Left) / TreeValue(root.Right)
	case "*":
		return TreeValue(root.Left) * TreeValue(root.Right)
	case "+":
		return TreeValue(root.Left) + TreeValue(root.Right)
	case "-":
		return TreeValue(root.Left) - TreeValue(root.Right)
	default:
		i, _ := strconv.Atoi(root.Val)
		return i
	}

	return 0
}

func Tree() *TreeNode {
	root := &TreeNode{
		Val: "*",
	}
	root.Left = &TreeNode{Val: "+"}
	root.Right = &TreeNode{Val: "+"}
	root.Left.Left = &TreeNode{Val: "3"}
	root.Left.Right = &TreeNode{Val: "2"}
	root.Right.Left = &TreeNode{Val: "4"}
	root.Right.Right = &TreeNode{Val: "5"}
	return root
}
