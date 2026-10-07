package main

import (
	"errors"
	"fmt"
	"slices"
	"encoding/base64"
)

var ErrKeyDoesNotExist = errors.New("Key does not exist")
var ErrEmptyKey = errors.New("empty key")
var ErrStoreFull = errors.New("Store is full")

type Store struct {
	data map[string]string
	maxSize int // 0 = unlimited
}

func NewStore(maxSize int) *Store {

	if maxSize < 0 {
		fmt.Println("Size should be greater than 0")
	}
	
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) SetKeyWithEncryption(key, val string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(val))
	if err := s.Set(key, encoded); err != nil {
		return "", err
	}

	return s.Get(key)
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
		return "", ErrEmptyKey
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
		return ErrEmptyKey
	}

	_, exists := s.data[key]

	// we are at capacity && is this a new key
	if s.maxSize > 0 && s.Len() >= s.maxSize && !exists {
		return fmt.Errorf("Set(%q): %w", key, ErrStoreFull)
	}
	
	s.data[key] = value
	return nil
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}