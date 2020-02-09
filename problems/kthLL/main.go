package main

import "fmt"

//Given a singly linked list and an integer k, remove the kth last element from the list. k is guaranteed to be smaller than the length of the list.
//The list is very long, so making more than one pass is prohibitively expensive.

type Node struct {
	Val  int
	Next *Node
}

func main() {
	head := llGen()
	remove(4, head)
	for head != nil {
		fmt.Printf("%d-->", head.Val)
		head = head.Next
	}
}

//note extra code would be required if k could be >= len of the LL
func remove(k int, head *Node) {

	var prev *Node
	kNode := head

	for i := 1; i < k; i++ {
		head = head.Next

	}

	for head.Next != nil {
		head = head.Next
		prev = kNode
		kNode = kNode.Next
	}

	prev.Next = kNode.Next
}

func llGen() *Node {
	head := &Node{
		Val: 1,
	}
	head.Next = &Node{
		Val: 2,
	}
	head.Next.Next = &Node{
		Val: 3,
	}
	head.Next.Next.Next = &Node{
		Val: 4,
	}
	head.Next.Next.Next.Next = &Node{
		Val: 5,
	}
	return head
}
