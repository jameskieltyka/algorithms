package queue

import "sync"

//Queue implements a linkedlist FIFO queue
type Queue struct {
	mut    *sync.Mutex
	head   *Data
	tail   *Data
	length int
}

//Data provides a data node for the LL
type Data struct {
	Value interface{}
	Next  *Data
}

//New creates a new queue
func New() *Queue {
	return &Queue{
		&sync.Mutex{},
		nil,
		nil,
		0,
	}
}

//Add inserts new data O(1)
func (q *Queue) Add(value interface{}) {
	q.mut.Lock()
	defer q.mut.Unlock()

	node := &Data{
		Value: value,
		Next:  nil,
	}

	if q.head == nil {
		q.head = node
		q.tail = node
	}

	q.tail.Next = node
	q.tail = node
	q.length++

}

//Pop retreives the first element in the queue O(1)
func (q *Queue) Pop() interface{} {
	q.mut.Lock()
	defer q.mut.Unlock()
	if q.head == nil {
		return nil
	}

	data := q.head.Value
	q.length--
	q.head = q.head.Next

	if q.head == nil {
		q.tail = nil
	}
	return data
}

//Len returns the length of the queue
func (q *Queue) Len() int {
	return q.length
}
