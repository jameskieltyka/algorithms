package main

type Graph struct {
	V   int
	E   int
	Adj []map[int]bool
}

func main() {

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
