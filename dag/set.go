package dag


// Set is a set data structure
type Set interface {
	Add(...interface{})
	Remove(...interface{})
	Contains(...interface{}) bool
	Len() int
	List() []interface{}
}

var _ Set = (*hashSet)(nil)

type hashSet struct {
	items map[interface{}]bool
}

// NewSet retuns an initialised set object
func NewSet() *hashSet {
	return &hashSet {
		items: make(map[interface{}]bool),
	}
}

// Add adds one or more items to the set
func (s *hashSet) Add(values ...interface{}) {
	for _, value := range values {
		s.items[value] = true
	}
}

// Remove removes one or more items from the set  
func (s *hashSet) Remove (values ...interface{}) {
	for _, value := range values {
		delete(s.items, value)
	}
}

// Contains checks whether all items are present in the set
// If one of the items is not present, it will return false 
func (s *hashSet) Contains(values ...interface{}) bool {
	for _, value := range values {
		if _, ok := s.items[value]; !ok {
			return false
		}
	}

	return true
}

// List returns the set as a slice
func (s *hashSet) List() []interface{} {
	l := make([]interface{}, s.Len())
	for k := range s.items {
		l = append(l, k)
	}

	return l
}

// Len returns the lenght of the Set 
func (s *hashSet) Len() int {
	return len(s.items)
}