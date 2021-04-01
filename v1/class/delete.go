package class

import "github.com/zytekaron/class/v1/storage"

func Delete(id string) error {
	return storage.Delete(id)
}
