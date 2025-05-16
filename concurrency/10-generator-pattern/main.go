package main

import (
	"fmt"
	"time"
)

type Message struct {
	Code    int
	Message string
}

func NewMessage(code int, msg string) Message {
	return Message{code, msg}
}

func main() {
	ch := Generate(5, time.Millisecond*1)
	// sig := Receive(ch)
	// <-sig
	<-Receive(ch)
	println("End of main")
}

func Generate(n uint, d time.Duration) chan Message {
	ch := make(chan Message)
	//time.Sleep(time.Second * 1)
	go func() {
		a, b := 0, 1
		for i := 1; i <= int(n); i++ {
			time.Sleep(d)
			ch <- NewMessage(i, fmt.Sprint("some fibonachi nubmer-->", a))
			a, b = b, a+b
		}
		close(ch) // once it is closed.. you canonot send data or open it
		// close would not make the channel nil
	}()
	return ch
}

func Receive(ch <-chan Message) chan struct{} {
	sig := make(chan struct{})
	go func() {
		for msg := range ch { // it reads the data untile the channel is closed
			fmt.Println(msg)
		}
		sig <- struct{}{}
		close(sig)
	}()
	return sig
}
