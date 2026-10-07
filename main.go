package main

import (
	"fmt"
	"time"
)

func main() {
	kvStore := NewStore(20)

	encrypted, err := kvStore.SetKeyWithEncryption("a", "hey you can see me")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(encrypted)

	ttlKVStore := NewTTLStore()
	encryptedTTL, err := ttlKVStore.SetKeyWithEncryption("a", "Hey you can read me", 10*time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(encryptedTTL)
}
