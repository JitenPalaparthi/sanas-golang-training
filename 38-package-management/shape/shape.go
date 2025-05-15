package shape

import (
	"fmt"

	"github.com/google/uuid"
)

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

type Common struct {
	Data any
}

func (c *Common) GetUUID() string {
	return uuid.New().String()
}
func (c *Common) ToString() string {
	return fmt.Sprint(c.Data)
}
