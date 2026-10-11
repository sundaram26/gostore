package main

import (
	"github.com/sundaram26/gostore/store"
	"github.com/sundaram26/gostore/store/kv"

	"encoding/base64"
	"fmt"
)

func main() {
	// original := kv.NewStore(20)
	// logger := NewLoggingMiddleware(original)
	// metric := NewMetricMiddleware(logger)
	
	// newStore := original.Clone()
	// newStore.Set("env", "production")

	

	// encrypted, err := SetKeyWithEncryption(metric, "a", "hey you can see me")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(encrypted)
	// metric.Report()

	// fmt.Println(original)
	// fmt.Println(newStore)

	cmds := []Command{
		{Op: "SET", Key: "env", Value: "production"},
		{Op: "SET", Key: "version", Value: "0.0.1"},
		{Op: "SET", Key: "debug", Value: "true"},
		{Op: "GET", Key: "env"},
		{Op: "SET", Key: "region", Value: "eu-west-1"},
		{Op: "GET", Key: "version"},
	}

	s := kv.NewStore(0)
	RestoreOnBoot(s, cmds)

	fmt.Println("restored, store size: ", s.Len())
	fmt.Println("keys: ", s.Keys())
}

// func CreateStore() store.Storer {
// 	plain := kv.NewStore(20)
// 	logger := NewLoggingMiddleware(plain)
// 	metrics := NewMetricMiddleware(logger)

// 	return metrics
// }

func SetKeyWithEncryption(store store.Storer, key, val string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(val))
	if err := store.Set(key, encoded); err != nil {
		return "", err
	}

	store.Len()
	store.Get("key")

	return store.Get(key)
}
