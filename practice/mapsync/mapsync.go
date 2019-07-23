package main

import (
	"sync"
	"time"
)

type MutexMap struct {
	sync.RWMutex
	Map map[int]int
}

// MutexMapTest() 針對Map併發讀寫拋錯，進行RWMutex讀寫鎖封裝練習。
func MutexMapTest() {
	mm := new(MutexMap)
	mm.Map = make(map[int]int)

	go func() { //read
		for {
			mm.RLock()
			_ = mm.Map[1]
			mm.RUnlock()

		}
	}()

	go func() { //write
		for {
			mm.Lock()
			mm.Map[2] = 2
			mm.Unlock()
		}
	}()

	for {
		time.Sleep(100 * time.Millisecond)

	}
}
