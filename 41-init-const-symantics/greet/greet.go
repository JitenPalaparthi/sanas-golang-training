package greet

// import _ "demo/salutation"

func init() {
	println("Calling Greet Package-1")
}

func init() {
	println("Calling Greet Package-2")
}

func init() {
	println("Calling Greet Package-2")
}

func Greet() {
	println("Hello world")
}

type T struct{}

func (T) Greet() {
	println("Hello world")
}
