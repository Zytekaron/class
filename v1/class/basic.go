package class

import (
	"github.com/zytekaron/class/v1/storage"
	"github.com/zytekaron/class/v1/types"
)

// Create a new class with empty meta
func Create(id string) (*types.Class, error) {
	return storage.Create(id)
}

// Save a class instance
func Save(class *types.Class) error {
	return storage.Save(class)
}

// Get a class by its id
func Get(id string) (class *types.Class, err error) {
	return storage.Get(id)
}

// Batch get classes by their ids
func Batch(ids []string) ([]*types.Class, error) {
	return storage.Batch(ids)
}

// Get all classes in the db
func All() ([]*types.Class, error) {
	return storage.All()
}

// Delete a class by its id
func Delete(id string) error {
	return storage.Delete(id)
}
