package state

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	var nilSlice []Operation
	tests := []struct {
		name           string
		reconciliation Reconciliation
		want           []Operation
	}{
		{
			"No diff leads to no operations",
			Reconciliation{
				Before: make(map[string]interface{}),
				After: make(map[string]interface{}),
			},
			nilSlice,
		},
		{
			"No diff if same keys and values leads to no operations",
			Reconciliation{
				Before: map[string]interface{}{
					"a": 1,
				},
				After: map[string]interface{}{
					"a": 1,
				},
			},
			nilSlice,
		},
		{
			"Deleted keys lead to Delete operations",
			Reconciliation{
				Before: map[string]interface{}{
					"a": 1,
				},
				After: map[string]interface{}{},
			},
			[]Operation{
				Delete{
					Key: "a",
					Value: 1,
				},
			},
		},
		{
			"Created keys lead to Create operations",
			Reconciliation{
				Before: map[string]interface{}{},
				After: map[string]interface{}{
					"a": 1,
				},
			},
			[]Operation{
				Create{
					Key: "a",
					Value: 1,
				},
			},
		},
		{
			"Updates lead to update operations",
			Reconciliation{
				Before: map[string]interface{}{
					"a": 1,
				},
				After: map[string]interface{}{
					"a": 2,
				},
			},
			[]Operation{
				Update{
					Key: "a",
					Before: 1,
					After: 2,
				},
			},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {
			ops := tc.reconciliation.Diff()
			if !reflect.DeepEqual(tc.want, ops) {
				t.Errorf("Expected operations: %v, actual operations: %v", tc.want, ops)
			}
		})
	}
}
