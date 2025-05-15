package shape

var (
	Global1 struct{} // exportable , unrestricted
	global2 struct{} // exportable , unrestricted
)

func Greet1() { // exportable , unrestricted
	greet2()
}

func greet2() { // unexportable , restricted
	println("Hello World!")
}

type T1 struct { // exportable , unrestricted
	A1, B1, c1, d1 struct{}
}

func (t *T1) SayHi() {
	greet2()
}

func (t *T1) sayHi() {
	greet2()
}

type t2 struct { // exportable , unrestricted
	A1, B1, c1, d1 struct{}
}

func (t *t2) SayHi() {
	greet2()
}

func (t *t2) sayHi() {
	greet2()
}

type D1 struct{}
type D2 struct{}

func (D1) Greet() {
	println("Hello World")
}

func (D2) Greet(msg string) {
	println(msg)
}
