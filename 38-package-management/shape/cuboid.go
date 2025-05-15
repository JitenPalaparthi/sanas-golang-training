package shape

type Cuboid struct {
	L, B, H float32
	Common
}

func NewCuboid(l, b, h float32) *Cuboid {
	return &Cuboid{l, b, h, Common{}}
}

func (r *Cuboid) Area() float32 {
	return r.L * r.B * r.H
}

func (r *Cuboid) Perimeter() float32 {
	return 2 * (r.L + r.B + r.H)
}

func (_ *Cuboid) What() string { // can omit the id is you are not using it
	return "Cuboid"
}
