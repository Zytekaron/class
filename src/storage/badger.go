package storage

import (
	"class/src/types"
	"github.com/dgraph-io/badger/v3"
)

func Load(name string) (*types.Class, error) {
	return load(name, false)
}

func Save(class *types.Class) error {
	bytes, err := class.ToBytes()
	if err != nil {
		return err
	}

	return db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(class.ID), bytes)
	})
}

func load(id string, upsert bool) (*types.Class, error) {
	var data []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(id))
		if err != nil {
			return err
		}

		d, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}

		data = d
		return nil
	})
	if err == badger.ErrKeyNotFound && upsert {
		class := types.NewClass(id)
		err = Save(class)
		return class, err
	}
	if err != nil {
		return nil, err
	}

	return types.ParseClass(data)
}

