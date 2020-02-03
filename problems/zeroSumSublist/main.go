package main

///leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	history := make(map[int]*ListNode)

	temp := &ListNode{
		Next: head,
	}
	sum := 0
	head = temp
	for head != nil {
		sum += head.Val
		if n, ok := history[sum]; ok {
			tempN := n.Next
			tempSum := sum
			for tempN != head {
				tempSum += tempN.Val
				delete(history, tempSum)
				tempN = tempN.Next
			}
			n.Next = head.Next
			head = n
		} else {
			history[sum] = head
		}

		head = head.Next
	}
	return temp.Next
}
