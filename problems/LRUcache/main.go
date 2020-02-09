package main

type LRUCache struct {
	capacity int
	size     int
	head     *node
	tail     *node
	access   map[int]*node
}

type node struct {
	key  int
	val  int
	prev *node
	next *node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		access:   make(map[int]*node, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.access[key]; ok {
		if n == this.head {
			return n.val
		}
		this.remove(n)
		this.add(n)
		return n.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	var n *node

	ns, ok := this.access[key]
	if ok {
		ns.val = value
		this.remove(ns)
		this.add(ns)
		return
	}

	n = &node{
		key: key,
		val: value,
	}
	this.access[key] = n
	this.add(n)

	if len(this.access) > this.capacity {
		delete(this.access, this.tail.key)
		this.remove(this.tail)
	}
}

func (this *LRUCache) add(node *node) {
	node.prev = nil
	node.next = this.head
	if this.head != nil {
		this.head.prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
	}
}

func (this *LRUCache) remove(node *node) {
	if node != this.head {
		node.prev.next = node.next
	} else {
		this.head = node.next
	}
	if node != this.tail {
		node.next.prev = node.prev
	} else {
		this.tail = node.prev
	}
}
