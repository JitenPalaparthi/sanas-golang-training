package shape

type Square float32

// func NewSquare(s float32) Square {
// 	var sq Square = Square(s)
// 	return sq
// }

func (s *Square) Area() float32 {
	if s != nil {
		return float32(*s * *s)
	}
	return 0
}

func (r Square) Perimeter() float32 {
	return 4 * float32(r)
}

func (r Square) What() string {
	return "Square"
}
