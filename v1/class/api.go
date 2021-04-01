package class

import (
	"github.com/zytekaron/class/v1/storage"
)

// Open the database connection
//
// This will connect to the current working directory of the program,
// and create a `./class` database folder if one does not exist.
func Open() error {
	return storage.Open()
}

// Close the database connection
func Close() error {
	return storage.Close()
}
