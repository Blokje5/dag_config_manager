package dag

// Graph is a data structure that represents a Graph
type Graph struct {
	edges Set
	vertices Set
}

type Edge struct {}

type Vertice struct {}

// NewGraph creates a new Graph
func NewGraph() *Graph {
	return &Graph{
		edges: NewSet(),
		vertices: NewSet(),
	}
}

// AddVertices adds one or more vertices to the graph
func (g *Graph) AddVertices(v []*Vertice) {
	g.vertices.Add(v)
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