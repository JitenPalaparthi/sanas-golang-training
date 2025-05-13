package main

import "fmt"

const PI float32 = 3.14

var Global int = 9999

func main() {

	//PI = 123.12
	var num1 uint8 = 100
	var num2 int32 = 1231

	{
		num10 := 100
		println(num10)
	}

	var (
		num4,
		num5 uint64
		float1 float32
		float2 float64
		ok1    bool
		str1   string
	) // zero values

	{
		a, b, c := 10, 20, 30 // shorthand declartion, it is evaluated by the compier at compile time

		// t := a
		// a = b
		// b = t

		a, b, c = b, c, a // swap

	}

	var age = 40 // type inference, the type is inferred by the compiler
	var float3 = 3.14

	var ok2 = true

	a, b, c := 10, 20, 30 // shorthand declartion, it is evaluated by the compier at compile time

	// t := a
	// a = b
	// b = t

	a, b, c = b, c, a // swap

	ok3 := true

	str2 := "Hello World!"
	str2 = "Hello Universe"

	fmt.Println(num1, num2, num4, num5, float1, float2, ok1, str1, age, ok2, float3, ok3, str2)

	c1 := 123.34 + 123.2i

	c2 := complex(123.123, 1.3) // compelx128 or complex64

	var f1, im1 float32 = 123.124, 1.3
	c3 := complex(f1, im1)

	c4 := c1 + c2
	c5 := c1 - c2
	c6 := c1 * c2
	fmt.Println(c1, c2, c3, c4, c5, c6)

	char1 := 'A'

	char2 := 'B'

	var char3 rune = 68

	var byte1 uint8 = 'D'

	char4 := char1 + char2 + char3

	var char5 int32 = 65

	fmt.Println(char1 == char5)

	fmt.Println(char4, byte1)

	str4 := "Hello World!" // 12
	str5 := "Hello, 世界"    // 7 + 6 -->13

	fmt.Printf("length of str1: %d str2: %d\n", len(str4), len(str5))
}

// numbers : int8,int16,int32,int64,uint8,uint16,uint32,uint64,float32,float64,int,uint
// alias   : rune(int32), byte(uint8)
// complex : complex64, complex128
// boolean: bool
// strings --> string
// specialtype --> any / interface{}

// type inference
// zero value

// comptime --> compiler does many checks during compilation itslef

// unicode characters --> 1-4 of length
