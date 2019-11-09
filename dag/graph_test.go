package dag

import (
	"testing"
)

func TestGraph_IsCyclic(t *testing.T) {
	tests := []struct {
		name string
		adjacencyList map[int][]int
		expected bool
	}{
		{
			"Simple linear graph does not contain cycles",
			map[int][]int {
				1: []int{2},
				2: []int{3},
			},
			false,
		},
		{
			"Simple cycle",
			map[int][]int {
				1: []int{2},
				2: []int{3},
				3: []int{1},
			},
			false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := Graph{
				adjacencyList: tc.adjacencyList,
			}
			isCyclic := g.IsCyclic()
			if isCyclic != tc.expected {
				t.Errorf("Expected cycle: %v but instead was %v", tc.expected, isCyclic)
			}
		})
	}
}
