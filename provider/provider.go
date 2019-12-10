package provider

import (
	"github.com/blokje5/dag_config_manager/provider/state"
	"github.com/blokje5/dag_config_manager/storage"
)

// The Provider interface encapsulates a provider for
// tasks to be executed on a host. Each provider
// is passed a connection and a state store in order to
// perform operations on a remote system.
//
// The expected state is passed to the provider
// in the form of a json string.
type Provider interface {
	// Init initializes the provider with the required information
	Init(store storage.Storage, stateJSON []byte)
	// Detect determines the current state on the host
	// and should update the storage backend with the
	// required information to run the reconciliation loop.
	// A task id is passed to determine whether a previous state was available
	Detect(id string) ([]byte, []byte, error)

	// Reconcile takes in the current requested state
	// and reconciles the host
	Reconcile(ops []state.Operation)
}
