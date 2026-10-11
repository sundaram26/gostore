package main

import (
	"fmt"
	"time"

	"github.com/sundaram26/gostore/store"
)

// Metrics Middleware tracks usage statistics for the store.
// Like LoggingMiddleware , it wraps a Storer(the decorator pattern) and records how many time each operation get called
//
// TODO: declare the MetricsMiddleware struct below.

type MetricsMiddleware struct {
	inner store.Storer 
	
	getCalls int
	setCalls int
	deleteCalls int
	lenCalls int 

	getMisses int

	totalGetLatency time.Duration
	totalSetLatency time.Duration
}

func NewMetricMiddleware(inner store.Storer) *MetricsMiddleware {
	return &MetricsMiddleware{inner: inner}
}

func (m *MetricsMiddleware) Get(key string) (string, error) {
	start := time.Now()
	val, err := m.inner.Get(key)
	if err != nil {
		m.getMisses++
	}
	m.getCalls++
	m.totalGetLatency += time.Since(start)
	return val, err
}

func (m *MetricsMiddleware) Set(key string, value string) error {
	start := time.Now()
	err := m.inner.Set(key, value)
	m.setCalls++
	m.totalSetLatency += time.Since(start)
	return err
}

func (m *MetricsMiddleware) Delete(key string) {
	m.inner.Delete(key)
	m.deleteCalls++
}

func (m *MetricsMiddleware) Len() int {
	len := m.inner.Len()
	m.lenCalls++
	return len
}

func (m *MetricsMiddleware) Keys() []string {
	got := m.inner.Keys()
	return got
}

func (m *MetricsMiddleware) Report() {
	fmt.Printf("GET calls: %d. (missed %d)\n", m.getCalls, m.getMisses)
	fmt.Printf("SET calls: %d\n", m.setCalls)
	fmt.Printf("DELETE calls: %d\n", m.deleteCalls)
	fmt.Printf("LEN calls: %d\n", m.lenCalls)

	if m.getCalls > 0 {
		avg := m.totalGetLatency / time.Duration(m.getCalls)
		fmt.Printf("Avg GET latency: %v\n", avg)
	}
	
	if m.setCalls > 0 {
		avg := m.totalSetLatency / time.Duration(m.setCalls)
		fmt.Printf("Avg SET latency: %v\n", avg)
	}
}