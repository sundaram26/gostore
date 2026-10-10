package store

import "errors"

var ErrKeyDoesNotExist = errors.New("Key does not exist")
var ErrEmptyKey = errors.New("empty key")
var ErrStoreFull = errors.New("Store is full")

type Storer interface {
	Get(key string) (string, error)
	Set(key, value string) error
	// Keys() []string
	Delete(key string)
	Len() int
}