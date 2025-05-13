package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {

	var arr1 [5]int // zero values

	// what are value of the array -->  [0 0 0 0 0]
	// what is the type of the array --> [5]int

	fmt.Println("Values of array:", arr1, "Type:", reflect.TypeOf(arr1))

	for i := range arr1 {
		arr1[i] = rand.Intn(100) // mutating within the bounds
	}

	fmt.Println(arr1)

	arr2 := [5]int{12, 12, 4, 16, 34}
	arr3 := [...]int{12, 12, 4, 16, 34, 23, 32} // [...]tells that it is an array
	fmt.Println("Values of array:", arr2, "Type:", reflect.TypeOf(arr2))
	fmt.Println("Values of array:", arr3, "Type:", reflect.TypeOf(arr3))

	s := SumOf(arr1)
	println("SumOf:", s)
	s = SumOf(arr2)
	println("SumOf:", s)

	s = SumOf7(arr3)
	println("SumOf:", s)
	// cannot type caste one arry type to another array type

	//var arr4 [5]int = [5]int(arr3) // cannot type cast arrays of different types(essentially lengths)

	var arr4 [5]int

	for i := 0; i < min(len(arr4), len(arr3)); i++ {
		arr4[i] = arr3[i]
	}

	arr2d := [2][2]int{{1, 2}, {3, 4}}
	arr3d := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}

	println("2d array")
	for _, a1 := range arr2d {
		for _, v := range a1 {
			println(v)
		}
	}

	println("3d array")
	for _, a1 := range arr3d {
		for i := 0; i < len(a1); i++ {
			for j := 0; j < len(a1[i]); j++ {
				println(a1[i][j])
			}
		}
	}

	// var any1 any = arr2d

	fmt.Println(reflect.TypeOf(arr2d).Kind().String())

}

func SumOf(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumOf7(arr [7]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
