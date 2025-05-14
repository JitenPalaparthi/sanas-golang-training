package main

import (
	"fmt"
	"unsafe"
)

var Greet1 string = "Hello Sanas!"
var Greet2 string = "Hello Sanas---India!"

func main() {

	p1 := Person{Name: "Jiten", Email: "JitenP@Outlook.com", Address: Address{"560086", "India", "Active"}}
	println(p1.Address.Country)
	e1 := Employee{Name: "Jiten", Email: "JitenP@Outlook.com", Address: Address{"560086", "India", "Active"}}
	println(e1.Country)
	println(e1.Address.Status)

	s1 := Student{
		Name: "Jiten", Email: "JitenP@Outlook.com",
		Address: Address{"560086", "India", "Active"},
		SocialMedia: struct {
			Accounts []string
			Emails   []string
		}{[]string{"linlkedin.com/jpalaparthi"}, []string{"jiten.palaparthi@gmail.com", "jitenp@outlook.com"}},
	}

	// var company struct { // anonymous structure
	// 	Id          int
	// 	Name, Email string
	// }

	fmt.Println(s1)
	c1 := struct { // anonymous structure
		Id          int
		Name, Email string
	}{101, "Sanas", "sanas@sanas.com"}

	// map1 := map[string]any{"560086": 1, "560096": "Blr-2", "fn1": func() {
	// 	println("hello World")
	// }}
	// slice1 := []int{10, 20, 30}

	// num := 100
	// num1 := uint8(100)

	fmt.Println(c1)

	var empty1 Empty
	fmt.Printf("Size of empty1:%d Address:%p\n", unsafe.Sizeof(empty1), &empty1)

	var empty2 Empty
	fmt.Printf("Size of empty2:%d Address:%p\n", unsafe.Sizeof(empty2), &empty2)

	var any1 Any
	fmt.Printf("Size of any1:%d any1:%p\n", unsafe.Sizeof(any1), &any1)

	fmt.Printf("Address of c1:%p\n", &c1)
	fmt.Printf("Address of Greet:%p\n", &Greet1)
	fmt.Printf("Address of Greet:%p\n", &Greet2)
}

type Person struct {
	Name, Email string
	Address     Address
}

type Address struct {
	PinCode, Country, Status string
}

type Employee struct {
	Name, Email, Status string
	Address             // Promoted filed
}

type Student struct {
	Name, Email string
	SocialMedia struct {
		Accounts []string
		Emails   []string
	}
	Address // promoted field
}

type Empty struct{}

type Any struct{}

// 10016d190
// 102b55190

// 10016d190 D main.Greet1
// 10016d1a0 D main.Greet2

// Address of Greet:0x102655190
// Address of Greet:0x1026551a0
