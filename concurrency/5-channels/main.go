package main

import "time"

func main() {
	var ch chan int // declaration of the channel but not instantiation
	if ch == nil {
		println("Yes nil channel")
	}
	ch = make(chan int) // unbuffered channel
	go func() {
		v := <-ch
		time.Sleep(time.Second * 5)
		println(v)
	}()
	// it is not executing the next step
	ch <- 100 // sender
	time.Sleep(time.Second * 6)
}

// Do not communicate by sharing memory; instead, share memory by communicating.
// channel is a queue.. the sender sends the data and receiver receives
// the sender and the receiver is blocked if any of the them is blocked.

// until the receiver receives the data , the sender is blocked
// until the sender sends the data, the revceiver is bloced
// order of the sender or receiver go routines does not matter
