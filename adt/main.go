package main

import (
	"fmt"

	priority "github.com/jkieltyka/algorithms/adt/priorityQueue"
	"github.com/jkieltyka/algorithms/adt/queue"
	"github.com/jkieltyka/algorithms/adt/stack"
)

func main() {
	q := queue.New()
	q.Add(1)
	q.Add(2)
	q.Add(3)

	s := stack.New()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	p := priority.New()
	p.Insert(1, 7)
	p.Insert(5, 10)
	p.Insert(3, 9)
	p.Insert(2, 8)

	for q.Len() != 0 {
		fmt.Println(q.Pop().(int))
	}

	for s.Len() != 0 {
		fmt.Println(s.Pop().(int))
	}

	fmt.Println()
	for p.Len() != 0 {
		fmt.Println(p.Remove())
	}
}
