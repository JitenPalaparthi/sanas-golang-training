package main

import "fmt"

func main() {

	r1 := Rect{12.12, 14.565, 0, 0}
	a1 := Area(r1)
	fmt.Printf("Area of r1:%.3f\n", a1)

	//a2 := (&r1).Area()
	a2 := r1.Area()
	fmt.Printf("Area of r1:%.3f\n", a2)

	fmt.Printf("Area of r1:%.3f\n", r1.A)

}

type Rect struct {
	L, B float32
	A, P float32
}

func Area(r Rect) float32 {
	r.A = r.L * r.B
	return r.A
}
