package db

import (
	"github.com/zytekaron/class/v1/types"
)

type Database interface {
	// Open the dbName connection
	Open() error

	// Close the dbName connection
	Close() error

	// Get a class from the database by its id
	Get(string) (*types.Class, error)

	// Batch get classes from a list of ids
	Batch([]string) ([]*types.Class, error)

	// Get all classes from the database
	All() ([]*types.Class, error)

	// Get a class iterator containing all classes in the database
	Iterator() Iterator

	// Insert a class into the database
	Insert(*types.Class) error

	// Delete a class from the database
	Delete(string) error

	// Set a class name
	SetName(string, string) error

	// Find classes by their name
	FindName(string) (*types.Class, error)

	// Add a tag to a class
	AddTag(string, string) error

	// Add tags to a class
	AddTags(string, []string) error

	// Remove a tag from a class
	RemoveTag(string, string) error

	// Remove tags from a class
	RemoveTags(string, []string) error

	// Find classes that contain at least one of the selected tags
	FindAnyTags([]string) ([]*types.Class, error)

	// Find classes that contain all of the selected tags
	FindAllTags([]string) ([]*types.Class, error)

	// Add metadata to the class
	AddMeta(string, string, interface{}) error

	// Add metadata to the class
	AddMetaBulk(string, map[string]interface{}) error

	// Remove metadata from the class
	RemoveMeta(string, string) error

	// Remove metadata from the class
	RemoveMetaBulk(string, []string) error

	// Find classes that have the specified meta key regardless of value
	FindMetaExists(string) ([]*types.Class, error)

	// Find classes that have the specified meta key with a specific value
	FindMetaExact(string, interface{}) ([]*types.Class, error)
}

// A simple Iterator type which uses a channel to send values.
//
// Error only needs to be called once the channel is closed to check for any errors.
type Iterator interface {
	Chan() chan *types.Class
	Error() error
}

func stringsContains(slice []string, str string) bool {
	for _, e := range slice {
		if e == str {
			return true
		}
	}
	return false
}
