package storage

import (
	"github.com/dgraph-io/badger/v3"
	"os"
	"path"
)

var pwd string
var db *badger.DB

func init() {
	var err error
	pwd, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}

func Init() error {
	dir := path.Join(pwd, ".class")
	opts := badger.DefaultOptions(dir).WithLogger(nil)

	database, err := badger.Open(opts)
	db = database
	return err
}

func Close() error {
	return db.Close()
}
