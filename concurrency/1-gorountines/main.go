package main

import (
	"runtime"
	"time"
)

func main() {
	defer println("Hello World--->")
	go func() {
		println("Hello World")
		defer time.Sleep(time.Second * 5)
		// panic("Dummy Panic")
	}()
	println("Hello end of main")
	//time.Sleep(time.Millisecond * 5) // this is not a solution
	runtime.Goexit() // it it carsh
	// XX --> Crashes
}

// Main is also a goroutine
// No gorutine waits for other goroution to complete its execution
// There is no concept of parent and child go routines
// cannot guarantee the order of execution
// just use go keyword to run a block of code as a goroutine
