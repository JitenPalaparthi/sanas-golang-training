package main

import (
	"time"
)

func main() {
	ch := make(chan int)
	sig := make(chan struct{})

	go Generate(100, time.Millisecond*100, ch)
	go Receive(10, ch, sig)
	<-sig // Why blocked  // future
}

func Generate(n uint, d time.Duration, ch chan<- int) {
	//time.Sleep(time.Second * 1)
	a, b := 0, 1
	for i := 1; i <= int(n); i++ {
		//time.Sleep()
		ch <- a
		a, b = b, a+b
	}
}

func Receive(n int, ch <-chan int, sig chan<- struct{}) {
	for i := 0; i < n; i++ {
		v := <-ch
		println(v)
	}
	sig <- struct{}{}
}
