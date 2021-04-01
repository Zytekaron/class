package class

import "class/src/storage"

func Delete(id string) error {
	return storage.Delete(id)
}
