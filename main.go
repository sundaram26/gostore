package main

import (

	"github.com/sundaram26/gostore/store"
	"github.com/sundaram26/gostore/store/kv"

	"encoding/base64"
	"fmt"
)

func main() {
	plain := kv.NewStore(20)
	logger := NewLoggingMiddleware(plain)
	metric := NewMetricMiddleware(logger)
	encrypted, err := SetKeyWithEncryption(metric, "a", "hey you can see me")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(encrypted)
	metric.Report()
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
