package main

import (
	"demo/messaging"
	"time"
)

func main() {
	ch := messaging.Generate(time.Second*5, time.Millisecond*100)
	<-messaging.Receive(ch)
	println("End of main")
}
