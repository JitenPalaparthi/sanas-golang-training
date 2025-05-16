package messaging

import (
	"fmt"
	"runtime"
	"sync"
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
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		<-timeOut
		wg.Done()
	}()

	go func() {
		a, b := 0, 1
		i := 1
		done := false
		go func() {
			wg.Wait()
			done = true
			runtime.Goexit() // what it exit
		}()

		for {
			if !done {
				ch <- NewMessage(i, fmt.Sprint("some fibonachi nubmer-->", a))
				a, b = b, a+b
			} else {
				close(ch)
				break
			}
			i++
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
