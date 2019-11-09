package dag

// T represents interface{}. It simplifies specifying interface{} methods.
type T interface{}

type Stack struct {
	top int
	items []T
}

// NewStack returns a Stack initialized with the items.
func NewStack(ts ...T) (*Stack){
	return &Stack{
		top: len(ts),
		items: ts,
	}
}

// Push pushes the items onto the Stack.
func (s *Stack) Push(ts ...T) {
	s.items = append(s.items, ts...)
	s.top += len(ts)
}

// Pop removes the top item from the Stack 
func (s *Stack) Pop() (T, bool) {
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
func (s *Stack) Peek() (T, bool) {
	if s.top == 0 {
		return nil, false
	}

	t := s.items[s.top-1]
	return t, true
}

// Size returns the size of the stack
func (s *Stack) Size() int {
	return len(s.items)
}