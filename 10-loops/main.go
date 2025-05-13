package main

import "math/rand/v2"

func main() {

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			print(i, " ")
		}
	}
	println()

	i := 1
	for i <= 10 { // looks like while loop syntax
		if i%2 == 0 {
			i++
			continue
		}
		print(i, " ")
		i++
	}
	println()

	i = 1
	for {
		if i > 10 {
			break
		}
		if i%2 == 0 {
			i++
			continue
		}
		print(i, " ")
		i++
	}
	println()

	for i := 1; ; i++ {
		if i > 10 {
			break
		}
		if i%2 == 0 {
			continue
		}
		print(i, " ")
	}
	println()

	i = 1
	for ; ; i++ {
		if i > 10 {
			break
		}
		if i%2 == 0 {
			continue
		}
		print(i, " ")
	}
	println()
	for i, j := 1, 10; i < j; i, j = i+1, j-1 {
		println("i:", i, "j:", j)
	}

	println("break and break outer kind of")

	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if i == 3 {
				break
			}
			println("i:", i, "j:", j)
		}
	}
	println("done-->break and break outer kind of ")

	done := false
	for i := 1; i <= 5; i++ {
		if done {
			break
		}
		for j := 1; j <= 5; j++ {
			if i == 3 {
				done = true
				break
			}
			println("i:", i, "j:", j)
		}
	}

	println("lable-->break and break outer kind of ")

outer:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if i == 3 {
				break outer
			}
			println("i:", i, "j:", j)
		}
	}

outer2:
	for {
		num := rand.IntN(100)
		switch {
		case num == 50:
			break outer2
		case num%2 == 0:
			println("Even number:", num)
		case num%2 != 0:
			println("Odd number:", num)
		}
	}
}

// break can be used in two areas 1.for 2.switch(but we dont use it unless ...)
