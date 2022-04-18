package main

import (
	"fmt"
	"time"
	"runtime"
)

type EventType string

const (
	one		EventType = "one"
	two				  = "two"		
	three			  = "three"
)

type EventTypeNumber uint8

const (
	oneNumber	EventTypeNumber = iota
	twoNumber
	threeNumber
)

type Event struct {
	ID			string
	CreatedAt	time.Time
	Type		EventType	
}

func main() {
	start := time.Now()
	var m runtime.MemStats
	var memstats []uint64

	for i := 0; i < 100000; i++ {
		if i % 1000 == 0 {
			runtime.ReadMemStats(&m)
			memstats = append(memstats, m.Alloc / 1024)
		}
		event := Event{ ID: "A", CreatedAt: time.Now().UTC(), Type: two }
		fmt.Println("Event", event)
	}

	fmt.Println("Duration", time.Now().Sub(start).Seconds())
	// runtime.ReadMemStats(&m)
	fmt.Printf("Mem stats %v", memstats)
}
