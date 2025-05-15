package main

import (
	"fmt"
)

func main() {

	r1 := NewRect(12.3, 12)
	s1 := Square(123.23)

	r2 := NewRect(12.3, 12)
	s2 := Square(123.23)

	c1 := NewCuboid(12.23, 23.4, 23.2)

	Shape(r1)
	Shape(r2)
	Shape(&s1)
	Shape(&s2)
	Shape(c1)

	c2 := 123.23 + 12.0i // notation
	c3 := complex(123.23, 1.34)

	fmt.Println(real(c2))
	fmt.Println(imag(c3))

}

func Shape(iShape IShape) {
	fmt.Printf("Area of %s :%.3f\n", iShape.What(), iShape.Area())
	fmt.Printf("Perimeter of %s:%.3f\n", iShape.What(), iShape.Perimeter())
}

// func real() {
// 	fmt.Println("somethign real")
// }
