package dag

import (
	"reflect"
	"testing"
)

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
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			if got := tc.g.Cycles(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Graph.Cycles() = %v, want %v", got, tc.want)
			}
		})
	}
}
