package main

import "fmt"

func main() {

	slice1 := make([]int, 10)
	slice1 = append(slice1, 10, 20, 30, 40, 50, 60)
	slice2 := []int{}
	slice2 = append(slice2, slice1...)
	//fmt.Println("a", "b", "Hello", true, 100)
	fmt.Println("Sumof", SumOf())
	fmt.Println("Sumof", SumOf(10, 20))
	fmt.Println("Sumof", SumOf(10, 20, 123, 123, 34, 35))
	fmt.Println("Sumof", SumOf(slice1...))

	arr1 := [5]int{10, 20, 30, 40, 50}
	//slice3 = arr1[:] // converting an array to a slice
	fmt.Println("Sumof", SumOf(arr1[:]...))
}

// variadic can be used only for functions/methods , or a definition to an interdfe
// variadic parameter must be the last parameter in a func
// func SumOf(nums ...int, multiplier int) int {

// }

func SumOf(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return 0
}
