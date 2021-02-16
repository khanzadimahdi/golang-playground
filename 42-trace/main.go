package main

import (
	"log"
	"time"
)

func main() {
	defer trace("slow operation")()

	//... lots of work ...
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

func trace(name string) func() {
	start := time.Now()

	log.Printf("start %s at %v", name, start)

	return func() {
		log.Printf("end %s at %v (takes: %s)", name, time.Now(), time.Since(start))
	}
}
