package main

import (
	"math/rand"
)

const DIV = 0

const MIN = 60
const SEC_DAY = 24*MIN + 10/MIN
const OK = true
const GREET = "Hello World"

// const ARR [5]int = {1,2,34,4,5}
var Rand = rand.Int()

func init() {
	println(Rand)
}

func init() {
	println("Calling init in main-2", SEC_DAY)
}

func init() {
	println("Calling init in main-3")
}

func init() {
	println("Calling init in main-4")
}

func init() {
	println(Rand)
}

func main() { // func-1
	num, div := 10, 0
	println("Start of main")
	func() { // func-1.1
		println("Start of func-1")
		var ptr *int
		*ptr = 1 // panic
		func() {
			println("Start of func-2")
			func() {
				println("Start of func-3")
				//_ = num / DIV
				_ = num / div // the systems panics
				println("Start of func-3")
			}()
			println("End of func-2")
		}()
		println("End of func-1")
	}()
	println("end of main")
}

// panic, defer , recover

// error
// panic
// fatal
