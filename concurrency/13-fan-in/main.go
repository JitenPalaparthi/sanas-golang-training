package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := Generator(10, "Gen-1")
	ch2 := Generator(10, "Gen-2")
	ch3 := Generator(10, "Gen-3")
	<-Receive(ch1, ch2, ch3)
}

func Generator(n uint, name string) chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < int(n); i++ {
			ch <- fmt.Sprint(name, "--->", i)
		}
		close(ch)
	}()
	return ch
}

func Receive(chs ...chan string) chan struct{} {
	sig := make(chan struct{})
	wg := new(sync.WaitGroup)
	wg.Add(len(chs))
	go func() {
		for _, ch := range chs {
			go func() {
				for s := range ch {
					println(s)
				}
				wg.Done()
			}()
		}
	}()

	go func() {
		wg.Wait()
		sig <- struct{}{}
		close(sig)
	}()
	return sig
}
