package dag

import (
	"reflect"
	"testing"
)

type testVertex struct {
	i int
}

func (t *testVertex) Hashcode() int {
	return t.i
}

func TestGraph_Cycles(t *testing.T) {
	tests := []struct {
		name string
		g    *Graph
		want [][]Vertex
	}{
		{
			name: "Empty graph should not contain cycles",
			g:    NewGraph(),
		},
		{
			name: "Graph with linear edges should not contain cycles",
			g:    NewGraph().AddEdge(&testVertex{1}, &testVertex{2}).AddEdge(&testVertex{2}, &testVertex{3}),
		},
		{
			name: "Graph with cycle should be detected",
			g:    NewGraph().AddEdge(&testVertex{1}, &testVertex{2}).AddEdge(&testVertex{2}, &testVertex{3}).AddEdge(&testVertex{3}, &testVertex{1}),
			want: [][]Vertex{[]Vertex{&testVertex{3}, &testVertex{2}, &testVertex{1}}},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			if got := tc.g.Cycles(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Graph.Cycles() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGraph_TopologicalSorting(t *testing.T) {
	tests := []struct {
		name string
		g    *Graph
		want []Vertex
	}{
		{
			name: "Empty graph sorted should return empty vertex",
			g:    NewGraph(),
		},
		{
			name: "Linear graph should return linear vertex slice",
			g:    NewGraph().AddEdge(&testVertex{1}, &testVertex{2}).AddEdge(&testVertex{2}, &testVertex{3}),
			want: []Vertex{&testVertex{1}, &testVertex{2}, &testVertex{3}},
		},
		{
			name: "Multibranched graph should return correct topological sorting",
			g:    NewGraph().AddEdge(&testVertex{1}, &testVertex{2}).AddEdge(&testVertex{2}, &testVertex{3}).AddEdge(&testVertex{4}, &testVertex{5}).AddEdge(&testVertex{1}, &testVertex{5}).AddEdge(&testVertex{2}, &testVertex{5}),
			want: []Vertex{&testVertex{1}, &testVertex{2}, &testVertex{3}, &testVertex{4}, &testVertex{5}},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			if got := tc.g.TopologicalSorting(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Graph.TopologicalSorting() = %v, want %v", got, tc.want)
			}
		})
	}
}
