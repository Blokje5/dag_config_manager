package dag

// Graph is a data structure that represents a Graph
type Graph struct {
	edges Set
	vertices Set
	adjacencyList map[int][]int
}

type Edge struct {}

type Vertice struct {}

// NewGraph creates a new Graph
func NewGraph() *Graph {
	return &Graph{
		edges: NewSet(),
		vertices: NewSet(),
		adjacencyList: make(map[int][]int),
	}
}

// AddVertices adds one or more vertices to the graph
func (g *Graph) AddVertices(v []*Vertice) {
	g.vertices.Add(v)
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(source Hashable, target Hashable) {
	g.adjacencyList[source.Hashcode()] = append(g.adjacencyList[source.Hashcode()], target.Hashcode())
}

// Vertices returns a list of vertices
func (g *Graph) Vertices() []Vertice {
	l := g.vertices.List()
	v := make([]Vertice, g.vertices.Len())
	for _, i := range l {
		v = append(v, i.(Vertice))
	}

	return v
}

// Edges returns a list of edges
func (g *Graph) Edges() []Edge {
	l := g.edges.List()
	e := make([]Edge, g.edges.Len())
	for _, i := range l {
		e = append(e, i.(Edge))
	}

	return e
}

// DFS
func (g *Graph) DFS() {
	visited := make([]bool, len(g.adjacencyList))
	for v := range g.adjacencyList {
		if !visited[v] {
			g.dfcRecursive(v, visited)
		}
	}
	
}

func (g *Graph) dfcRecursive(v int, visited []bool) {
	visited[v] = true

	neighbours := g.adjacencyList[v]
	for n := range neighbours {
		if !visited[n] {
			g.dfcRecursive(n, visited)
		}
	}
}

func (g *Graph) IsCyclic() bool {
	visited := make([]bool, len(g.adjacencyList))
	stack := make([]bool, len(g.adjacencyList))
	for v := range g.adjacencyList {
		if g.isCyclicRecursive(v, visited, stack) {
			return true
		}
	}

	return false
}

func (g *Graph) isCyclicRecursive(v int, visited []bool, stack []bool) bool {
	if stack[v] {
		return true
	}

	if visited[v] {
		return false
	}

	stack[v] = true
	visited[v] = true

	neighbours := g.adjacencyList[v]
	for n := range neighbours {
		if g.isCyclicRecursive(n, visited, stack) {
			return true
		}
	}

	stack[v] = false; 
  
	return false; 
}