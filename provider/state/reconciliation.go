package state

import "reflect"

// Reconciliation captures the changes made by the provider
type Reconciliation struct {
	// Before captures the previous state stored by the provider
	// and is normally retrieved from storage
	Before map[string]interface{} `json:"before"`
	// After captures the state that the provider should aim to achieve
	// This is the new state that will be stored if the provider is succesful
	After map[string]interface{} `json:"after"`
}

// Diff returns the operations required to go from the before state
// to the after state
func (r *Reconciliation) Diff() []Operation {
	before := r.Before
	after := r.After
	var operations []Operation
	// Check for the trivial case of no change
	if reflect.DeepEqual(before, after) {
		return operations
	}

	commonKeys := findCommonKeys(before, after)
	creations := findAddedKeys(before, after)
	updates := findUpdatedValues(before, after, commonKeys)
	deletions := findDeletedKeys(before, after)

	for _, c := range creations {
		op := c
		operations = append(operations, op)
	}

	for _, u := range updates {
		op := u
		operations = append(operations, op)
	}

	for _, d := range deletions {
		op := d
		operations = append(operations, op)
	}

	return operations
}

// Operation represents a step to be taken by the
// provider in the reconciliation loop
type Operation interface{}

func findCommonKeys(before, after map[string]interface{}) []string {
	var commonKeys []string
	for k := range before {
		if _, ok := after[k]; ok {
			commonKeys = append(commonKeys, k)
		}
	}

	return commonKeys
}

// Create captures newly added keys.
// The reconciliation loop
// will be responsible for correctly reading the values
// And executing the creation
type Create struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func findAddedKeys(before, after map[string]interface{}) []Create {
	var createdKeys []Create
	for k := range after {
		if _, ok := before[k]; !ok {
			create := Create{
				Key:   k,
				Value: after[k],
			}
			createdKeys = append(createdKeys, create)
		}
	}

	return createdKeys
}

// Update captures the change for a key
// if a value is updated. The reconciliation loop
// will be responsible for correctly reading the values
// And executing the update
type Update struct {
	Key    string      `json:"key"`
	Before interface{} `json:"before"`
	After  interface{} `json:"after"`
}

func findUpdatedValues(before, after map[string]interface{}, commonKeys []string) []Update {
	var updates []Update
	for _, k := range commonKeys {
		beforeValue := before[k]
		afterValue := after[k]
		if !isEqual(beforeValue, afterValue) {
			update := Update{
				Key:    k,
				Before: beforeValue,
				After:  afterValue,
			}

			updates = append(updates, update)
		}
	}

	return updates
}

// Checks for equality between two interfaces
func isEqual(v1, v2 interface{}) bool {
	// TODO make a rigorous implementation checking for non-comparable types
	return v1 == v2
}

// Delete captures removed keys.
// The reconciliation loop
// will be responsible for correctly reading the values
// And executing the deletion
type Delete struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func findDeletedKeys(before, after map[string]interface{}) []Delete {
	var deletedKeys []Delete
	for k := range before {
		if _, ok := after[k]; !ok {
			delete := Delete{
				Key:   k,
				Value: before[k],
			}
			deletedKeys = append(deletedKeys, delete)
		}
	}

	return deletedKeys
}
