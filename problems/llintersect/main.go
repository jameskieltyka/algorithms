package main

//Given two singly linked lists that intersect at some point, find the intersecting node. The lists are non-cyclical.

type node struct {
	val  int
	next *node
}

func intersect(root1 *node, root2 *node) int {
	length1, length2 := 0, 0
	n1 := root1
	for n1.next != nil {
		length1++
		n1 = n1.next
	}

	n2 := root2
	for n2.next != nil {
		length2++
		n2 = n2.next
	}

	if length1 > length2 {
		for i := 0; i < length1-length2; i++ {
			root1 = root1.next
		}
	} else {
		for i := 0; i < length2-length1; i++ {
			root2 = root2.next
		}
	}

	for root1.val != root2.val {
		root1 = root1.next
		root2 = root2.next
	}

	return root2.val
}
