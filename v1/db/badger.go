package db

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/zytekaron/class/v1/types"
)

type BadgerDatabase struct {
	dir string
	db  *badger.DB
}

func NewBadger(dir string) *BadgerDatabase {
	return &BadgerDatabase{dir: dir}
}

func (b *BadgerDatabase) Open() error {
	opts := badger.DefaultOptions(b.dir).WithLogger(nil)
	var err error
	b.db, err = badger.Open(opts)
	return err
}

func (b *BadgerDatabase) Close() error {
	return b.db.Close()
}

func (b *BadgerDatabase) Get(id string) (*types.Class, error) {
	var data []byte
	err := b.db.View(func(txn *badger.Txn) error {
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
	if err != nil {
		return nil, err
	}

	return types.ParseClass(data)
}

// note: this badger function doesn't use iterators, because deserializing classes from bytes is slow.
func (b *BadgerDatabase) Batch(ids []string) ([]*types.Class, error) {
	i := 0
	out := make([]*types.Class, len(ids))
	err := b.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			if !stringsContains(ids, string(item.Key())) {
				continue
			}
			return item.Value(func(value []byte) error {
				c, err := types.ParseClass(value)
				if err != nil {
					return err
				}
				out[i] = c
				i++
				return nil
			})
		}
		return nil
	})
	return out, err
}

func (b *BadgerDatabase) All() ([]*types.Class, error) {
	out := make([]*types.Class, 0)
	it := b.Iterator()
	for c := range it.Chan() {
		out = append(out, c)
	}
	return out, it.Error()
}

func (b *BadgerDatabase) Iterator() Iterator {
	ch := make(chan *types.Class)
	it := &BadgerIterator{ch: ch}

	go func() {
		defer close(ch)

		err := b.db.View(func(txn *badger.Txn) error {
			badgerIterator := txn.NewIterator(badger.DefaultIteratorOptions)
			defer badgerIterator.Close()

			for badgerIterator.Rewind(); badgerIterator.Valid(); badgerIterator.Next() {
				err := badgerIterator.Item().Value(func(value []byte) error {
					c, err := types.ParseClass(value)
					if err != nil {
						return err // propagate to outer for loop
					}
					it.ch <- c
					return nil
				})
				if err != nil {
					return err // propagate to b.database.View
				}
			}
			return nil
		})
		if err != nil {
			it.err = err
		}
	}()

	return it
}

func (b *BadgerDatabase) Insert(c *types.Class) error {
	bytes, err := c.ToBytes()
	if err != nil {
		return err
	}

	return b.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(c.ID), bytes)
	})
}

func (b *BadgerDatabase) Delete(id string) error {
	return b.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(id))
	})
}

func (b *BadgerDatabase) SetName(id string, name string) error {
	return b.update(id, func(class *types.Class) {
		class.Name = name
	})
}

func (b *BadgerDatabase) FindName(name string) (*types.Class, error) {
	return b.findOne(func(class *types.Class) bool {
		return class.Name == name
	})
}

func (b *BadgerDatabase) AddTag(id, tags string) error {
	return b.update(id, func(class *types.Class) {
		class.Tags.Add(tags)
	})
}

func (b *BadgerDatabase) AddTags(id string, tags []string) error {
	return b.update(id, func(class *types.Class) {
		class.Tags.AddAll(tags)
	})
}

func (b *BadgerDatabase) RemoveTag(id, tags string) error {
	return b.update(id, func(class *types.Class) {
		class.Tags.Remove(tags)
	})
}

func (b *BadgerDatabase) RemoveTags(id string, tags []string) error {
	return b.update(id, func(class *types.Class) {
		class.Tags.RemoveAll(tags)
	})
}

func (b *BadgerDatabase) FindAnyTags(tags []string) ([]*types.Class, error) {
	return b.findMany(func(class *types.Class) bool {
		return class.Tags.HasAny(tags)
	})
}

func (b *BadgerDatabase) FindAllTags(tags []string) ([]*types.Class, error) {
	return b.findMany(func(class *types.Class) bool {
		return class.Tags.HasAll(tags)
	})
}

func (b *BadgerDatabase) AddMeta(id, key string, value interface{}) error {
	return b.update(id, func(class *types.Class) {
		class.Meta[key] = value
	})
}

func (b *BadgerDatabase) AddMetaBulk(id string, meta map[string]interface{}) error {
	return b.update(id, func(class *types.Class) {
		for key, value := range meta {
			class.Meta[key] = value
		}
	})
}

func (b *BadgerDatabase) RemoveMeta(id, key string) error {
	return b.update(id, func(class *types.Class) {
		delete(class.Meta, key)
	})
}

func (b *BadgerDatabase) RemoveMetaBulk(id string, keys []string) error {
	return b.update(id, func(class *types.Class) {
		for _, key := range keys {
			delete(class.Meta, key)
		}
	})
}

func (b *BadgerDatabase) FindMetaExists(key string) ([]*types.Class, error) {
	return b.findMany(func(class *types.Class) bool {
		_, ok := class.Meta[key]
		return ok
	})
}

func (b *BadgerDatabase) FindMetaExact(key string, value interface{}) ([]*types.Class, error) {
	return b.findMany(func(class *types.Class) bool {
		val, ok := class.Meta[key]
		return ok && val == value
	})
}

func (b *BadgerDatabase) update(id string, predicate func(class *types.Class)) error {
	doc, err := b.Get(id)
	if err != nil {
		return err
	}
	predicate(doc)
	return b.Insert(doc)
}

func (b *BadgerDatabase) findOne(predicate func(class *types.Class) bool) (*types.Class, error) {
	it := b.Iterator()
	for c := range it.Chan() {
		if predicate(c) {
			return c, nil
		}
	}
	return nil, it.Error()
}

func (b *BadgerDatabase) findMany(predicate func(class *types.Class) bool) ([]*types.Class, error) {
	out := make([]*types.Class, 0)
	it := b.Iterator()
	for c := range it.Chan() {
		if predicate(c) {
			out = append(out, c)
		}
	}
	return out, it.Error()
}

type BadgerIterator struct {
	ch  chan *types.Class
	err error
}

func (b *BadgerIterator) Chan() chan *types.Class {
	return b.ch
}

func (b *BadgerIterator) Error() error {
	return b.err
}
