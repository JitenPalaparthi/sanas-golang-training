package main

import (
	"fmt"
	"reflect"
)

func main() {

	funcslice1 := make([]any, 5) // func() int

	fn1 := func() {
		fmt.Println("hello Wrold")
	}

	fn2 := func(a, b int) int {
		return a + b
	}

	fn3 := func(a int) int {
		return a * a
	}

	fn4 := calc(10, 20, sub)
	fmt.Println(reflect.TypeOf(fn4))

	fn5 := sub
	funcslice1[0], funcslice1[1], funcslice1[2], funcslice1[3], funcslice1[4] = fn1, fn2, fn3, fn4, fn5

	for _, v := range funcslice1 {
		switch v1 := v.(type) {
		case func():
			// v.(func())()
			v1()
		case func(int, int) int:
			fmt.Println(v.(func(int, int) int)(10, 20))
		case func(int) int:
			fmt.Println(v1(100))
		case int:
			println(v1)
		}

	}

}

func calc(a, b int, fn func(int, int) int) int {
	if fn != nil {
		return fn(a, b)
	}
	return -1
}

func sub(a, b int) int {
	return a - b
}
