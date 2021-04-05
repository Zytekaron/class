package db

import (
	"context"
	"github.com/zytekaron/class/v1/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	uri        string
	dbName     string
	colName    string
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewMongo(uri, db, collection string) *MongoDatabase {
	return &MongoDatabase{
		uri:     uri,
		dbName:  db,
		colName: collection,
	}
}

func (m *MongoDatabase) Open() error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(m.uri))
	if err != nil {
		return err
	}
	m.client = client
	m.database = client.Database(m.dbName)
	m.collection = m.database.Collection(m.colName)
	return nil
}

func (m *MongoDatabase) Close() error {
	return m.client.Disconnect(context.Background())
}

func (m *MongoDatabase) Get(id string) (*types.Class, error) {
	res := m.collection.FindOne(context.Background(), bson.M{"_id": id})
	var class *types.Class
	err := res.Decode(&class)
	return class, err
}

// note: this badger function doesn't use iterators, because deserializing classes from bytes is slow.
func (m *MongoDatabase) Batch(ids []string) ([]*types.Class, error) {
	cursor, err := m.collection.Find(context.Background(), bson.M{"_id": bson.M{"$in": ids}})
	if err != nil {
		return nil, err
	}

	var results []*types.Class
	err = cursor.All(context.Background(), &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (m *MongoDatabase) All() ([]*types.Class, error) {
	out := make([]*types.Class, 0)
	it := m.Iterator()
	for c := range it.Chan() {
		out = append(out, c)
	}
	return out, it.Error()
}

func (m *MongoDatabase) Iterator() Iterator {
	ch := make(chan *types.Class)
	it := &MongoIterator{ch: ch}

	go func() {
		defer close(ch)

		cursor, err := m.collection.Find(context.Background(), bson.M{})
		if err != nil {
			it.err = err
			return
		}

		var results []*types.Class
		err = cursor.All(context.Background(), &results)
		if err != nil {
			it.err = err
			return
		}

		for _, class := range results {
			ch <- class
		}
	}()

	return it
}

func (m *MongoDatabase) Insert(class *types.Class) error {
	_, err := m.collection.InsertOne(context.Background(), class)
	return err
}

func (m *MongoDatabase) Delete(id string) error {
	_, err := m.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func (m *MongoDatabase) SetName(id string, name string) error {
	_, err := m.collection.UpdateByID(context.Background(), id, bson.M{"$set": bson.M{"name": name}})
	return err
}

func (m *MongoDatabase) FindName(name string) (*types.Class, error) {
	return m.findOne(func(class *types.Class) bool {
		return class.Name == name
	})
}

func (m *MongoDatabase) AddTag(id string, tag string) error {
	_, err := m.collection.UpdateByID(context.Background(), id, bson.M{"$push": bson.M{"tags": tag}})
	return err
}

func (m *MongoDatabase) AddTags(id string, tags []string) error {
	_, err := m.collection.UpdateByID(context.Background(), id, bson.M{"$push": bson.M{"tags": bson.M{"$each": tags}}})
	return err
}

func (m *MongoDatabase) RemoveTag(id string, tag string) error {
	_, err := m.collection.UpdateByID(context.Background(), id, bson.M{"$pull": bson.M{"tags": tag}})
	return err
}

func (m *MongoDatabase) RemoveTags(id string, tags []string) error {
	_, err := m.collection.UpdateByID(context.Background(), id, bson.M{"$pull": bson.M{"tags": bson.M{"$in": tags}}})
	return err
}

func (m *MongoDatabase) FindAnyTags(tags []string) ([]*types.Class, error) {
	cursor, err := m.collection.Find(context.Background(), bson.M{"tags": bson.M{"$in": tags}})
	return m.exhaustCursor(cursor, err)
}

func (m *MongoDatabase) FindAllTags(tags []string) ([]*types.Class, error) {
	cursor, err := m.collection.Find(context.Background(), bson.M{"tags": bson.M{"$all": tags}})
	return m.exhaustCursor(cursor, err)
}

func (m *MongoDatabase) AddMeta(id, key string, value interface{}) error {
	_, err := m.collection.UpdateByID(context.Background(), id, bson.M{"$set": bson.M{"meta." + key: value}})
	return err
}

func (m *MongoDatabase) AddMetaBulk(id string, data map[string]interface{}) error {
	doc := bson.M{}
	for key, value := range data {
		doc["meta."+key] = value
	}
	_, err := m.collection.UpdateByID(context.Background(), id, bson.M{"$set": doc})
	return err
}

func (m *MongoDatabase) RemoveMeta(id, key string) error {
	_, err := m.collection.UpdateByID(context.Background(), id, bson.M{"$unset": bson.M{"meta." + key: 1}})
	return err
}

func (m *MongoDatabase) RemoveMetaBulk(id string, keys []string) error {
	doc := bson.M{}
	for _, key := range keys {
		doc["meta."+key] = 1
	}
	_, err := m.collection.UpdateByID(context.Background(), id, bson.M{"$unset": doc})
	return err
}

func (m *MongoDatabase) FindMetaExists(key string) ([]*types.Class, error) {
	cursor, err := m.collection.Find(context.Background(), bson.M{"meta." + key: bson.M{"$exists": true}})
	return m.exhaustCursor(cursor, err)
}

func (m *MongoDatabase) FindMetaExact(key string, value interface{}) ([]*types.Class, error) {
	cursor, err := m.collection.Find(context.Background(), bson.M{"meta." + key: value})
	return m.exhaustCursor(cursor, err)
}

func (m *MongoDatabase) findOne(predicate func(class *types.Class) bool) (*types.Class, error) {
	it := m.Iterator()
	for c := range it.Chan() {
		if predicate(c) {
			return c, nil
		}
	}
	return nil, it.Error()
}

func (m *MongoDatabase) findMany(predicate func(class *types.Class) bool) ([]*types.Class, error) {
	out := make([]*types.Class, 0)
	it := m.Iterator()
	for c := range it.Chan() {
		if predicate(c) {
			out = append(out, c)
		}
	}
	return out, it.Error()
}

func (m *MongoDatabase) exhaustCursor(cursor *mongo.Cursor, err error) ([]*types.Class, error) {
	if err != nil {
		return nil, err
	}
	var list []*types.Class
	return list, cursor.All(context.Background(), &list)
}

type MongoIterator struct {
	ch  chan *types.Class
	err error
}

func (b *MongoIterator) Chan() chan *types.Class {
	return b.ch
}

func (b *MongoIterator) Error() error {
	return b.err
}
