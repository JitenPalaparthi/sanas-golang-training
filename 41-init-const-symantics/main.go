package main

import (
	//_ "demo/greet"
	"demo/salutation"
	"fmt"
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

	fmt.Println(salutation.T1)

}
