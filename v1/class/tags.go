package class

import (
	"github.com/zytekaron/class/v1/storage"
	"github.com/zytekaron/class/v1/types"
)

// Get the tags of a class
func GetTags(id string) (tags []string, err error) {
	err = storage.Handle(id, func(class *types.Class) bool {
		tags = class.Tags
		return false
	})
	return
}

// Add a tag to a class
func AddTag(id string, tag string) (err error) {
	return AddTags(id, []string{tag})
}

// Add multiple tags to a class
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

// Remove a tag from a class
func RemoveTag(id string, tag string) (err error) {
	return RemoveTags(id, []string{tag})
}

// Remove multiple tags from a class
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

// Clear all tags of a class
func ClearTags(id string) (err error) {
	err = storage.EnsureHandle(id, func(class *types.Class) bool {
		class.Tags = []string{}
		return true
	})
	return
}
