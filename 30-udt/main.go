package main

func main() {
	num := 10
	var f1 MyFunc = func() int {
		return num * num
	}

	v := f1()
	println(v)
	f1.Greet().Greet()

}

type MyFunc func() int

func (m *MyFunc) Greet() *MyFunc {
	println("Hello MyFunction")
	return m
}
