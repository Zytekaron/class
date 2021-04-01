package class

import (
	"github.com/zytekaron/class/v1/storage"
	"github.com/zytekaron/class/v1/types"
)

func Create(id string) (*types.Class, error) {
	return storage.Create(id)
}
