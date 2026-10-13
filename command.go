package main

import (
	"log"
	"sync"

	"github.com/sundaram26/gostore/store"
)

// Command is a parsed Redis-ish operation: opcode plus its arguments.
type Command struct {
	Op string
	Key string
	Value string
}

// dispatch is a tiny switch that turns a Command into a store call.
func dispatch(s store.Storer, c Command){
	switch c.Op {
		case "SET":
			log.Printf("SET key %s", c.Key)
			s.Set(c.Key, c.Value)
		case "GET":
			s.Get(c.Key)
		case "INCR":
			s.Incr(c.Key)
	}
}

// RestorationOnBoot replays a slice of commands against the store.
func RestoreOnBoot(s store.Storer, cmds []Command) {
	var wg sync.WaitGroup

	for _, c := range cmds {
		wg.Go(func(){
			dispatch(s, c)
		})

		// wg.Add(1)
		// go func() {
		// 	defer wg.Done()
		// 	dispatch(s, c)
		// }()
	}

	wg.Wait()
}