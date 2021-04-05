package types

import (
	"bytes"
	"encoding/gob"
)

type Class struct {
	ID   string                 `json:"_id" bson:"_id"`
	Name string                 `json:"name" bson:"name"`
	Tags TagList                `json:"tags" bson:"tags"`
	Meta map[string]interface{} `json:"meta" bson:"meta"`
}

type TagList []string

func NewClass(id string) *Class {
	return &Class{
		ID:   id,
		Tags: make([]string, 0),
		Meta: make(map[string]interface{}),
	}
}

func ParseClass(data []byte) (class *Class, err error) {
	reader := bytes.NewReader(data)
	err = gob.NewDecoder(reader).Decode(&class)
	return
}

func (c *Class) ToBytes() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(c)
	return buf.Bytes(), err
}

func (t *TagList) Add(tag string) {
	for _, e := range *t {
		if e == tag {
			return
		}
	}
	*t = append(*t, tag)
}

func (t *TagList) AddAll(tags []string) {
	for _, tag := range tags {
		t.Add(tag)
	}
}

func (t *TagList) Remove(tag string) {
	index := -1
	for i, e := range *t {
		if e == tag {
			index = i
			break
		}
	}

	slice := []string(*t)
	*t = append(slice[:index], slice[index+1:]...)
}

func (t *TagList) RemoveAll(tags []string) {
	for _, tag := range tags {
		t.Remove(tag)
	}
}

func (t *TagList) Has(tag string) bool {
	for _, e := range *t {
		if e == tag {
			return true
		}
	}
	return false
}

func (t *TagList) HasAll(tags []string) bool {
	for _, tag := range tags {
		if !t.Has(tag) {
			return false
		}
	}
	return true
}

func (t *TagList) HasAny(tags []string) bool {
	for _, tag := range tags {
		if t.Has(tag) {
			return true
		}
	}
	return false
}
