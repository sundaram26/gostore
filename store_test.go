package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestKeys_ReturnsAllKeysSorted(t *testing.T){
	store := NewStore(5)
	store.Set("charlie", "3")
	store.Set("alpha", "1")
	store.Set("brave", "2")

	got := store.Keys()
	want := []string{"alpha", "brave", "charlie"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("key() = %v, want %v", got, want)
	}
}

func TestKyes_EmptyStore(t *testing.T) {
	store := NewStore(2)

	got := store.Keys()

	if len(got) != 0 {
		t.Errorf("Keys() on empty store = %v, want empty slice", got)
	}
}

func TestSetGet_EmptyKeys(t *testing.T) {
	store := NewStore(1)

	if _, err := store.Get(""); err == nil || !errors.Is(err, ErrEmptyKey){
		t.Error("Get() failed")
	}
}

func TestSetGet_RoundTrip(t *testing.T) {
	store := NewStore(2)

	store.Set("hello", "world")

	if val, err := store.Get("hello"); err != nil || val != "world" {
		t.Errorf("Get() failed")
	}
	
	if val, err := store.Get("missing"); err == nil  || val != "" {
		t.Errorf("Get() failed")
	}
	
}


// Table driven test
func TestDelete_RemovesKey(t *testing.T) {
	tests := []struct {
		name string
		setup map[string]string
		deleteKey string
		wantLen int
	} {
		{"deletes existing key", map[string]string{"a": "1", "b": "2"}, "a", 1},
		{"no-op on missing key", map[string]string{"a": "1"}, "x", 1},
		{"no-op on empty store", map[string]string{}, "x", 0},
		{"deletes last remaining key", map[string]string{"only": "value"}, "only", 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			store := NewStore(5)
			for k, v := range tc.setup {
				store.Set(k, v)
			}

			store.Delete(tc.deleteKey)

			if got := store.Len(); got != tc.wantLen {
				t.Errorf("after Delete(%q): len() = %d, want %d", tc.deleteKey, got, tc.wantLen)
			}

			if _, err := store.Get(tc.deleteKey); err == nil {
				t.Errorf("after Delete(%q): key still present", tc.deleteKey)
			}
		})
	}
}