package dag

import (
	"fmt"
	"strings"
)

// T represents interface{}. It simplifies specifying interface{} methods.
type T interface{}

type Stack interface {
	Push(ts ...T)
	Pop() (T, bool)
	Peek() (T, bool)
	Len() int
	Contains(t T) bool
	fmt.Stringer
}

type stackImpl struct {
	top   int
	items []T
}

// NewStack returns a Stack initialized with the items.
func NewStack(ts ...T) Stack {
	return &stackImpl{
		top:   len(ts),
		items: ts,
	}
}

// Push pushes the items onto the Stack.
func (s *stackImpl) Push(ts ...T) {
	s.items = append(s.items, ts...)
	s.top += len(ts)
}

// Pop removes the top item from the Stack
func (s *stackImpl) Pop() (T, bool) {
	if s.top == 0 {
		return nil, false
	}

	t := s.items[s.top-1]
	s.top--
	s.items = s.items[0:s.top]

	return t, true
}

// Peek returns the top item if found without removing it from the stack.
// If it did not find an item, it will return nil, false
func (s *stackImpl) Peek() (T, bool) {
	if s.top == 0 {
		return nil, false
	}

	t := s.items[s.top-1]
	return t, true
}

// Size returns the size of the stack
func (s *stackImpl) Len() int {
	return len(s.items)
}

// Contains returns true if the Stack contains element "t".
func (s *stackImpl) Contains(t T) bool {
	for _, n := range s.items {
		if t == n {
			return true
		}
	}

	return false
}

func (s *stackImpl) String() string {
	strs := make([]string, s.Len())
	for i, item := range s.items {
		strs[i] = fmt.Sprintf("%v", item)
	}

	return strings.Join(strs, ",")
}
