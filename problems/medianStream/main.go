package main

import (
	"container/heap"
	"fmt"
)

//Compute the running median of a sequence of numbers. That is, given a stream of numbers, print out the median of the list so far on each new element.

func main() {
	stream := []int{2, 1, 5, 7, 2, 0, 5}
	hLower := &MaxHeap{}
	hUpper := &MinHeap{}
	heap.Init(hUpper)
	heap.Init(hLower)
	median := &MedianStream{
		lower: hLower,
		upper: hUpper,
	}

	for i := range stream {
		fmt.Println(median.Add(stream[i]))
	}
}

type MedianStream struct {
	lower *MaxHeap
	upper *MinHeap
}

func (m MedianStream) Add(i int) float32 {
	if len(*m.lower) == len(*m.upper) {
		if len(*m.lower) > 0 {
			lower := heap.Pop(m.lower).(int)
			upper := heap.Pop(m.upper).(int)
			if i < lower {
				heap.Push(m.lower, i)
				heap.Push(m.upper, lower)
				heap.Push(m.upper, upper)
				return float32(lower)
			} else if i > upper {
				heap.Push(m.lower, lower)
				heap.Push(m.upper, i)
				heap.Push(m.upper, upper)
				return float32(upper)
			} else {
				heap.Push(m.lower, lower)
				heap.Push(m.upper, i)
				heap.Push(m.upper, upper)
				return float32(i)
			}
		}
		heap.Push(m.upper, i)
		return float32(i)
	} else if len(*m.lower) > len(*m.upper) {
		lower := heap.Pop(m.lower).(int)
		if i < lower {
			heap.Push(m.lower, i)
			heap.Push(m.upper, lower)
			return m.PopMiddle()
		} else {
			heap.Push(m.lower, lower)
			heap.Push(m.upper, i)
			return m.PopMiddle()
		}
	} else {
		upper := heap.Pop(m.upper).(int)
		if i < upper {
			heap.Push(m.lower, i)
			heap.Push(m.upper, upper)
			return m.PopMiddle()
		} else {
			heap.Push(m.lower, upper)
			heap.Push(m.upper, i)
			return m.PopMiddle()
		}
	}
	return float32(0)
}

func (m MedianStream) PopMiddle() float32 {
	lower := heap.Pop(m.lower).(int)
	upper := heap.Pop(m.upper).(int)
	heap.Push(m.lower, lower)
	heap.Push(m.upper, upper)
	return float32((upper + lower)) / 2
}

type MaxHeap []int

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() interface{} {
	temp := *m
	res := temp[len(temp)-1]
	*m = temp[0 : len(*m)-1]
	return res
}

type MinHeap []int

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *MinHeap) Pop() interface{} {
	temp := *m
	res := temp[len(temp)-1]
	*m = temp[0 : len(*m)-1]
	return res
}
