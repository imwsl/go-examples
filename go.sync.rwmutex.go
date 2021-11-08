package main

import (
	"fmt"
	"sync"
)

var (
	x int64
	wg sync.WaitGroup
	lock sync.RWMutex
)

func read() {
	lock.RLock()
	fmt.Printf("read %d\n", x)
	lock.RUnlock()
	wg.Done()
}

func write() {
	lock.Lock()
	fmt.Println("write...\n")
	x = x + 1
	lock.Unlock()
	wg.Done()
}

func main() {
	for i := 0; i < 5; i ++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 10; i ++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()

	for {}
}
