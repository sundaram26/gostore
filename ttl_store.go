package main

import (
	"encoding/base64"
	"fmt"
	"time"
)

// ==========================================================================
// TTLStore --> A second engine that expires keys automatically
// ==========================================================================
// TTLStore wraps a map but adds a expiry time to every value.
// Get returns ErrKeyNotFound for expired keys, exactly like a missing key.

type TTLStore struct {
	data map[string]ttlEntry
}

type ttlEntry struct {
	value string
	expiresAt time.Time
}

func NewTTLStore() *TTLStore {
	return &TTLStore {
		data: make(map[string]ttlEntry),
	}
}

func (t *TTLStore) SetKeyWithEncryption(key, val string, ttl time.Duration) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(val))
	if err := t.Set(key, encoded, ttl); err != nil {
		return "", err
	}

	return t.GetWithTTL(key)
}

func (t *TTLStore) Set(key string, value string, ttl time.Duration) error {
	if key == "" {
		return ErrEmptyKey
	}

	expiryTime := time.Now().Add(ttl)

	t.data[key] = ttlEntry{
		value: value,
		expiresAt: expiryTime,
	}

	return nil	
}

func (t *TTLStore) GetWithTTL(key string) (string, error) {
	if key == ""{
		return "", ErrEmptyKey
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