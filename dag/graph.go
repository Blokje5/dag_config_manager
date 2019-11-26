package dag

type Vertex interface {
	Hashable
}

type Graph struct {
	vertices Set
	adjacencyList map[int][]Vertex
}

// NewGraph returns an initialized Graph 
func NewGraph() *Graph {
	return &Graph {
		vertices: NewSet(),
		adjacencyList: make(map[int][]Vertex),
	}
}

// AddEdge adds an edge from vertex "v1" to vertex "v2".
// It will add the vertices to the graph if both where not present.
// The graph is returned to make it easy to chain methods
func (g *Graph) AddEdge(v1, v2 Vertex) *Graph {
	g.vertices.Add(v1, v2)

	neighbours := g.Neighbours(v1)
	for _, v := range neighbours {
		if v2.Hashcode() == v.Hashcode() {
			return g
		}
	}

	g.adjacencyList[v1.Hashcode()] = append(g.adjacencyList[v1.Hashcode()], v2)
	return g
}

// Vertices returns a list of vertices
func (g *Graph) Vertices() []Vertex {
	l := g.vertices.List()
	v := make([]Vertex, 0, g.vertices.Len())
	for _, i := range l {
		v = append(v, i.(Vertex))
	}

	return v
}

// AdjacencyList returns the adjacency list representation of the graph
func (g *Graph) AdjacencyList() map[int][]Vertex {
	return g.adjacencyList
}

// Neighbours returns all neighbouring vertices of vertex "v".
// If "v" has no neighbours, a nil slice will be returned.
func (g *Graph) Neighbours(v Vertex) []Vertex {
	return g.adjacencyList[v.Hashcode()]
}

// Cycles returns the list of components in graph "g" that
// contain a cycle. Each component is represented as a Vertex slice. 
// TODO implement with Set instead of slice.
func (g *Graph) Cycles() [][]Vertex {
	sccs := StronglyConnectedComponents(g)
	var cycles [][]Vertex
	for _, scc := range sccs {
		// Cycle when there is a scc larger then size 1
		if len(scc) > 1 {
			cycles = append(cycles, scc)
		}
	}

	return cycles
}

// TopologicalSorting returns the topological sorted dag, with
// each vertex ordered from starting point to end point.
// The method assumes a valid dag as input.
func (g *Graph) TopologicalSorting() []Vertex {
	vs := g.Vertices()
	vl := len(vs)

	t := &topologicalSortingTraversal {
		visitedMap: make(map[int]struct{}, vl),
		sorted: make([]Vertex, vl),
		graph: g,
		time: -1,
	}

	for _, v := range  vs {
		if !t.visited(v) {
			topologicalSorting(t, v)
		}
	}

	var reversed []Vertex
	for i := vl - 1; i >= 0; i-- {
		reversed = append(reversed, t.sorted[i])
	}

	return reversed
}

type topologicalSortingTraversal struct {
	sorted []Vertex
	visitedMap map[int]struct{}
	graph *Graph
	time int
}

func (t *topologicalSortingTraversal) visited(v Vertex) bool {
	_, ok := t.visitedMap[v.Hashcode()]

	return ok
}

func (t *topologicalSortingTraversal) visit(v Vertex) {
	t.visitedMap[v.Hashcode()] = struct{}{}
}

func topologicalSorting(t *topologicalSortingTraversal, v Vertex) {
	t.visit(v)
	for _, u := range t.graph.Neighbours(v) {
		if !t.visited(u) {
			topologicalSorting(t, u)
		}
	}
	
	t.time++
	t.sorted[t.time] = v
}

// Traversal provides a basic interface for graph traversal
type Traversal interface {
	// Visited should return true if Vertex "v" has already been visited in this
	// traversal. If the same traversal is used multiple times, the state that
	// tracks visited nodes should be reset.
	Visited(v Vertex) bool

	// AdjacencyList should return the adjacency list representation of the graph.
	// This is represented as a map of the hash code of the Vertex to the list of
	// neigbouring Vertexs.
	AdjacencyList() map[int][]Vertex

	// Vertices should return the set of all Vertices in the graph.
	Vertices() Set
}

// StopFunc should return true when the traversal should end at Vertex "v".
type StopFunc func(v Vertex) bool

// DFS performs Depth First Search based on the Traversal, starting at Vertex "v".
func DFS(t Traversal, v Vertex, f StopFunc) {
	for _, v := range t.Vertices().List() {
		dfs(t, v.(Vertex), f)
	}
}

func dfs(t Traversal, v Vertex, f StopFunc) {
	a := t.AdjacencyList()
	s := NewStack(v)
	for s.Len() > 0 {
		next, _ := s.Pop()
		u := next.(Vertex)
		if t.Visited(u) {
			continue
		}
		if f(u) {
			return
		}

		s.Push(a[u.Hashcode()])
	}
}