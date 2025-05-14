package main

func main() {
	funcslice := make([]func() int, 5)
	funcslice1 := make([]any, 5) // func() int

	for i := 0; i < len(funcslice); i++ {
		funcslice[i] = func() int {
			return i * i
		}
	}

	for _, f := range funcslice {
		v := f()
		println(v)
	}
	// i := 0
	// for i < len(funcslice1) {
	// 	funcslice1[i] = func(i int) func() int {
	// 		return func() int {
	// 			return i * i
	// 		}
	// 	}(i)
	// 	i++
	// }
	i := 0
loop:
	funcslice1[i] = func(i int) func() int {
		return func() int {
			return i * i
		}
	}(i)
	i++
	if i < len(funcslice1) {
		goto loop
	}

	for _, f := range funcslice {
		v := f()
		println(v)
	}

	for _, f := range funcslice1 {
		switch f := f.(type) {
		case func() int:
			v := f()
			println(v)
		}
	}
}
