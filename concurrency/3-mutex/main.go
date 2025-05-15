package main

import (
	"sync"
)

var (
	wg      *sync.WaitGroup
	mu      *sync.Mutex
	Counter int // do not use shared variable
)

func init() {
	wg = new(sync.WaitGroup)
	mu = new(sync.Mutex)
}

func main() {
	wg.Add(1)
	go func() {
		for i := 1; i <= 1000; i++ {
			mu.Lock()
			Counter++
			mu.Unlock()
		}
		wg.Done()
	}()
	wg.Add(1)
	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		Decrement(mu)
		wg.Done()
	}(wg, mu)

	wg.Wait()
	println(Counter)
}

func Decrement(mu *sync.Mutex) {
	for i := 1; i <= 1000; i++ {
		mu.Lock()
		Counter--
		mu.Unlock()
	}
}
