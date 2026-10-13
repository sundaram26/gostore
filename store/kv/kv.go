package kv

import (
	"fmt"
	"slices"
	"strconv"
	"sync"
	"time"

	"github.com/sundaram26/gostore/store"
)


type Store struct {
	mu sync.RWMutex
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

	s.mu.RLock()
	defer s.mu.RUnlock()
	
	for k, v := range s.data {
		cp.data[k] = v
	}

	return cp
}

func (s *Store) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

// Keys returns every key currently in the store, sorted alphabetically.
func (s *Store) Keys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
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

	s.mu.RLock()
	defer s.mu.RUnlock()

	time.Sleep(time.Millisecond) //simulate some work

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

	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.data[key]

	// we are at capacity && is this a new key
	if s.maxSize > 0 && len(s.data) >= s.maxSize && !exists {
		return fmt.Errorf("Set(%q): %w", key, store.ErrStoreFull)
	}
	
	s.data[key] = value
	return nil
}

func (s *Store) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}

// Incr increments the integer value stored at a key by q.
// If the key doesn't exists, it's initialized to 0 before incrementing.
// Mirrors Redis INCR command.
func (s *Store) Incr(key string) (int, error) {	
	if key == "" {
		return 0, store.ErrEmptyKey
	}
	
	s.mu.Lock()
	defer s.mu.Unlock()
	
	// Read
	current := 0
	if raw, ok := s.data[key]; ok {
		v, err := strconv.Atoi(raw) // string -> int
		if err != nil {
			return 0, fmt.Errorf("INCR %q: %w", key, err)
		}
		current = v
	}

	// Modify
	current++
	time.Sleep(time.Microsecond) //simulate some work

	// Write
	s.data[key] = strconv.Itoa(current) // int -> string

	return current, nil	
}