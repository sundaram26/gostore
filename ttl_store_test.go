package main

import (
	"testing"
	"time"
)

func TestTTLStore_GetBeforeExpiry(t *testing.T) {
	s := NewTTLStore()
	_ = s.Set("k", "v", 1*time.Second)

	got, err := s.GetWithTTL("k")
	if err != nil || got != "v" {
		t.Fatalf("expected v, got %q %v", got, err)
	}
}

func TestTTLStore_ExpiredKeyEvicted(t *testing.T) {
	s := NewTTLStore()
	_ = s.Set("k", "v", 50*time.Millisecond)
}