package messaging

import (
	"fmt"
	"runtime"
	"time"
)

// var Ch chan int

// func init() {
// 	Ch = make(chan int)
// }

type Message struct {
	Code    int
	Message string
}

func NewMessage(code int, msg string) Message {
	return Message{code, msg}
}

func Generate(d time.Duration, sleep time.Duration) <-chan Message {
	ch := make(chan Message)
	timeOut := After(d)
	//time.Sleep(time.Second * 1)
	go func() {
		a, b := 0, 1
		i := 1
		for {
			select {
			case ch <- NewMessage(i, fmt.Sprint("some fibonachi nubmer-->", a)):
				a, b = b, a+b
				i++
			case <-timeOut:
				close(ch)
				runtime.Goexit()

				// case <-func(d time.Duration) chan struct{} {
				// 	ch := make(chan struct{})
				// 	go func() {
				// 		time.Sleep(d) // after d duration
				// 		ch <- struct{}{}
				// 		close(ch)
				// 	}()
				// 	return ch
				// }(d):
			}
		}
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

func After(d time.Duration) chan struct{} {
	ch := make(chan struct{})
	go func() {
		time.Sleep(d) // after d duration
		ch <- struct{}{}
		close(ch)
	}()
	return ch
}
