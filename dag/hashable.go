package dag

// Hashable represents a hashable entity
type Hashable interface {
	Hashcode() int
}
