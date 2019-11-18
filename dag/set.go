package dag


// Set is a set data structure
type Set interface {
	Add(...Hashable)
	Remove(...Hashable)
	Contains(...Hashable) bool
	Len() int
	List() []Hashable
}

var _ Set = (*hashSet)(nil)

type hashSet struct {
	items map[int]Hashable
}

// NewSet retuns an initialised set object
func NewSet() *hashSet {
	return &hashSet {
		items: make(map[int]Hashable),
	}
}

// Add adds one or more items to the set
func (s *hashSet) Add(values ...Hashable) {
	for _, v := range values {
		s.items[v.Hashcode()] = v
	}
}

// Remove removes one or more items from the set  
func (s *hashSet) Remove (values ...Hashable) {
	for _, v := range values {
		delete(s.items, v.Hashcode())
	}
}

// Contains checks whether all items are present in the set
// If one of the items is not present, it will return false 
func (s *hashSet) Contains(values ...Hashable) bool {
	for _, v := range values {
		if _, ok := s.items[v.Hashcode()]; !ok {
			return false
		}
	}

	return true
}

// List returns the set as a slice
func (s *hashSet) List() []Hashable {
	l := make([]Hashable, s.Len())
	for _, k := range s.items {
		l = append(l, k)
	}

	return l
}

// Len returns the lenght of the Set 
func (s *hashSet) Len() int {
	return len(s.items)
}