package main

type IShape interface {
	IArea
	IPerimeter
	IWhat
}

type IArea interface {
	Area() float32
}

type IPerimeter interface {
	Perimeter() float32
}

type IWhat interface {
	What() string
}
