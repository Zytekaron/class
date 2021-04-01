package storage

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/zytekaron/class/v1/types"
)

// Get all classes from the db
func All() ([]*types.Class, error) {
	out := make([]*types.Class, 0)
	err := db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			err := item.Value(func(value []byte) error {
				c, err := types.ParseClass(value)
				if err != nil {
					return err
				}
				out = append(out, c)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	return out, err
}

// Batch get classes from the db by their id
func Batch(ids []string) ([]*types.Class, error) {
	i := 0
	out := make([]*types.Class, len(ids))
	err := db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			err := item.Value(func(value []byte) error {
				c, err := types.ParseClass(value)
				if err != nil {
					return err
				}
				if contains(ids, c.ID) {
					out[i] = c
					i++
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	return out, err
}
