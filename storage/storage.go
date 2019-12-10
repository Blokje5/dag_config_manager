package storage

// Storage defines the interface for the storage backend.
// The storage backend is used to maintain metadata on the current execution.
type Storage interface {
	// Read the key from the storage backend
	Read(key string) (interface{}, bool)
	// Write the key to the storage backend
	Write(key string, value interface{})
}

type memStore struct {
	data map[string]interface{}
}

// NewStore returns an in memory storage backend.
func NewStore() Storage {
	return &memStore{
		data: map[string]interface{}{},
	}
}

func (s *memStore) Read(key string) (interface{}, bool) {
	value, ok := s.data[key]
	return value, ok
}

func (s *memStore) Write(key string, value interface{}) {
	s.data[key] = value
}
