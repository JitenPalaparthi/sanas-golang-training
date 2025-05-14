package main

func main() {
	var f1 = func() {
		println("hello world")
	}

	if f1 == nil {
		println("it is nil")
	}
	f1()

	var f2 = func(a, b int) int {
		return a + b
	}(10, 20)
	println(f2)

	var fn1 func(int, int) int

	if fn1 == nil {
		println("Yes is nil")
	}

	fn1 = sub
	r := fn1(10, 20)
	println(r)

	fn1 = func(i1, i2 int) int {
		return i1 * i2
	}
	r = fn1(10, 20)
	println(r)

}

func sub(a, b int) int {
	return a - b
}
