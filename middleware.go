package main

import (
	"log"
	"os"

	"github.com/sundaram26/gostore/store"
)

type LoggingMiddleware struct {
	inner store.Storer // any Storer - the thing we're decorating
	logger *log.Logger
}


func NewLoggingMiddleware(inner store.Storer) *LoggingMiddleware {
	return &LoggingMiddleware{
		inner: inner,
		logger: log.New(os.Stdout, "[log] ", log.Ltime/log.Lmicroseconds),
	}
}

func (l *LoggingMiddleware) Get(key string) (string, error) {
	l.logger.Printf("GET %q", key)
	val, err := l.inner.Get(key) // delegate to inner store

	if err != nil {
		l.logger.Printf("GET %q -> miss (%v)", key, err)
	} else {
		l.logger.Printf("GET %q -> hit", key)
	}

	return val, nil	
}

func (l *LoggingMiddleware) Set(key, value string) error {
	l.logger.Printf("SET %q", key) // side effect
	return l.inner.Set(key, value)
}

func (l *LoggingMiddleware) Delete(key string) {
	l.logger.Printf("DELETE %q", key) 
	l.inner.Delete(key)
}

func (l *LoggingMiddleware) Len() int {
	got := l.inner.Len()
	l.logger.Printf("LENGTH %d", got)
	return got;
}

func (l *LoggingMiddleware) Keys() []string {
	keys := l.inner.Keys()
	l.logger.Printf("KEYS %v", keys)
	return keys;
}

