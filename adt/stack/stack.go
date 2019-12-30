package stack

import "sync"

//Stack implements a linkedlist FIFO queue
type Stack struct {
	mut  *sync.Mutex
	data []interface{}
}

//New creates a new stack with an initial capacity
func New(size int) *Stack {
	return &Stack{
		&sync.Mutex{},
		make([]interface{}, 0, size),
	}
}

//Push adds a new item to the stack
func (s *Stack) Push(data interface{}) {
	s.mut.Lock()
	defer s.mut.Unlock()
	s.data = append(s.data, data)
}

// Pop returns the last item added tot the stack,
// note the slice is not resized if too few items are added to the stack
func (s *Stack) Pop(data interface{}) interface{} {
	s.mut.Lock()
	defer s.mut.Unlock()
	lenStack := len(s.data)
	if lenStack == 0 {
		return nil
	}

	item := s.data[lenStack-1]
	s.data = s.data[:lenStack-1]
	return item
}
