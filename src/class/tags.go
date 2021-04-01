package class

import (
	"class/src/storage"
	"class/src/types"
)

func GetTags(id string) (tags []string, err error) {
	err = storage.Handle(id, func(class *types.Class) bool {
		tags = class.Tags
		return false
	})
	return
}

func AddTag(id string, tag string) (err error) {
	return AddTags(id, []string{tag})
}

func AddTags(id string, tags []string) (err error) {
	if len(tags) == 0 {
		return nil
	}
	err = storage.EnsureHandle(id, func(class *types.Class) bool {
		for _, tag := range tags {
			class.Tags.Add(tag)
		}
		return true
	})
	return
}

func RemoveTag(id string, tag string) (err error) {
	return RemoveTags(id, []string{tag})
}

func RemoveTags(id string, tags []string) (err error) {
	if len(tags) == 0 {
		return nil
	}
	err = storage.EnsureHandle(id, func(class *types.Class) bool {
		for _, tag := range tags {
			class.Tags.Remove(tag)
		}
		return true
	})
	return
}

func ClearTags(id string) (err error) {
	err = storage.EnsureHandle(id, func(class *types.Class) bool {
		class.Tags = []string{}
		return true
	})
	return
}
