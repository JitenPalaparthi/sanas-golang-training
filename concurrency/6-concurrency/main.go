package main

import (
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wg := new(sync.WaitGroup)
	defer wg.Wait()
	wg.Add(1)
	go func() {
		Generate(10, time.Millisecond*100, ch)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		Receive(10, ch)
		wg.Done()
	}()
}

func Generate(n uint, d time.Duration, ch chan<- int) {
	a, b := 0, 1
	for i := 1; i <= int(n); i++ {
		time.Sleep(d)
		ch <- a
		a, b = b, a+b
	}
}

func Receive(n int, ch <-chan int) {
	for i := 0; i < n; i++ {
		println(<-ch)
	}
}
