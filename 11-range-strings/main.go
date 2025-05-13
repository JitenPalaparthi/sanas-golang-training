package main

import "fmt"

func main() {

	str := "Hello World!" // string is collection of bytes.. Byte is nothing but uint8

	for i := 0; i < len(str); i++ {
		print(string(str[i]), " ")
	}
	println()

	str = "Hello, 世界"
	for i := 0; i < len(str); i++ {
		print(string(str[i]), " ")
	}
	println()

	// for range loop
	// arrays, strings, slices, maps , channels

	for i, c := range str {
		println("Index:", i, "-->", string(c))
	}

	for _, c := range str {
		println("-->", string(c))
	}

	c := 0
	for i, _ := range str {
		c++
		println("-->", i)
	}
	fmt.Println("NUmber of chars:", c)
}
