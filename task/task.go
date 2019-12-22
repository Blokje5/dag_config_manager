package task

// Task represents an operation to be performed by a provider
type Task struct {
	// Name is the unique identifier of the task
	Name string
	// Provider is the name of the provider implementing the task
	Provider string
}