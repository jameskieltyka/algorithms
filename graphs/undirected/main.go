package main

import "fmt"

type Graph struct {
	V   int
	E   int
	Adj []map[int]bool
}

func main() {
	g := New(6)
	g.AddEdge(0, 2)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(0, 5)
	g.AddEdge(3, 5)
	g.AddEdge(3, 2)
	g.AddEdge(4, 2)
	g.AddEdge(3, 4)

	d := &DFS{
		marked: make([]bool, 6),
		count:  0,
	}
	d.dfs(g, 0)
	fmt.Println(d.count)

	b := &BFS{
		marked: make([]bool, 6),
		count:  0,
		queue:  make(chan int, 6),
	}
	b.bfs(g, 0)
	fmt.Println(b.count)

}

func New(vert int) *Graph {
	adjList := make([]map[int]bool, vert)
	for i := range adjList {
		adjList[i] = make(map[int]bool)
	}
	return &Graph{
		V:   vert,
		E:   0,
		Adj: adjList,
	}
}

func (g *Graph) RemoveEdge(v, w int) {
	delete(g.Adj[v], w)
	delete(g.Adj[w], v)
	g.E--
}

func (g *Graph) AddEdge(v, w int) {
	g.Adj[v][w] = true
	g.Adj[w][v] = true
	g.E++
}

func (g *Graph) AddVertex() int {
	g.Adj = append(g.Adj, make(map[int]bool))
	g.V++
	return g.V
}

type DFS struct {
	marked []bool
	count  int
}

func (d *DFS) dfs(g *Graph, v int) {
	d.count++
	d.marked[v] = true
	for i := range g.Adj[v] {
		if !d.marked[i] {
			d.dfs(g, i)
		}
	}
}

type BFS struct {
	marked []bool
	queue  chan int
	count  int
}

func (b *BFS) bfs(g *Graph, v int) {
	b.queue <- v
	b.marked[v] = true
	b.count++
	for {
		select {
		case v := <-b.queue:
			for i := range g.Adj[v] {
				if !b.marked[i] {
					b.marked[i] = true
					b.queue <- i
					b.count++
				}
			}
		default:
			close(b.queue)
			return
		}
	}
}
