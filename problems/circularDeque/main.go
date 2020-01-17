package main

func main() {

}

//leetcode.com/problems/design-circular-deque

type MyCircularDeque struct {
	len   int
	max   int
	start int
	end   int
	data  []int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		data:  make([]int, k),
		max:   k,
		start: k - 1,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.len == this.max {
		return false
	}

	this.data[this.start] = value
	pos := (this.start + this.max - 1) % this.max
	this.start = pos
	this.len++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.len == this.max {
		return false
	}

	this.data[this.end] = value
	pos := (this.end + 1) % this.max
	this.end = pos
	this.len++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.len == 0 {
		return false
	}
	pos := (this.start + 1) % this.max
	this.start = pos
	this.len--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.len == 0 {
		return false
	}
	pos := (this.end + this.max - 1) % this.max
	this.end = pos
	this.len--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.len == 0 {
		return -1
	}
	return this.data[(this.start+1)%this.max]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.len == 0 {
		return -1
	}
	return this.data[(this.end-1+this.max)%this.max]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.len == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.len == this.max
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
