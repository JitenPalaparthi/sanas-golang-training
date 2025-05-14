package main

import "fmt"

func main() {
	var p1 Person
	p1.Id = 101
	p1.Name = "Jiten"
	p1.Email = "JItenP@Outlook.com"
	var p2 Person = Person{Id: 102, Name: "Ravi", Email: "Ravi@Gmail.com"}
	p3 := Person{Id: 102, Name: "Ravi", Email: "Ravi@Gmail.com"}
	p4 := &Person{Id: 102, Name: "Ravi", Email: "Ravi@Gmail.com"}
	p5 := new(Person)
	p5.Id = 101
	p5.Name = "Jiten"
	p5.Email = "JItenP@Outlook.com"

	// var t1 T1

	fmt.Println(p1, p2, p3, p4, p5)
	// fmt.Println("Size of t1:", unsafe.Sizeof(t1))

	// var t2 T2

	// fmt.Println("Size of t2:", unsafe.Sizeof(t2))

	c1 := ColourCode{
		string:   "Red",
		int:      10011,
		mystring: "Some Red",
	}
	println(c1.int, c1.string, c1.mystring)
	c2 := ColourCode{int: 12312}
	println(c2.int, c2.string)

	c3 := ColourCode{10012, "Red", "Another Read"}
	println(c3)
}

type mystring = string // alias

type Person struct {
	Id    int
	Name  string
	Email string
}

type ColourCode struct { // anonymous fields
	int
	string
	mystring
}

// type T1 struct {
// 	b1 byte // 1
// 	Id int  // 8
// 	b2 byte // 1
// }

// type T2 struct {
// 	Id  int  // 8
// 	b1  byte // 1
// 	b2  byte // 1
// 	i16 uint16
// 	f32 float32
// }
