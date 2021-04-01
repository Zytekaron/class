package class

import (
	"class/src/storage"
	"class/src/types"
)

func GetMeta(id string) (meta map[string]string, err error) {
	err = storage.Handle(id, func(class *types.Class) bool {
		meta = class.Meta
		return false
	})
	return
}

func AddMeta(id, key, value string) (err error) {
	err = storage.EnsureHandle(id, func(class *types.Class) bool {
		class.Meta[key] = value
		return true
	})
	return
}

func BulkAddMeta(id string, meta map[string]string) (err error) {
	err = storage.EnsureHandle(id, func(class *types.Class) bool {
		for key, value := range meta {
			class.Meta[key] = value
		}
		return true
	})
	return
}

func RemoveMeta(id, key string) (err error) {
	err = storage.EnsureHandle(id, func(class *types.Class) bool {
		delete(class.Meta, key)
		return true
	})
	return
}

func BulkRemoveMeta(id string, keys []string) (err error) {
	err = storage.EnsureHandle(id, func(class *types.Class) bool {
		for _, key := range keys {
			delete(class.Meta, key)
		}
		return true
	})
	return
}

func SetMeta(id string, meta map[string]string) (err error) {
	err = storage.EnsureHandle(id, func(class *types.Class) bool {
		class.Meta = meta
		return true
	})
	return
}

func ClearMeta(id string) (err error) {
	err = storage.EnsureHandle(id, func(class *types.Class) bool {
		class.Meta = map[string]string{}
		return true
	})
	return
}
