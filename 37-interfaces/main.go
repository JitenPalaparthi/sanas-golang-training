package main

import (
	"fmt"
	"math/rand"
)

func main() {

	r1 := NewRect(12.3, 12)
	s1 := Square(123.23)

	r2 := NewRect(12.3, 12)
	s2 := Square(123.23)

	c1 := NewCuboid(12.23, 23.4, 23.2)

	ishapes := []IShape{r1, r2, &s1, &s2, c1}
	ShapeAll(ishapes)

	println("Runtime")

	for i := 1; i <= 10; i++ {
		r := rand.Intn(len(ishapes))
		Shape(ishapes[r])
	}

}

func Shape(iShape IShape) {
	fmt.Printf("Area of %s :%.3f\n", iShape.What(), iShape.Area())
	fmt.Printf("Perimeter of %s:%.3f\n", iShape.What(), iShape.Perimeter())
}

func ShapeAll(ishapes []IShape) {
	for _, s := range ishapes {
		fmt.Printf("Area of %s :%.3f\n", s.What(), s.Area())
		fmt.Printf("Perimeter of %s:%.3f\n", s.What(), s.Perimeter())
	}
}

// func real() {
// 	fmt.Println("somethign real")
// }
