package class

import (
	"github.com/zytekaron/class/v1/storage"
	"github.com/zytekaron/class/v1/types"
)

func GetDesc(id string) (desc string, err error) {
	err = storage.Handle(id, func(class *types.Class) bool {
		desc = class.Desc
		return false
	})
	return
}

func SetDesc(id string, desc string) (err error) {
	err = storage.EnsureHandle(id, func(class *types.Class) bool {
		class.Desc = desc
		return true
	})
	return
}
