package task

import (
	"fmt"

	"github.com/blokje5/dag_config_manager/provider"
	"gopkg.in/yaml.v2"
)

// Task represents an operation to be performed by a provider
type Task struct {
	// Name is the unique identifier of the task
	Name string `yaml:"name"`
	// Provider is the name of the provider implementing the task
	Provider string `yaml:"provider"`
	// Data represents the leftover keys and values that are used to configure the task
	Data map[string]interface{} `yaml:",inline"`
}

// Parse parses a set of bytes into a task struct
func Parse(data []byte) (Task, error) {
	var task Task
	if err := yaml.Unmarshal(data, &task); err != nil {
		return task, err
	}
	return task, nil
}

func (t Task) Initialize(registry provider.Registry) error {
	provider, ok := registry.FindProvider(t.Provider)
	if !ok {
		return fmt.Errorf("Unable to find provider for task: %v", t.Name)
	}

	
}
