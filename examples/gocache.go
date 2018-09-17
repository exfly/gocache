package main

import (
	"log"
	"time"

	"github.com/ExFly/gocache/memory"
)

func main() {
	cacher := memory.NewMemCache(10, 5*time.Second)
	err := cacher.Set("key", "val")
	if err != nil {
		log.Print(err)
	}
	val, err := cacher.Get("key")
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("get the val: %v", val)
	}
}
