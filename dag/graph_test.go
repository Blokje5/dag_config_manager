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
		name   string
		g *Graph
		want   [][]Vertex
	}{
		{
			name: "Empty graph should not contain cycles",
			g: NewGraph(),
		},
		{
			name: "Graph with linear edges should not contain cycles",
			g: NewGraph().AddEdge(&testVertex{1}, &testVertex{2}).AddEdge(&testVertex{2}, &testVertex{3}),
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
