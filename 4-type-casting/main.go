package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var num1 uint8 = 123

	var num2 uint16 = uint16(num1) // type casting

	println(num2)
	var num3 uint64 = 12312312 //1011101111011110 11111000

	var num4 uint8 = uint8(num3) // there would be data loss
	println(num4)

	fmt.Printf("num3:%b\n", num3)

	var num5 uint8 = 0b11111000

	fmt.Printf("num5:%d\n", num5)

	var float1 float32 = 123123.123

	var num6 = int(float1)

	fmt.Printf("num6:%d\n", num6)

	//ok1 := true // 1 byte

	//var num7 uint8 = uint8(ok1)

	var char1 rune = 47000
	var char2 rune = 'A' // This is also integer

	var str1 string = string(char1) // what happen here

	str2 := string(char2)

	fmt.Println(str1)
	fmt.Println(char2)
	fmt.Println(str2)

	//str3:= "47000"

	str3 := fmt.Sprint(char1)
	fmt.Println("Data:", str3, "Type of:", reflect.TypeOf(str3))

	str4 := strconv.Itoa(int(char1))
	fmt.Println("Data:", str4, "Type of:", reflect.TypeOf(str4))

	str5 := "47000"

	v1, err := strconv.Atoi(str5)

	if err != nil {
		println(err.Error())
	} else {
		println("covnverted to int", v1)
	}

	a1, s1, m1, d1 := calc(12, 23)
	fmt.Println(a1, s1, m1, d1)
	_, _, _, d2 := calc(12, 6)
	fmt.Println(d2)

	str5 = "123123.123"
	//var v2 float64
	v2, err := strconv.ParseFloat(str5, 32)
	if err != nil {
		println(err.Error())
	} else {
		println("covnverted to flopat64", v2)
	}

	ok1, _ := strconv.ParseBool("Hello World")
	println("covnverted to bool", ok1)

	var any1 any // inteface

	any1 = 100 //
	//var num7 int = int(any1) // cant do type casting
	var num7 int = any1.(int)
	fmt.Println(num7)

	any1 = true

	// var str6 string = any1.(string)
	// println(str6)

	//var str6 string
	str6, ok := any1.(string)
	if ok {
		println(str6)
	} else {
		b1, ok := any1.(bool)
		if ok {
			println(b1)
		} else {
			println("unable to assert")
		}
	}

	println("Hello World")

}

func calc(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}

// ssa
