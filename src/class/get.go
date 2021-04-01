package class

import (
	"class/src/storage"
	"class/src/types"
)

func Get(id string) (class *types.Class, err error) {
	err = storage.Handle(id, func(c *types.Class) bool {
		class = c
		return false
	})
	return
}
