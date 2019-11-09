package dag

type Vertice interface {
	Neighbours() []Vertice
	Hashable
}

// DFS performs Depth First Search on the adjecency list
func DFS(adjacencyList []Vertice) {
	visited := make(map[int]bool, len(adjacencyList))
	for _, v := range adjacencyList {
		dfs(v, visited)
	}
}

func dfs(v Vertice, visited map[int]bool) {
	s := NewStack(v)
	for s.Size() > 0 {
		next, _ := s.Pop()
		u := next.(Vertice)
		if !visited[u.Hashcode()] {
			visited[u.Hashcode()] = true
		}

		s.Push(u.Neighbours)
	}
}