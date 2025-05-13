package main

import (
	_ "os"
)

func main() {

	a, b := 10, 20

	c := a + b/a - b*4 + a + (a + 20) - (b-10)*2

	println(c)

	d := a+b/a-b*4+a+(a+20)-(b-10)*2 > 100
	println(d)

	_ = d // blank identifier
}
