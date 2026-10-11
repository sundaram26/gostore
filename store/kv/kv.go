package kv

import (
	"fmt"
	"slices"

	"github.com/sundaram26/gostore/store"
)


type Store struct {
	data map[string]string
	maxSize int // 0 = unlimited
}

func NewStore(maxSize int) *Store {
	return &Store{
		data: make(map[string]string),
		maxSize: maxSize,
	}
}

func (s *Store) Clone() *Store {
	cp := &Store {
		data: make(map[string]string, len(s.data)),
		maxSize: s.maxSize,
	}

	for k, v := range s.data {
		cp.data[k] = v
	}

	return cp
}

func (s *Store) Len() int {
	return len(s.data)
}

// Keys returns every key currently in the store, sorted alphabetically.
func (s *Store) Keys() []string {
	keys := make([]string, 0, len(s.data))
	// sort the keys alphabetically
	for key := range s.data {
		keys = append(keys, key)
	}

	slices.Sort(keys)

	return keys
}

func (s *Store) Get(key string) (string, error) {
	if key == ""{
		return "", store.ErrEmptyKey
	}

	val, ok := s.data[key]
	if !ok {
		return "", fmt.Errorf("key %s does not exists", key)
	}
	return val, nil
}

// same function as above
// func Get(s *Store, key string) (string, bool) {

// }

func (s *Store) Set(key, value string) error {
	if key == "" {
		return store.ErrEmptyKey
	}

	_, exists := s.data[key]

	// we are at capacity && is this a new key
	if s.maxSize > 0 && s.Len() >= s.maxSize && !exists {
		return fmt.Errorf("Set(%q): %w", key, store.ErrStoreFull)
	}
	
	s.data[key] = value
	return nil
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}