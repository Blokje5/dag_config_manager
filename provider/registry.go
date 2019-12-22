package provider

import (
	"github.com/blokje5/dag_config_manager/storage"
)


// Registry hold a map of names to providers
type Registry interface {
	Register(Provider)
	FindProvider(n string) (Provider, bool)
}

// InMemRegistry returns an in-memory registry
type InMemRegistry struct {
	store storage.Storage
}

// NewInMemRegistry returns a new in memory registry
func NewInMemRegistry(store storage.Storage) (*InMemRegistry) {
	return &InMemRegistry{
		store: store,
	}
}

// Register registers a provider with the registry
func (r *InMemRegistry) Register(p Provider) {
	n := p.Register()
	r.store.Write(n, p)
}

// FindProvider returns the provider for the given provider name
// If no provider is found, the boolean is false.
func (r *InMemRegistry) FindProvider(n string) (Provider, bool) {
	p, ok := r.store.Read(n)
	if !ok {
		return nil, false
	}

	return p.(Provider), true
}