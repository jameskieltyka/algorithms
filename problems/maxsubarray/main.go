package main

//Given an array of integers and a number k, where 1 <= k <= length of the array,
// compute the maximum values of each subarray of length k.

type dequeue struct {
	size    int
	maxSize int
	head    *node
	tail    *node
}

type node struct {
	val  int
	prev *node
	next *node
}

func NewDeque(size int) *dequeue {
	return &dequeue{
		maxSize: size,
	}
}

func (d *dequeue) AddHead(data int) {
	n := &node{
		val: data,
	}
	d.size++
	if d.head == nil {
		d.head = n
		d.tail = n
	} else {
		n.next = d.head
		d.head.prev = n
		d.head = n
	}
}

func (d *dequeue) AddTail(data int) {
	n := &node{
		val: data,
	}
	d.size++
	if d.tail == nil {
		d.head = n
		d.tail = n
	} else {
		n.prev = d.tail
		temp := d.tail
		temp.next = n
		d.tail = n
	}
}

func (d *dequeue) PopTail() *int {
	if d.tail == nil {
		return nil
	}
	d.size--
	data := d.tail.val
	temp := d.tail
	d.tail = temp.prev
	temp = nil
	return &data

}

func (d *dequeue) PopHead() *int {
	if d.head == nil {
		return nil
	}
	d.size--
	data := d.head.val
	temp := d.head
	d.head = temp.next
	temp = nil
	return &data

}

func (d *dequeue) PeekHead() *int {
	if d.head == nil {
		return nil
	}
	return &d.head.val
}

func (d *dequeue) PeekTail() *int {
	if d.tail == nil {
		return nil
	}
	return &d.tail.val
}

func maxSubarray(arr []int, k int) []int {
	res := make([]int, 0, k)
	dequeue := NewDeque(k)

	count := 0
	for i := 0; i < len(arr); i++ {
		if dequeue.size == 0 || arr[i] > arr[*dequeue.PeekHead()] {
			for dequeue.PeekHead() != nil {
				dequeue.PopHead()
			}
			dequeue.AddHead(i)
		} else {
			for arr[*dequeue.PeekTail()] < arr[i] {
				dequeue.PopTail()
			}
			dequeue.AddTail(i)
		}
		count++
		if count%k == 0 {
			for *dequeue.PeekHead() < i+1-k {
				dequeue.PopHead()
			}
			count = k - 1
			res = append(res, arr[*dequeue.PeekHead()])
		}
	}

	return res
}
