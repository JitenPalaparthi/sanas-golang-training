package main

import "fmt"

const MAX = 10

func main() {
	var slice1 []int
	PrintSliceHeader("slice1", slice1)
	// if slice1 == nil {
	// 	println("slice 1 is nil ")
	// }
	slice1 = make([]int, 5)
	fmt.Println("slice1 -------------------------")
	fmt.Printf("Address of slice:%p\n", &slice1)
	if slice1 != nil && len(slice1) > 0 {
		fmt.Println("Ptr in the slice:", &slice1[0])
	} else {
		fmt.Println("Ptr in the slice:nil")
	}
	fmt.Println("Len:", len(slice1))
	fmt.Println("Cap:", cap(slice1))
	fmt.Println(slice1)
	fmt.Println("-----------------------")

	PrintSliceHeader("slice1", slice1) // call by value

	// slice2 := make([]int, 5, 10)
	// fmt.Println(slice2)
	// slice3 := []int{}
	// slice4 := []int{10, 11, 12, 13, 14}
	// fmt.Println(slice4)

	// if slice3 == nil {
	// 	println("slice 3 is nil ")
	// }

	a := 10
	b := a // copies the value of a to b
	b++
	println("Value of:", a)

}

func PrintSliceHeader(name string, slice []int) {
	fmt.Println(name, "-------------------------")
	fmt.Printf("Address of slice:%p\n", &slice)
	if slice != nil && len(slice) > 0 {
		fmt.Println("Ptr in the slice:", &slice[0])
	} else {
		fmt.Println("Ptr in the slice:nil")
	}
	fmt.Println("Len:", len(slice))
	fmt.Println("Cap:", cap(slice))
	fmt.Println(slice)
	fmt.Println("-----------------------")
}
