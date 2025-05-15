package main

func main() {

	// db.selct().filter().orderby().where()

	//ic.Add(123).Sub(123).Mul(123).Add(-123).Div(123).Get()

	result := NewCalc(100).Add(10).Sub(5).Mul(2).Div(4).Add(10).Get()
	println(result)
}

type ICalc interface {
	Add(int) ICalc
	Sub(int) ICalc
	Mul(int) ICalc
	Div(int) ICalc
	Get() int
}

type Calc struct {
	Data int
}

func NewCalc(d int) *Calc {
	return &Calc{d}
}

func (c *Calc) Add(d int) ICalc {
	c.Data += d
	return c
}

func (c *Calc) Sub(d int) ICalc {
	c.Data -= d
	return c
}

func (c *Calc) Mul(d int) ICalc {
	c.Data *= d
	return c
}

func (c *Calc) Div(d int) ICalc {
	c.Data /= d
	return c
}

func (c *Calc) Get() int {
	return c.Data
}
