package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num1 uint8 = 123
	var num2 uint16 = 12312
	var num3 float32 = 1231.2
	var num4 float64 = 123232.23
	var any1 any = 12312
	var any2 any = 12312.123
	var ok1 bool = true
	var str1 string = "12312"

	var sum float64

	sum = float64(num1) + float64(num2) + float64(num3) + num4 + float64(any1.(int)) + any2.(float64)

	if ok1 {
		sum += 1
	}

	v1, _ := strconv.Atoi(str1)
	sum += float64(v1)

	fmt.Printf("Sum of:%.3f", sum)

}
