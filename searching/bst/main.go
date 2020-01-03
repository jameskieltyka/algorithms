package main

import "fmt"

type BST struct {
	root      *Node
	NodeCount int
}

type Node struct {
	Key       string
	Value     int
	NodeCount int
	Left      *Node
	Right     *Node
}

func main() {
	b := &BST{}
	b.put("5", 10)
	b.put("4", 9)
	b.put("7", 11)
	b.put("8", 13)
	//b.put("5", 13)

	b.inorderTrav()

	fmt.Println(b.root)
	fmt.Println(b.search("5"))
	fmt.Println(b.search("4"))

	b.delete("5")
	fmt.Println(b.search("5"))
	fmt.Println(b.search("6"))
	fmt.Println(b.search("8"))
	fmt.Println(b.search("4"))

}

func (b *BST) search(key string) *Node {
	return b.root.search(key)
}

func (n *Node) search(key string) *Node {
	if n == nil {
		return nil
	}
	if n.Key == key {
		return n
	} else if key < n.Key {
		return n.Left.search(key)
	}

	return n.Right.search(key)
}

func (b *BST) put(key string, value int) {
	if b.root == nil {
		b.root = &Node{
			Key:   key,
			Value: value,
		}
		return
	}
	b.root.put(key, value)
}

func (n *Node) put(key string, value int) {
	if n == nil {
		n = &Node{
			Key:   key,
			Value: value,
		}
		return
	}

	if n.Key == key {
		n.Value = value
		return
	}

	if n.Key < key {
		if n.Right == nil {
			n.Right = &Node{
				Key:   key,
				Value: value,
			}
		}
		n.Right.put(key, value)
		return
	}

	if n.Key > key {
		if n.Left == nil {
			n.Left = &Node{
				Key:   key,
				Value: value,
			}
		}
		n.Left.put(key, value)
		return
	}
}

func (b *BST) delete(key string) {
	b.root = b.root.delete(key)
}

func (n *Node) delete(key string) *Node {
	if n == nil {
		return nil
	}

	if n.Key == key {
		if n.Right == nil {
			return n.Left
		}
		if n.Left == nil {
			return n.Right
		}
		nTemp := n.Right.successor()
		nTemp.Right = n.Right.successorChild()
		nTemp.Left = n.Left

		return nTemp
	}

	if n.Key < key {
		n.Right = n.Right.delete(key)

	}

	if n.Key > key {
		n.Left = n.Left.delete(key)
	}
	return nil
}

func (n Node) successorChild() *Node {

	if n.Left == nil {
		return n.Right
	}

	n.Left = n.Left.successorChild()
	return &n
}

func (n Node) successor() *Node {

	if n.Left == nil {
		return &n
	}

	return n.Left.successor()
}

func (b *BST) inorderTrav() {
	if b.root == nil {
		return
	}
	b.root.inorder()

}

func (n *Node) inorder() {
	if n.Left != nil {
		n.Left.inorder()
	}
	fmt.Printf(n.Key + "->")
	if n.Right != nil {
		n.Right.inorder()
	}
}

func (b *BST) preorderTrav() {
	if b.root == nil {
		return
	}
	b.root.preorder()

}

func (n *Node) preorder() {
	fmt.Printf(n.Key + "->")
	if n.Left != nil {
		n.Left.inorder()
	}
	if n.Right != nil {
		n.Right.inorder()
	}
}

func (b *BST) postorderTrav() {
	if b.root == nil {
		return
	}
	b.root.preorder()

}

func (n *Node) postorder() {

	if n.Left != nil {
		n.Left.inorder()
	}
	if n.Right != nil {
		n.Right.inorder()
	}
	fmt.Printf(n.Key + "->")
}
