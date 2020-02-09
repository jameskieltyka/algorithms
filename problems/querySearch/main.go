package main

import "fmt"

//Implement an autocomplete system. That is, given a query string s and a set of all possible query strings, return all strings in the set that have s as a prefix.
//For example, given the query string de and the set of strings [dog, deer, deal], return [deer, deal].

type Trie struct {
	root *Node
}

type Node struct {
	isWord bool
	val    rune
	next   map[rune]*Node
}

func main() {
	trie := NewTrie()
	trie.Add("dog")
	trie.Add("deer")
	trie.Add("deal")

	fmt.Println(trie.Search("de"))

}

func NewTrie() *Trie {
	return &Trie{
		root: &Node{
			next: make(map[rune]*Node, 26),
		},
	}
}

func (t *Trie) Add(s string) {
	node := t.root
	for _, c := range s {
		if n, ok := node.next[c]; ok {
			node = n
		} else {
			n := &Node{
				val:  c,
				next: make(map[rune]*Node, 26),
			}
			node.next[c] = n
			node = n
		}
	}
	node.isWord = true
}

func (t *Trie) Search(sub string) []string {
	words := []string{}
	node := t.root
	for _, c := range sub {
		if n, ok := node.next[c]; ok {
			node = n
		} else {
			return words
		}
	}
	fmt.Println(sub)
	node.SearchWord(sub, &words)
	return words
}

func (n *Node) SearchWord(sub string, res *[]string) {
	if n.isWord {
		*res = append(*res, sub)
	}

	for c, node := range n.next {
		tmp := sub + string(c)
		node.SearchWord(tmp, res)
	}
}
