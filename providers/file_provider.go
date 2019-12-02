package providers

import (
	"encoding/json"
	"os"

	"github.com/blokje5/dag_config_manager/storage"
)

// FileProvider implements the provider interface
// and provides the ability to create files on
// a remote host.
type FileProvider struct {
	store storage.Storage
	state fileState
}

type fileState struct {
	path string `json:omitempty`
}

func (p *FileProvider) Init(store storage.Storage, stateJSON []byte) {
	p.store = store
	var state fileState
	json.Unmarshal(stateJSON, &state)
	p.state = state
}

func (p *FileProvider) Detect(id string) {
	stateJSON, ok := p.store.Read(id)
	if ok {
		currentState := stateJSON.(fileState)
		if _, err := os.Stat(currentState.path); os.IsNotExist(err) {
			// TODO implement required state changes after detect
		}
	}
}

func (p *FileProvider) Reconcile() {
	// TODO reconcile current state with desired state
}
