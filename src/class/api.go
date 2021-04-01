package class

import (
	"class/src/storage"
)

func Init() error {
	return storage.Init()
}

func Close() error {
	return storage.Close()
}
