package provider

import (
	"reflect"
	"testing"

	"github.com/blokje5/dag_config_manager/providers"
	"github.com/blokje5/dag_config_manager/storage"
)

func TestInMemRegistry_RegisterAndFind(t *testing.T) {
	fileProvider := providers.FileProvider{}
	registry := NewInMemRegistry(storage.NewStore())
	registry.Register(&fileProvider)
	p, ok := registry.FindProvider(fileProvider.Register())
	if !ok {
		t.Fatal("Expected to find provider in registry")
	}

	if !reflect.DeepEqual(p, &fileProvider) {
		t.Fatal("Expected returned provider to be equal to the registered provider")
	}
}
