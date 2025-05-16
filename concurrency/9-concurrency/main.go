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
	ch := make(chan Message)
	sig := make(chan struct{})
	go Generate(5, time.Second*10, ch)
	go Receive(ch, sig)
	<-sig // Why blocked  // future

	println("End of main")
}

func Generate(n uint, d time.Duration, ch chan<- Message) {
	//time.Sleep(time.Second * 1)
	a, b := 0, 1
	for i := 1; i <= int(n); i++ {
		time.Sleep(d)
		ch <- NewMessage(i, fmt.Sprint("some fibonachi nubmer-->", a))
		a, b = b, a+b
	}
	close(ch) // once it is closed.. you canonot send data or open it
	// close would not make the channel nil
}

func Receive(ch <-chan Message, sig chan<- struct{}) {
	for msg := range ch { // it reads the data untile the channel is closed
		fmt.Println(msg)
	}
	sig <- struct{}{}
}
