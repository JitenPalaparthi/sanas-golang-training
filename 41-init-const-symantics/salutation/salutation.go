package salutation

import "demo/greet"

var T1 greet.T

var t1 greet.T

func init() {
	println("Calling Salutation Package-1")
}

func init() {
	println("Calling Salutation Package-2")
}

func init() {
	println("Calling Salutation Package-2")
}

func Greet() {
	// println("Hello world")
	greet.Greet()
}

// type T struct{}

// func (T) Greet() {
// 	println("Hello world")
// }
