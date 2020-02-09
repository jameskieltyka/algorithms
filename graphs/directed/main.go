package main

import "fmt"

func main() {
	g := NewGraph(5)
	g.addEdge(0, 1)
	g.addEdge(1, 3)
	g.addEdge(2, 3)
	g.addEdge(3, 4)

	g.topoorder()
}

type Digraph struct {
	V   int
	E   int
	Adj [][]int
}

func NewGraph(v int) *Digraph {
	slice := make([][]int, v)
	for i := 0; i < v; i++ {
		slice[i] = make([]int, 0)
	}

	return &Digraph{
		V:   v,
		Adj: slice,
	}
}

func (d *Digraph) addEdge(v, w int) {
	d.Adj[v] = append(d.Adj[v], w)
	d.E++
}

func (d *Digraph) reverse() *Digraph {
	r := NewGraph(d.V)
	for i := 0; i < d.V; i++ {
		for _, w := range d.Adj[i] {
			r.addEdge(w, i)
		}
	}
	return r
}

func (d *Digraph) topoorder() {
	marked := make([]bool, d.V)
	onStack := make([]bool, d.V)
	topoorder := make([]int, 0, d.V)
	for i := 0; i < d.V; i++ {
		if !marked[i] {
			d.dfs(marked, onStack, &topoorder, i)
		}
	}

	for i := range topoorder {
		fmt.Printf("%d->", topoorder[len(topoorder)-1-i])
	}
}

func (d *Digraph) dfs(marked []bool, onStack []bool, topoorder *[]int, s int) {
	onStack[s] = true
	marked[s] = true
	for _, v := range d.Adj[s] {
		if onStack[v] {
			fmt.Println("cycle found")
			return
		}
		if !marked[v] {
			d.dfs(marked, onStack, topoorder, v)
		}
	}
	*topoorder = append(*topoorder, s)
	onStack[s] = false
}
