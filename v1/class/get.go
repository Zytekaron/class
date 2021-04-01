package class

import (
	"github.com/zytekaron/class/v1/storage"
	"github.com/zytekaron/class/v1/types"
)

func Get(id string) (class *types.Class, err error) {
	return storage.Load(id)
}
