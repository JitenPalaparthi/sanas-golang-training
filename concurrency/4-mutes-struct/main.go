package main

import (
	"sync"
)

var (
	wg *sync.WaitGroup = new(sync.WaitGroup)
)

type Counter struct {
	Count int
	mu    *sync.Mutex
	//mur   *sync.RWMutex
}

func NewCounter() *Counter {
	return &Counter{0, new(sync.Mutex)}
}

func (c *Counter) Decrement() {
	for i := 1; i <= 1000; i++ {
		c.mu.Lock()
		c.Count--
		c.mu.Unlock()
	}
}

func (c *Counter) Incrment() {
	for i := 1; i <= 1000; i++ {
		c.mu.Lock()
		c.Count++
		c.mu.Unlock()
	}
}

func main() {

	c := NewCounter()

	wg.Add(2)
	go func() {
		c.Incrment()
		wg.Done()
	}()
	go func() {
		c.Decrement()
		wg.Done()
	}()

	wg.Wait()
	println("Counter:", c.Count)

}
