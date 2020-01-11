package main

import (
	"fmt"
	"strings"
)

type node struct {
	val   string
	left  *node
	right *node
}

//Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.
func main() {
	root := tree()
	var s strings.Builder
	serialise(root, &s)
	fmt.Println(s.String())

	rootDeserialised := deserialise(s.String())
	fmt.Println(rootDeserialised.right.left.val)
}

func serialise(n *node, s *strings.Builder) {
	if n == nil {
		s.WriteString("nil,")
		return
	}

	s.WriteString(fmt.Sprintf("%s,", n.val))
	serialise(n.left, s)
	serialise(n.right, s)
}

func deserialise(s string) *node {
	strings := strings.Split(s, ",")
	root, _ := deserialiseNode(strings, 0)
	return root
}

func deserialiseNode(strs []string, index int) (*node, int) {
	if strs[index] == "nil" {
		return nil, index
	}
	node := &node{
		val: strs[index],
	}

	node.left, index = deserialiseNode(strs, index+1)
	node.right, index = deserialiseNode(strs, index+1)
	return node, index
}

func tree() *node {
	root := &node{
		val: "test",
	}

	left := &node{
		val: "left 1",
	}

	right := &node{
		val: "right 1",
	}

	rightl := &node{
		val: "right left",
	}

	rightr := &node{
		val: "right right",
	}
	right.left = rightl
	right.right = rightr
	root.left = left
	root.right = right

	return root
}
