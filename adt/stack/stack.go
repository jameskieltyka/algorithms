package stack

import "sync"

//Stack implements LIFO queue
type Stack struct {
	mut  *sync.Mutex
	head *Data
	size int
}

//Data provides a data node for the LL
type Data struct {
	Value interface{}
	Next  *Data
}

//New creates a new stack with an initial capacity
func New() *Stack {
	return &Stack{
		&sync.Mutex{},
		nil,
		0,
	}
}

//Push adds a new item to the stack
func (s *Stack) Push(data interface{}) {
	s.mut.Lock()
	defer s.mut.Unlock()
	node := &Data{
		Value: data,
		Next:  s.head,
	}
	s.head = node
	s.size++
}

// Pop returns the last item added to the stack,
func (s *Stack) Pop() interface{} {
	s.mut.Lock()
	defer s.mut.Unlock()
	if s.head == nil {
		return nil
	}

	item := s.head
	s.head = s.head.Next
	s.size--
	return item.Value
}

//Len returns the size of the stack
func (s *Stack) Len() int {
	return s.size
}
