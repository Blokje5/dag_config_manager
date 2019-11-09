package dag

import (
	"reflect"
	"testing"
)

func TestStack_Push(t *testing.T) {
	tests := []struct {
		name   string
		stack *Stack
		elements []T
		expected *Stack
	}{
		{
			"Adding an element to an empty stack",
			NewStack(),
			[]T{1},
			NewStack(1),
		},
		{
			"Adding multiple elements to an empty stack",
			NewStack(),
			[]T{1, 2, 3},
			NewStack(1, 2, 3),
		},
		{
			"Adding multiple elements to an prefilled stack",
			NewStack(1),
			[]T{2, 3, 4},
			NewStack(1, 2, 3, 4),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.stack
			s.Push(tc.elements...)
			if !reflect.DeepEqual(s,tc.expected) {
				t.Errorf("Stack.Push() got = %v, expected %v", s.items, tc.expected.items)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type result struct {
		item T
		found bool
	} 
	tests := []struct {
		name   string
		stack *Stack
		expectedStack *Stack
		expectedResult result
	}{
		{
			"Popping an element of an empty stack",
			NewStack(),
			NewStack(),
			result{nil, false},
		},
		{
			"Popping an element of a filled stack",
			NewStack(1),
			NewStack(),
			result{1, true},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.stack
			item, found := s.Pop()
			// TODO (leijsackers): Currently I have no better way to deal with nil vs empty slice
			if s.top != tc.expectedStack.top || s.Size() != tc.expectedStack.Size() {
				t.Errorf("Stack.Pop() got stack = %v, expected stack %v", s, tc.expectedStack)
			}

			if item != tc.expectedResult.item || found != tc.expectedResult.found {
				t.Errorf("Stack.Pop() got return %v, %v,  expected return %v, %v", t, found, tc.expectedResult.item, tc.expectedResult.found)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	type result struct {
		item T
		found bool
	} 
	tests := []struct {
		name   string
		stack *Stack
		expectedStack *Stack
		expectedResult result
	}{
		{
			"Peek an element in an empty stack",
			NewStack(),
			NewStack(),
			result{nil, false},
		},
		{
			"Peek an element in a filled stack",
			NewStack(1),
			NewStack(1),
			result{1, true},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.stack
			item, found := s.Peek()
			if !reflect.DeepEqual(s,tc.expectedStack) {
				t.Errorf("Stack.Peek() got stack = %v, expected stack %v", s, tc.expectedStack)
			}

			if item != tc.expectedResult.item || found != tc.expectedResult.found {
				t.Errorf("Stack.Peek() got return %v, %v,  expected return %v, %v", t, found, tc.expectedResult.item, tc.expectedResult.found)
			}
		})
	}
}

func TestStack_Size(t *testing.T) {
	tests := []struct {
		name   string
		stack *Stack
		expected   int
	}{
		{
			"Empty stack size should be zero",
			NewStack(),
			0,
		},
		{
			"Initialized stack size should be equal to length of elements",
			NewStack(1, 2, 3),
			3,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.stack
			if got := s.Size(); got != tc.expected {
				t.Errorf("Stack.Size() = %v, want %v", got, tc.expected)
			}
		})
	}
}
