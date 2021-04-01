package class

import (
	"github.com/zytekaron/class/v1/storage"
)

func Init() error {
	return storage.Init()
}

func Close() error {
	return storage.Close()
}
