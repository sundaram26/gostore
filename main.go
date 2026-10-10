package main

import (
	"github.com/sundaram26/gostore/store"
	"github.com/sundaram26/gostore/store/kv"
	"github.com/sundaram26/gostore/store/ttl"

	"encoding/base64"
	"fmt"
	"time"
)

func main() {
	kvStore := kv.NewStore(20)

	encrypted, err := SetKeyWithEncryption(kvStore, "a", "hey you can see me")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(encrypted)

	ttlKVStore := ttl.NewTTLStore(time.Second * 2)
	encryptedTTL, err := SetKeyWithEncryption(ttlKVStore, "a", "Hey you can read me")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(encryptedTTL)
}

func SetKeyWithEncryption(store store.Storer, key, val string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(val))
	if err := store.Set(key, encoded); err != nil {
		return "", err
	}

	return store.Get(key)
}
