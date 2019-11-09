package dag

import (
	"reflect"
	"testing"
)

func TestSet_Add(t *testing.T) {
	tests := []struct {
		name   string
		values []interface{}
		expected_size int
	}{
		{
			"Can add single item",
			[]interface{}{1},
			1,
		},
		{
			"Can add multiple items",
			[]interface{}{1, 2, 3},
			3,
		},
		{
			"Is a set",
			[]interface{}{1, 2, 2},
			2,
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {
			s := NewSet()
			s.Add(tc.values...)
			if s.Len() != tc.expected_size {
				t.Errorf("Expected %d items in Set, actual items in Set: %d", tc.expected_size, s.Len())
			}
		})
	}
}

func TestSet_Remove(t *testing.T) {
	t.Run("Can remove items", func(t *testing.T) {
		s := NewSet()
		s.Add(1, 2, 3, 4)
		s.Remove(1)
		if s.Len() != 3 {
			t.Errorf("Expected %d items in Set after removal, actual items in Set: %d", 3, s.Len())
		}

		s.Remove(1)
		if s.Len() != 3 {
			t.Errorf("Expected %d items in Set after removal, actual items in Set: %d", 3, s.Len())
		}

		s.Remove(2, 3)
		if s.Len() != 1 {
			t.Errorf("Expected %d items in Set after removal, actual items in Set: %d", 1, s.Len())
		}
	})
}

func TestSet_Contains(t *testing.T) {
	t.Run("Can remove items", func(t *testing.T) {
		s := NewSet()
		s.Add(1, 2, 3, 4)
		if ok := s.Contains(1); !ok {
			t.Errorf("Expected item to be contained in set")
		}
		
		if ok := s.Contains(1, 3); !ok {
			t.Errorf("Expected items to be contained in set")
		}

		if ok := s.Contains(5); ok {
			t.Errorf("Expected items not to be contained in set")
		}
	})
}

func TestSet_List(t *testing.T) {
	tests := []struct {
		name   string
		values []interface{}
	}{
		{
			"List returns single Add call",
			[]interface{}{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSet()
			s.Add(tc.values...)
			l := s.List()
			if reflect.DeepEqual(l, tc.values) {
				t.Error("Expected list to return values")
			}
		})
	}
}