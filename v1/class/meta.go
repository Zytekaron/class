package class

import (
	"github.com/zytekaron/class/v1/storage"
	"github.com/zytekaron/class/v1/types"
)

// Get all the meta of a class
func GetMeta(id string) (meta map[string]string, err error) {
	err = storage.Handle(id, func(class *types.Class) bool {
		meta = class.Meta
		return false
	})
	return
}

// Add meta of a class
func AddMeta(id, key, value string) (err error) {
	err = storage.EnsureHandle(id, func(class *types.Class) bool {
		class.Meta[key] = value
		return true
	})
	return
}

// Add a set of meta to a class
func BulkAddMeta(id string, meta map[string]string) (err error) {
	err = storage.EnsureHandle(id, func(class *types.Class) bool {
		for key, value := range meta {
			class.Meta[key] = value
		}
		return true
	})
	return
}

// Remove meta from a class
func RemoveMeta(id, key string) (err error) {
	err = storage.EnsureHandle(id, func(class *types.Class) bool {
		delete(class.Meta, key)
		return true
	})
	return
}

// Remove a set of meta from a class
func BulkRemoveMeta(id string, keys []string) (err error) {
	err = storage.EnsureHandle(id, func(class *types.Class) bool {
		for _, key := range keys {
			delete(class.Meta, key)
		}
		return true
	})
	return
}

// Overwrite a class's meta
func SetMeta(id string, meta map[string]string) (err error) {
	err = storage.EnsureHandle(id, func(class *types.Class) bool {
		class.Meta = meta
		return true
	})
	return
}

// Clear the meta of a class
func ClearMeta(id string) (err error) {
	err = storage.EnsureHandle(id, func(class *types.Class) bool {
		class.Meta = map[string]string{}
		return true
	})
	return
}
