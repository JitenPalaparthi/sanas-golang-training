package main

import "fmt"

func main() {
	println("start of main")
	defer println("end of main")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
				for i := 1; i <= 10; i++ {
					print(i, " ") // panics
				}
			}
		}()
		func() {
			num := 100
			for i := 10; i >= 0; i-- {
				print(num/i, " ") // panics
			}

			// -- defer
			for i := 1; i <= 10; i++ {
				print(i, " ") // panics
			}
			println()
		}()
	}()
	func() {
		a, b := 0, 1
		for i := 1; i <= 20; i++ {
			println(a)
			a, b = b, a+b
		}
	}()
}
