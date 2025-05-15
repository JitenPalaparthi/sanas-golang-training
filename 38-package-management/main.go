package main

import (
	"fmt"
	. "math"
	"math/rand"
	"shapes/shape"
	. "time"
	tm "time"
)

func main() {

	r1 := shape.NewRect(12.3, 12)
	s1 := shape.Square(123.23)

	r2 := shape.NewRect(12.3, 12)
	s2 := shape.Square(123.23)

	c1 := shape.NewCuboid(12.23, 23.4, 23.2)

	ishapes := []shape.IShape{r1, r2, &s1, &s2, c1}
	ShapeAll(ishapes)

	println("Runtime")

	for i := 1; i <= 10; i++ {
		r := rand.Intn(len(ishapes))
		Shape(ishapes[r])
	}

	fmt.Println(tm.Now().UTC())

	shape.D1{}.Greet()
	shape.D2{}.Greet("Hey Sanas Minds!")

	r1.Data = r1
	r1.GetUUID()
	c1.Data = c1
	c1.GetUUID()

	fmt.Println(r1.ToString())
	fmt.Println(c1.ToString())
	println(Abs(123.23))
	println(Now().UTC())
}

func Shape(iShape shape.IShape) {
	fmt.Printf("Area of %s :%.3f\n", iShape.What(), iShape.Area())
	fmt.Printf("Perimeter of %s:%.3f\n", iShape.What(), iShape.Perimeter())
}

func ShapeAll(ishapes []shape.IShape) {
	for _, s := range ishapes {
		fmt.Printf("Area of %s :%.3f\n", s.What(), s.Area())
		fmt.Printf("Perimeter of %s:%.3f\n", s.What(), s.Perimeter())
	}
}

// func real() {
// 	fmt.Println("somethign real")
// }
