package class

import (
	"class/src/storage"
	"class/src/types"
)

func Create(id string) (*types.Class, error) {
	return storage.Create(id)
}
