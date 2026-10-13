package ttl

import (
	"fmt"
	"slices"
	// "strconv"
	"time"

	"github.com/sundaram26/gostore/store"
)

// ==========================================================================
// TTLStore --> A second engine that expires keys automatically
// ==========================================================================
// TTLStore wraps a map but adds a expiry time to every value.
// Get returns ErrKeyNotFound for expired keys, exactly like a missing key.

type TTLStore struct {
	data map[string]ttlEntry
	ttl time.Duration
}

type ttlEntry struct {
	value string
	expiresAt time.Time
}

func NewTTLStore(ttl time.Duration) *TTLStore {
	return &TTLStore {
		data: make(map[string]ttlEntry),
		ttl: ttl,
	}
}

func (s *TTLStore) Keys() []string {
	keys := make([]string, 0, len(s.data))
	// sort the keys alphabetically
	for key := range s.data {
		keys = append(keys, key)
	}

	slices.Sort(keys)

	return keys
}

// func (t *TTLStore) SetKeyWithEncryption(key, val string, ttl time.Duration) (string, error) {
// 	encoded := base64.StdEncoding.EncodeToString([]byte(val))
// 	if err := t.Set(key, encoded); err != nil {
// 		return "", err
// 	}

// 	return t.Get(key)
// }

func (t *TTLStore) Set(key string, value string) error {
	if key == "" {
		return store.ErrEmptyKey
	}

	expiryTime := time.Now().Add(t.ttl)

	t.data[key] = ttlEntry{
		value: value,
		expiresAt: expiryTime,
	}

	return nil	
}

func (t *TTLStore) Get(key string) (string, error) {
	if key == ""{
		return "", store.ErrEmptyKey
	}

	entry, ok := t.data[key]

	if !ok || time.Now().After(entry.expiresAt) {
		// lazy eviction - no background job scanning for expired keys
		// deleted on demand.
		delete(t.data, key)

		return "", fmt.Errorf("key %s does not exists", key)
	}
	return entry.value, nil
}

func (t *TTLStore) Delete(key string) {
	delete(t.data, key)
}

func (t *TTLStore) Len() int {
	return len(t.data)
}

func (t *TTLStore) Incr(key string) (int, error) {
	// if key == "" {
	// 	return 0, store.ErrEmptyKey
	// }
	
	// // Read
	// current := 0
	// if raw, ok := t.data[key]; ok {
	// 	v, err := strconv.Atoi(raw) // string -> int
	// 	if err != nil {
	// 		return 0, fmt.Errorf("INCR %q: %w", key, err)
	// 	}
	// 	current = v
	// }

	// // Modify
	// current++
	// time.Sleep(time.Microsecond) //simulate some work

	// // Write
	// t.data[key] = strconv.Itoa(current) // int -> string

	// return current, nil
	return 0, nil	
}