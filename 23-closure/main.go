package main

import "math"

func main() {

	s := calc(10, 20, func(i1, i2 int) int {
		return i1 + i2
	})
	println("sum:", s)

	s = calc(10, 20, sub)
	println("substract:", s)

	fn1 := func(a, b int) int {
		return a * b
	}

	s = calc(10, 20, fn1)

	println("multiplication:", s)

	var Fn func(int, int) int

	Fn = func(i1, i2 int) int {
		return i1 / i2
	}

	s = calc(10, 20, Fn)
	println("division:", s)
	fn2 := SetVal(10)
	v := fn2()
	println("function return:", v)

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

func SetVal(n int) func() int {
	return func() int {
		return n * n
	}
}

func MakeAFunc(n int, calltype string) any {
	switch calltype {
	case "sq", "square", "SQ", "SQUARE":
		return func() int {
			return n * n
		}
	case "cube", "CUBE", "Cube":
		return func() int {
			return n * n * n
		}

	case "sqroot":
		return func() int {
			return int(math.Sqrt(float64(n)))
		}
	case "modivideby2":
		return func() (int, int) {
			return n / 2, n % 2
		}
	default:
		return nil
	}
}
