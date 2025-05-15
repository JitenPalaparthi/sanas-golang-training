package main

import (
	"fmt"
	"unsafe"
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

	var any1 interface{} = 10
	any1 = "Hello Sanas"

	var empty1 IEmpty = 10
	empty1 = "Hello World"
	empty1 = r1
	empty1 = s1
	empty1 = c1
	fmt.Println(any1, empty1)
}

type IEmpty interface{}

func Shape(iShape IShape) {
	fmt.Println("Size of iShape:", unsafe.Sizeof(iShape))
	switch iShape.(type) {
	case *Rect:
		fmt.Printf("Area of Rect:%.3f\n", iShape.Area())
		fmt.Printf("Perimeter of Rect:%.3f\n", iShape.Perimeter())
	case *Square:
		fmt.Printf("Area of Square:%.3f\n", iShape.Area())
		fmt.Printf("Perimeter of Square:%.3f\n", iShape.Perimeter())
	case *Cuboid:
		fmt.Printf("Area of Cuboid:%.3f\n", iShape.Area())
		fmt.Printf("Perimeter of Cuboid:%.3f\n", iShape.Perimeter())
	}
}
