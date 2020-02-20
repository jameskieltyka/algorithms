package main

import (
	"container/heap"
	"fmt"
)

func main() {
	arr := []int{2, 4, 1, 3, 5}
	fmt.Println(outOfOrder(arr))

	arr = []int{5, 4, 3, 2, 1}
	fmt.Println(outOfOrder(arr))
}

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

func outOfOrder(arr []int) [][]int {
	h := &IntHeap{}
	heap.Init(h)
	heap.Push(h, arr[0])

	var res [][]int

	for i := 1; i < len(arr); i++ {
		tempStore := []int{}
		hLen := h.Len()
		temp := heap.Pop(h).(int)
		
		for hLen > 0 && temp > arr[i] {
			tempStore = append(tempStore, temp) 
			hLen = h.Len()
			if hLen > 0 {
				temp = heap.Pop(h).(int)
			}
		}

		heap.Push(h, arr[i])
		if len(tempStore) == 0 {
			heap.Push(h, temp)
			continue
		}

		for j := range tempStore {
			heap.Push(h, tempStore[j])
			res = append(res, []int{tempStore[j], arr[i]})
		}
	}
	
	return res
}

