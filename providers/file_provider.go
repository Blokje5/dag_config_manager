package providers

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/blokje5/dag_config_manager/provider/state"
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

func (p *FileProvider) Detect(id string) ([]byte, []byte, error) {
	stateJSON, ok := p.store.Read(id)
	if ok {
		currentState := stateJSON.(fileState)
		if _, err := os.Stat(currentState.path); os.IsNotExist(err) {
			currentState.path = ""
			cur, err := json.Marshal(currentState)
			if err != nil {
				return nil, nil, err
			}

			next, err := json.Marshal(p.state)
			if err != nil {
				return nil, nil, err
			}

			return cur, next, nil
		}
	}

	return nil, nil, errors.New("Could not read state")
}

func (p *FileProvider) Reconcile(ops []state.Operation) {
	for _, op := range ops {
		switch op.(type) {
		case state.Create:
			create := op.(state.Create)
			if create.Key == "path" {
				path := create.Value.(string)
				os.Create(path)
			}
		case state.Delete:
			delete := op.(state.Delete)
			if delete.Key == "path" {
				path := delete.Value.(string)
				os.Remove(path)
			}
		case state.Update:
			update := op.(state.Update)
			if update.Key == "path" {
				pathSrc := update.Before.(string)
				pathTarget := update.Before.(string)
				os.Rename(pathSrc, pathTarget)
			}
		}
	}
}
