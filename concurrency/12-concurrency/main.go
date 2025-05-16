package main

import (
	"demo/messaging"
	"time"
)

func main() {
	ch := messaging.Generate(time.Millisecond*500, time.Millisecond*100)
	<-messaging.Receive(ch)
	println("End of main")
}
