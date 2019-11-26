package dag

// StronglyConnectedComponents Returns the strongly connected components of graph "G".
// Uses the Tarjan algorihtm, see:
// https://en.wikipedia.org/wiki/Tarjan%27s_strongly_connected_components_algorithm
func StronglyConnectedComponents(g *Graph) [][]Vertex {
	vs := g.Vertices()
	t := &tarjanTraversal{
		nextIndex: 0,
		vertexIndex: make(map[int]int, len(vs)),
		stack: NewStack(),
		graph: g,
	}

	for _, v := range  vs {
		if !t.visited(v) {
			stronglyConnectedComponents(t, v)
		}
	}

	return t.stronglyConnectedComponents
}

type tarjanTraversal struct {
	nextIndex int
	vertexIndex map[int]int
	stack Stack
	graph *Graph
	stronglyConnectedComponents [][]Vertex
}

func stronglyConnectedComponents(t *tarjanTraversal, v Vertex) int {
	idx := t.visit(v)
	minIdx := idx

	for _, u := range t.graph.Neighbours(v) {
		if !t.visited(u) {
			// Vertex "u" has not yet been visited, recursivly search "u"
			minIdx = min(minIdx, stronglyConnectedComponents(t, u))
		} else if t.stack.Contains(u) {
			// Vertex u is is in the Stack, hence part of the current SCC
			minIdx = min(minIdx, t.vertexIndex[u.Hashcode()])
		}
	}

	if minIdx == idx {
		// Start a new scc from the root of this scc
		var scc []Vertex
		for {
			el, _ := t.stack.Pop()
			v2 := el.(Vertex)
			scc = append(scc, v2)
			if v2 == v {
				break
			}
		}
		t.stronglyConnectedComponents = append(t.stronglyConnectedComponents, scc)
	}

	return minIdx
}

func (t *tarjanTraversal) visited(v Vertex) bool {
	_, ok := t.vertexIndex[v.Hashcode()]

	return ok
}


func (t *tarjanTraversal) visit(v Vertex) int {
	idx := t.nextIndex
	t.stack.Push(v)
	t.vertexIndex[v.Hashcode()] = idx
	t.nextIndex++
	return idx
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}