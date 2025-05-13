package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type small = uint16

func main() {
	var b1 byte = 69
	var c1 rune = 'A'
	var s1 small
	fmt.Println("Type of s1: ", reflect.TypeOf(s1))
	fmt.Printf("type of s1 %T\n", s1)
	fmt.Printf("type of b1 %T\n", b1)
	fmt.Printf("type of c1 %T\n", c1)

	str1 := "Hello World!"
	str2 := "Hello Universe!"
	var str3 string

	fmt.Println("Size of str1:", unsafe.Sizeof(str1))
	fmt.Println("Size of str2:", unsafe.Sizeof(str2))
	fmt.Println("Size of str3:", unsafe.Sizeof(str3))

	var any1 any

	fmt.Printf("Type of any1:%T Data of any1:%v\n", any1, any1)

	any1 = str1
	fmt.Printf("Type of any1:%T Data of any1:%v\n", any1, any1)

	any1 = 100

	fmt.Printf("Type of any1:%T Data of any1:%v\n", any1, any1)

	any1 = true

	fmt.Printf("Type of any1:%T Data of any1:%v\n", any1, any1)

	any1 = 123123.123

	fmt.Printf("Type of any1:%T Data of any1:%v\n", any1, any1)

	//any1 = Person("Jiten", 40)

}
