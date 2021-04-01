package types

import (
	"bytes"
	"encoding/gob"
)

type Class struct {
	ID   string
	Name string
	Desc string
	Tags TagList
	Meta map[string]interface{}
}

type TagList []string

func NewClass(id string) *Class {
	return &Class{
		ID: id,
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
