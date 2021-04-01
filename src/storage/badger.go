package storage

import (
	"class/src/types"
	"github.com/dgraph-io/badger/v3"
)

func Load(name string) (*types.Class, error) {
	return load(name, false)
}

func Create(id string) (*types.Class, error) {
	class := types.NewClass(id)
	return class, Save(class)
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

func Delete(id string) error {
	return db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(id))
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
		return Create(id)
	}
	if err != nil {
		return nil, err
	}

	return types.ParseClass(data)
}
