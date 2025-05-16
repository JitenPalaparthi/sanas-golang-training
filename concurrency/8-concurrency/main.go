package main

import (
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)
	sig := make(chan struct{})
	go Generate(20, time.Millisecond*100, ch)
	go Receive(ch, sig)
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
	close(ch) // once it is closed.. you canonot send data or open it
	// close would not make the channel nil
}

func Receive(ch <-chan int, sig chan<- struct{}) {
	for {
		v, ok := <-ch
		if ok {
			println(v)
		} else {
			sig <- struct{}{}
			runtime.Goexit()
		}
	}
}
