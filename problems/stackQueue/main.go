package main

import "fmt"

//Implement a queue using two stacks. Recall that a queue is a FIFO (first-in, first-out) data structure with the following methods: enqueue, which inserts an element into the queue, and dequeue, which removes it.

func main() {
	queue := StackQueue{
		0,
		&stack{},
		&stack{},
	}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	for queue.size != 0 {
		fmt.Println(queue.Dequeue())
	}
	
}

type stack []int

func (s *stack) Push(v int) {
    *s = append(*s, v)
}

func (s *stack) Pop() int {
    res:=(*s)[len(*s)-1]
    *s=(*s)[:len(*s)-1]
    return res
}

type StackQueue struct {
	size int
	instack *stack
	outstack *stack
}

func (s *StackQueue) Enqueue(i int) {
	s.size++
	s.instack.Push(i)
}

func (s *StackQueue) Dequeue() int {
	s.size--
	if len(*s.outstack) != 0 {
		return s.outstack.Pop()
	}

	for len(*s.instack) != 0 {
		s.outstack.Push(s.instack.Pop())
	}

	return s.outstack.Pop()
}