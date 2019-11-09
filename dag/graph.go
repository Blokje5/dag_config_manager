package dag

type Vertice interface {

	Hashable
}

// Traversal provides a basic interface for graph traversal
type Traversal interface {
	// Visited should return true if vertice "v" has already been visited in this
	// traversal. If the same traversal is used multiple times, the state that
	// tracks visited nodes should be reset.
	Visited(v Vertice) bool

	// AdjacencyList should return the adjacency list representation of the graph.
	// This is represented as a map of the hash code of the vertice to the list of
	// neigbouring vertices.
	AdjacencyList() map[int][]Vertice

	// Vertices should return the set of all Vertices in the graph.
	Vertices() Set
}

// StopFunc should return true when the traversal should end at vertice "v".
type StopFunc func(v Vertice) bool

// DFS performs Depth First Search based on the Traversal, starting at vertice "v".
func DFS(t Traversal, v Vertice, f StopFunc) {
	for _, v := range t.Vertices().List() {
		dfs(t, v.(Vertice), f)
	}
}

func dfs(t Traversal, v Vertice, f StopFunc) {
	a := t.AdjacencyList()
	s := NewStack(v)
	for s.Size() > 0 {
		next, _ := s.Pop()
		u := next.(Vertice)
		if t.Visited(u) {
			continue
		}
		if f(u) {
			return
		}

		s.Push(a[u.Hashcode()])
	}
}