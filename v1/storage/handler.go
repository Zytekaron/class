package storage

import (
	"github.com/zytekaron/class/v1/types"
)

// A function that takes in a class pointer and
// returns whether it was updated (should be saved)
type ClassFunc func(*types.Class) bool

// View a classifile by its id.
//
// Returns true and calls the callback if the class exists.
func Handle(id string, cb ClassFunc) error {
	return handle(id, cb, false)
}

// View a classifile, create it if it doesn't exist, and call the callback.
func EnsureHandle(id string, cb ClassFunc) error {
	return handle(id, cb, true)
}

func handle(id string, cb func(*types.Class) bool, upsert bool) error {
	class, err := load(id, upsert)
	if err != nil {
		return err
	}

	changed := cb(class)
	if !changed {
		return nil
	}

	err = Save(class)
	return err
}
