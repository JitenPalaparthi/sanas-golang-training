package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg *sync.WaitGroup = new(sync.WaitGroup)
)

func main() {
	ch1 := Generator(100, "Gen-1")
	sig := make(chan struct{})
	go Receiver("recv-1", ch1, sig)
	go Receiver("recv-2", ch1, sig)
	go Receiver("recv-3", ch1, sig)
	wg.Add(1)
	go ReceiverWg("recv-wg-1", ch1, wg)
	wg.Add(1)
	go ReceiverWg("recv-wg-2", ch1, wg)
	wg.Add(1)
	go ReceiverWg("recv-wg-3", ch1, wg)
	<-sig
	<-sig
	<-sig
	wg.Wait()

}

func Generator(n uint, name string) chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < int(n); i++ {
			time.Sleep(time.Millisecond * 100)
			ch <- fmt.Sprint(name, "--->", i)
		}
		close(ch)
	}()
	return ch
}

func Receiver(name string, ch chan string, sig chan<- struct{}) {
	//sig := make(chan struct{})
	for c := range ch {
		println(name, "-->", c)
	}
	sig <- struct{}{}
}

func ReceiverWg(name string, ch chan string, wg *sync.WaitGroup) {
	//sig := make(chan struct{})
	for c := range ch {
		println(name, "-->", c)
	}
	wg.Done()
}
