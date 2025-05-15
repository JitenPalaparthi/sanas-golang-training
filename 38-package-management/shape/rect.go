package shape

type Rect struct {
	L, B float32
	Common
}

func NewRect(l, b float32) *Rect {
	return &Rect{l, b, Common{}}
}

func (r *Rect) Area() float32 {
	return (*r).L * r.B
}
