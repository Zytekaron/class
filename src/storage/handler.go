package storage

import (
	"class/src/types"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"log"
	"os"
)

// A function that takes in a class pointer and
// returns whether it was updated (should be saved)
type ClassFunc func(*types.Class) bool

// View a classifile by its id.
//
// Returns true and calls the callback if the class exists.
func Handle(id string, cb ClassFunc) bool {
	return handle(id, cb, false)
}

// View a classifile, create it if it doesn't exist, and call the callback.
func EnsureHandle(id string, cb ClassFunc) {
	handle(id, cb, true)
}

func handle(id string, cb func(*types.Class) bool, upsert bool) bool {
	class, err := load(id, upsert)
	// may only occur when !upsert
	if err == badger.ErrKeyNotFound {
		return false
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	changed := cb(class)
	if !changed {
		return true
	}

	err = Save(class)
	if err != nil {
		log.Fatal(err)
	}
	return true
}
