package class

import (
	"github.com/zytekaron/class/v1/storage"
	"github.com/zytekaron/class/v1/types"
)

// Get the name of a class
func GetName(id string) (name string, err error) {
	err = storage.Handle(id, func(class *types.Class) bool {
		name = class.Name
		return false
	})
	return
}

// Set the name of a class
func SetName(id string, name string) (err error) {
	err = storage.EnsureHandle(id, func(class *types.Class) bool {
		class.Name = name
		return true
	})
	return
}
