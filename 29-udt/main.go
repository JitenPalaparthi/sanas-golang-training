package main

import "fmt"

type integer = int // alias

type myint1 int // newtype

type myint2 myint1

type myint3 myint2

func (i myint1) ToString() string {
	return fmt.Sprint(i)
}

func (i myint2) Sq() int {
	return int(i * i)
}

func (i myint3) Cube() int {
	return int(i * i * i)
}

func main() {

	var (
		i1 int    = 30
		i2 myint1 = 20
		i3 myint2 = 30
		i4 myint3 = 40
	)

	s1 := myint1(i1).ToString()
	sq1 := myint2(i1).Sq()
	cb1 := myint3(i1).Cube()

	fmt.Println("i1--> ToString:", s1, "Sq:", sq1, "Cube:", cb1)

	s2 := i2.ToString()
	sq2 := myint2(i1).Sq()
	cb2 := myint3(i1).Cube()

	fmt.Println("i2-> ToString:", s2, "Sq:", sq2, "Cube:", cb2)

	s3 := myint1(i3).ToString()
	sq3 := i3.Sq()
	cb3 := myint3(i3).Cube()

	fmt.Println("i3-> ToString:", s3, "Sq:", sq3, "Cube:", cb3)

	s4 := myint1(i4).ToString()
	sq4 := myint2(i4).Sq()
	cb4 := i4.Cube()

	fmt.Println("i4-> ToString:", s4, "Sq:", sq4, "Cube:", cb4)

	var float1 float64 = 12312.123

	s5 := myint1(float1).ToString()
	sq5 := myint2(float1).Sq()
	cb5 := myint3(float1).Cube()

	fmt.Println("float1-> ToString:", s5, "Sq:", sq5, "Cube:", cb5)

}
